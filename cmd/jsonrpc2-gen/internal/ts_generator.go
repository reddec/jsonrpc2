package internal

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/reddec/godetector/deepparser"
	"go/ast"
	"strings"
	"text/template"
)

func (result *generationResult) GenerateTS(shimFiles ...string) string {
	var tsg tsGenerator
	tsg.BeforeInspect = deepparser.RemoveJsonIgnoredFields
	fm := sprig.TxtFuncMap()
	fm["firstLine"] = func(text string) string {
		return strings.Split(text, "\n")[0]
	}
	fm["typescript"] = func(field interface{}) string {
		switch v := field.(type) {
		case typed:
			return tsg.MapType(v.AST)
		case arg:
			return tsg.MapType(v.AST)
		case *deepparser.StField:
			return tsg.MapField(v)
		default:
			panic("ts?")
		}
	}
	fm["definitions"] = func() []*deepparser.Definition {
		if tsg.Ordered == nil {
			for _, mth := range result.UsedMethods {
				for _, lt := range mth.LocalTypes(result.Import) {
					tsg.Add(lt.Inspect)
				}
			}
		}
		var ans []*deepparser.Definition
		for _, d := range tsg.Ordered {
			if len(d.StructFields()) > 0 {
				ans = append(ans, d)
			}
		}
		return ans
	}
	fm["enums"] = func() []*deepparser.Definition {
		var ans []*deepparser.Definition
		for _, d := range tsg.Ordered {
			if d.IsTypeAlias() {
				fv := d.FindEnumValues()
				if len(fv) > 0 {
					ans = append(ans, d)
				}
			}
		}
		return ans
	}
	fm["shim"] = func(def *deepparser.Definition) *typeShim {
		return tsg.FindShim(def.Import.Path + "@" + def.TypeName)
	}
	tsg.Shim(typeShim{
		Qual:    "time@Duration",
		Content: "type Duration = string; // suffixes: ns, us, ms, s, m, h",
	})
	tsg.Shim(typeShim{
		Qual:    "time@Time",
		Content: "type Time = string;",
	})
	for _, file := range shimFiles {
		if file != "" {
			err := tsg.ShimFromYamlFile(file)
			if err != nil {
				panic(err)
			}
		}
	}
	t := template.Must(template.New("").Funcs(fm).Parse(string(MustAsset("ts.gotemplate"))))
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, result)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

type tsGenerator struct {
	deepparser.Typer
	typesShim
}

func (tsg *tsGenerator) mapBase(typeName string) string {
	if strings.HasPrefix(typeName, "int") ||
		strings.HasPrefix(typeName, "float") ||
		strings.HasPrefix(typeName, "uint") ||
		typeName == "byte" {
		return "number"
	}
	if typeName == "string" {
		return "string"
	}
	if typeName == "bool" {
		return "boolean"
	}
	return typeName
}

func (tsg *tsGenerator) MapType(t ast.Expr) string {
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

func (tsg *tsGenerator) MapField(st *deepparser.StField) string {
	tp := tsg.MapType(st.AST.Type)
	if st.Omitempty {
		return tp + " | null"
	}
	return tp
}
