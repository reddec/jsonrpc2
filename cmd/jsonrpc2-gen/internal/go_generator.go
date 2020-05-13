package internal

import (
	"bytes"
	"github.com/dave/jennifer/jen"
	"github.com/reddec/godetector"
	"github.com/reddec/godetector/deepparser"
	"go/ast"
)

func (result *generationResult) GenerateGo(pkg string, linkedTypes bool) string {
	var typer GoType
	typer.BeforeInspect = deepparser.RemoveJsonIgnoredFields
	for _, mth := range result.UsedMethods {
		for _, lt := range mth.LocalTypes(result.Import) {
			typer.Add(lt.Inspect)
		}
	}

	var types = jen.Empty()
	if !linkedTypes {
		for _, definition := range typer.Ordered {
			if definition.Import.Type == godetector.GoRoot {
				continue
			}
			if !definition.IsTypeAlias() {
				continue
			}
			fd, _ := typer.MapDefinition(definition)
			types.Type().Id(definition.TypeName).Add(fd).Line()
			types.Const().DefsFunc(func(constTypes *jen.Group) {
				for _, value := range definition.FindEnumValues() {
					constTypes.Id(value.Name).Id(definition.TypeName).Op("=").Id(value.Value)
				}
			}).Line()

		}

		for _, definition := range typer.Ordered {
			if definition.Import.Type == godetector.GoRoot {
				continue
			}
			if !shouldBeDefined(definition) {
				continue
			}
			if !definition.IsTypeAlias() {
				types.Comment(definition.Type.Comment.Text()).Line().Type().Id(definition.TypeName).StructFunc(func(group *jen.Group) {
					for _, field := range definition.StructFields() {
						fd, _ := typer.MapField(field)
						group.Id(field.Name).Add(fd).Tag(map[string]string{"json": field.Tag})
					}
				}).Line().Line()
			}
		}
	}

	apiClient := result.Service.Name + "Client"

	if result.DocAddr != "" {
		types.Func().Id("Default").Params().Op("*").Id(apiClient).BlockFunc(func(group *jen.Group) {
			group.Return().Op("&").Id(apiClient).ValuesFunc(func(init *jen.Group) {
				init.Id("BaseURL").Op(":").Lit(result.DocAddr)
			})
		}).Line().Line()
	}

	types.Type().Id(apiClient).StructFunc(func(group *jen.Group) {
		group.Id("BaseURL").String()
		group.Id("sequence").Uint64()
	}).Line().Line()

	for _, method := range result.UsedMethods {
		ret, _ := typer.MapTyped(method.Return())
		if linkedTypes {
			ret = method.Return().Qual(result.Import)
		}
		types.Comment(method.Comment()).Line().Func().Params(jen.Id("impl").Op("*").Id(apiClient)).Id(method.Name).ParamsFunc(func(params *jen.Group) {
			params.Id("ctx").Qual("context", "Context")
			for _, param := range method.Args() {
				if linkedTypes {
					params.Id(param.Name).Add(param.Qual(result.Import))
				} else {
					fd, _ := typer.MapTyped(param.typed)
					params.Id(param.Name).Add(fd)
				}
			}
		}).Params(
			jen.Id("reply").Add(ret),
			jen.Err().Error(),
		).BlockFunc(func(methodType *jen.Group) {
			methodType.Err().Op("=").Qual("github.com/reddec/jsonrpc2/client", "CallHTTP").CallFunc(func(callParams *jen.Group) {
				callParams.Id("ctx")
				callParams.Id("impl").Dot("BaseURL")
				callParams.Lit(result.Generator.Qual(method))
				callParams.Qual("sync/atomic", "AddUint64").Call(jen.Op("&").Id("impl").Dot("sequence"), jen.Lit(1))
				callParams.Op("&").Id("reply")
				for _, param := range method.Args() {
					callParams.Id(param.Name)
				}
			})
			methodType.Return()
		}).Line().Line()
	}

	out := jen.NewFile(pkg)
	out.Add(types)

	var buf bytes.Buffer
	err := out.Render(&buf)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

type GoType struct {
	deepparser.Typer
}

func (py *GoType) MapType(t ast.Expr) (jen.Code, bool) {
	if v, ok := t.(*ast.Ident); ok {
		if py.isDefined(v.Name) {
			return jen.Id(v.Name), true
		}
		return jen.Id(v.Name), true
	}
	if acc, ok := t.(*ast.SelectorExpr); ok {
		if py.isDefined(acc.Sel.Name) {
			return jen.Id(acc.Sel.Name), true
		}
		return jen.Interface(), false
	}
	if ptr, ok := t.(*ast.StarExpr); ok {
		v, ok := py.MapType(ptr.X)
		if ok {
			return jen.Op("*").Add(v), ok
		}
		return v, ok
	}

	if arr, ok := t.(*ast.ArrayType); ok {
		v, _ := py.MapType(arr.Elt)
		return jen.Index().Add(v), false
	}
	return jen.Interface(), false
}

func (py *GoType) MapField(st *deepparser.StField) (jen.Code, bool) {
	if st.Definition != nil && st.Definition.Import.Type == godetector.GoRoot {
		return jen.Qual(st.Definition.Import.Path, rebuildTypeNameWithoutPackage(st.AST.Type)), true
	}
	return py.MapType(st.AST.Type)
}

func (py *GoType) MapDefinition(def *deepparser.Definition) (jen.Code, bool) {
	var ops = jen.Empty()
	_, isRef := def.Type.Type.(*ast.StarExpr)
	if isRef {
		ops = jen.Op("*")
	}

	if def.Import.Type == godetector.GoRoot {
		return ops.Qual(def.Import.Path, rebuildTypeNameWithoutPackage(def.Type.Type)), true
	}
	return py.MapType(def.Type.Type)
}

func (py *GoType) MapTyped(tp typed) (jen.Code, bool) {
	var ops = jen.Id(tp.Ops)
	def := py.Parsed[tp.Import+"@"+tp.Type]
	if def == nil {
		return py.MapType(tp.AST)
	}

	if def.Import.Type == godetector.GoRoot {
		return ops.Qual(def.Import.Path, def.TypeName), true
	}

	return py.MapType(tp.AST)
}

func (py *GoType) isDefined(name string) bool {
	for _, p := range py.Ordered {
		if p.TypeName == name && shouldBeDefined(p) {
			return true
		}
	}
	return false
}

func shouldBeDefined(p *deepparser.Definition) bool {
	if (p.IsTypeAlias() && len(p.FindEnumValues()) > 0) || len(p.StructFields()) > 0 {
		return true
	}
	return false
}
