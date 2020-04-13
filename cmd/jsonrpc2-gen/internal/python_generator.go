package internal

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/reddec/godetector/deepparser"
	"go/ast"
	"strings"
	"text/template"
)

func (result *generationResult) GeneratePython() string {
	var python Python

	fm := sprig.TxtFuncMap()
	fm["python"] = func(field interface{}) string {
		switch v := field.(type) {
		case typed:
			return python.MapType(v.AST)
		case arg:
			return python.MapType(v.AST)
		case *deepparser.StField:
			return python.MapField(v)
		default:
			panic("ts?")
		}
	}
	fm["firstLine"] = func(text string) string {
		return strings.Split(text, "\n")[0]
	}
	fm["definitions"] = func() []*deepparser.Definition {
		if python.Ordered == nil {
			for _, mth := range result.UsedMethods {
				for _, lt := range mth.LocalTypes(result.Import) {
					python.Add(lt.Inspect)
				}
			}
		}
		return python.Ordered
	}
	fm["imports"] = func() map[string]map[string]bool {
		return python.Modules
	}
	t := template.Must(template.New("").Funcs(fm).Parse(string(MustAsset("python.gotemplate"))))
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, result)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

type Python struct {
	deepparser.Typer
	Modules map[string]map[string]bool
}

func (py *Python) mapBase(typeName string) string {
	if strings.HasPrefix(typeName, "int") ||
		strings.HasPrefix(typeName, "uint") {
		return "int"
	}
	if typeName == "byte" {
		return "byte"
	}
	if strings.HasPrefix(typeName, "float") {
		return "float"
	}
	if typeName == "string" {
		return "str"
	}
	if typeName == "bool" {
		return "bool"
	}
	py.AddModule("typing", "Any")
	return "Any"
}

func (py *Python) MapType(t ast.Expr) string {
	if v, ok := t.(*ast.Ident); ok {
		return py.mapBase(v.Name)
	}
	if _, ok := t.(*ast.SelectorExpr); ok {
		py.AddModule("typing", "Any")
		return "Any"
	}
	if ptr, ok := t.(*ast.StarExpr); ok {
		return py.MapType(ptr.X)
	}

	if arr, ok := t.(*ast.ArrayType); ok {
		py.AddModule("typing", "List")
		return "List[" + py.MapType(arr.Elt) + "]"
	}
	py.AddModule("typing", "Any")
	return "Any"
}

func (py *Python) MapField(st *deepparser.StField) string {
	tp := py.MapType(st.AST.Type)
	if st.Omitempty {
		py.AddModule("typing", "Optional")
		return "Optional[" + tp + "]"
	}
	return tp
}

func (py *Python) AddModule(pyImportName string, names ...string) {
	if len(names) == 0 {
		return
	}
	if py.Modules == nil {
		py.Modules = make(map[string]map[string]bool)
	}
	if py.Modules[pyImportName] == nil {
		py.Modules[pyImportName] = make(map[string]bool)
	}
	for _, name := range names {
		py.Modules[pyImportName][name] = true
	}
}
