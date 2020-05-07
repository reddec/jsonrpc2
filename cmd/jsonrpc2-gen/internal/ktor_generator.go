package internal

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/reddec/godetector/deepparser"
	"go/ast"
	"strings"
	"text/template"
)

func (result *generationResult) GenerateKtor() string {
	var typer ktorGenerator
	typer.BeforeInspect = deepparser.RemoveJsonIgnoredFields
	for _, mth := range result.UsedMethods {
		for _, lt := range mth.LocalTypes(result.Import) {
			typer.Add(lt.Inspect)
		}
	}
	fm := sprig.TxtFuncMap()
	fm["firstLine"] = func(text string) string {
		return strings.Split(text, "\n")[0]
	}
	fm["kotlin"] = func(field interface{}) string {
		switch v := field.(type) {
		case typed:
			return typer.MapType(v.AST)
		case arg:
			return typer.MapType(v.AST)
		case *ast.TypeSpec:
			return typer.MapType(v.Type)
		case *deepparser.StField:
			return typer.MapField(v)
		default:
			panic("ktor?")
		}
	}
	fm["kotlinDefault"] = func(field interface{}) string {
		switch v := field.(type) {
		case typed:
			return typer.Default(v.AST)
		case arg:
			return typer.Default(v.AST)
		case *ast.TypeSpec:
			return typer.Default(v.Type)
		case *deepparser.StField:
			return typer.Default(v.AST.Type)
		default:
			panic("ktor?")
		}
	}
	fm["definitions"] = func() []*deepparser.Definition {
		if typer.Ordered == nil {
			for _, mth := range result.UsedMethods {
				for _, lt := range mth.LocalTypes(result.Import) {
					typer.Add(lt.Inspect)
				}
			}
		}
		var ans []*deepparser.Definition
		for _, d := range typer.Ordered {
			if len(d.StructFields()) > 0 {
				ans = append(ans, d)
			}
		}
		return ans
	}
	fm["enums"] = func() []*deepparser.Definition {
		var ans []*deepparser.Definition
		for _, d := range typer.Ordered {
			if d.IsTypeAlias() {
				fv := d.FindEnumValues()
				if len(fv) > 0 {
					ans = append(ans, d)
				}
			}
		}
		return ans
	}

	t := template.Must(template.New("").Funcs(fm).Parse(string(MustAsset("ktor.gotemplate"))))
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, result)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

type ktorGenerator struct {
	deepparser.Typer
}

func (tsg *ktorGenerator) mapBase(typeName string) string {
	if strings.HasPrefix(typeName, "int") {
		return "Long"
	}
	if typeName == "float32" {
		return "Float"
	}
	if typeName == "float64" {
		return "Double"
	}
	if strings.HasPrefix(typeName, "uint") {
		return "Long"
	}
	if typeName == "byte" {
		return "Byte"
	}
	if typeName == "string" {
		return "String"
	}
	if typeName == "bool" {
		return "Boolean"
	}
	return typeName
}

func (tsg *ktorGenerator) Default(t ast.Expr) string {
	if _, isRef := t.(*ast.StarExpr); isRef {
		return "null"
	}
	if isArray(t) {
		return "arrayOf()"
	}
	typeName := tsg.MapType(t)
	if typeName == "String" {
		return "\"\""
	}
	if typeName == "Boolean" {
		return "false"
	}
	if typeName == "Long" || typeName == "Int" || typeName == "Byte" || typeName == "Double" || typeName == "Float" {
		return "0.to" + typeName + "()"
	}
	return typeName + ".default()"
}

func (tsg *ktorGenerator) MapType(t ast.Expr) string {
	if v, ok := t.(*ast.Ident); ok {
		return tsg.mapBase(v.Name)
	}
	if acc, ok := t.(*ast.SelectorExpr); ok {
		return acc.Sel.Name
	}
	if ptr, ok := t.(*ast.StarExpr); ok {
		return tsg.MapType(ptr.X)
	}

	if arr, ok := t.(*ast.ArrayType); ok {
		return "Array<" + tsg.MapType(arr.Elt) + ">"
	}

	return "any"
}

func (tsg *ktorGenerator) MapField(st *deepparser.StField) string {
	tp := tsg.MapType(st.AST.Type)
	if st.Omitempty {
		return tp + "?"
	}
	return tp
}
