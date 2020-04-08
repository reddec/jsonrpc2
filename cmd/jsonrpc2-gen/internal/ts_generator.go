package internal

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"go/ast"
	"strings"
	"text/template"
)

func (result *generationResult) GenerateTS() string {
	var tsg tsGenerator
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
		case *stField:
			return tsg.MapField(v)
		default:
			panic("ts?")
		}
	}
	fm["definitions"] = func() []*Definition {
		if tsg.Ordered == nil {
			for _, mth := range result.UsedMethods {
				for _, lt := range mth.LocalTypes(result.Import) {
					tsg.Inspect(lt.Inspect)
				}
			}
		}
		return tsg.Ordered
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
	Ordered []*Definition
	Parsed  map[string]*Definition
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

func (tsg *tsGenerator) MapField(st *stField) string {
	tp := tsg.MapType(st.AST.Type)
	if st.Omitempty {
		return tp + " | null"
	}
	return tp
}

func (tsg *tsGenerator) Inspect(def *Definition) {
	uid := def.Import.Path + "@" + def.TypeName
	_, ok := tsg.Parsed[uid]
	if ok {
		return
	}
	if tsg.Parsed == nil {
		tsg.Parsed = make(map[string]*Definition)
	}
	tsg.Ordered = append(tsg.Ordered, def)
	def.removeJSONIgnoredFields()
	tsg.Parsed[uid] = def

	for _, f := range def.StructFields() {
		alias := detectPackageInType(f.AST.Type)
		typeName := rebuildTypeNameWithoutPackage(f.AST.Type)
		def := findDefinitionFromAst(typeName, alias, def.File, def.FileDir)

		if def != nil {
			tsg.Inspect(def)
		}
	}
}
