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
	python.BeforeInspect = deepparser.RemoveJsonIgnoredFields
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
	fm["isdefined"] = func(st *deepparser.StField) bool {
		return python.isDefined(deepparser.RebuildTypeNameWithoutPackage(st.AST.Type))
	}
	fm["to_json"] = func(field interface{}, param string) string {
		var st ast.Expr
		switch v := field.(type) {
		case typed:
			st = v.AST
		case arg:
			st = v.AST
		case *deepparser.StField:
			st = v.AST.Type
		default:
			panic("ts?")
		}

		if python.isBytes(st) {
			return "encodebytes(" + param + ")"
		}
		if !python.isDefined(deepparser.RebuildTypeNameWithoutPackage(st)) {
			return param
		}
		if isArray(st) {
			return "[x.to_json() for x in " + param + "]"
		}
		return param + ".to_json()"
	}
	fm["from_json"] = func(field interface{}, param string) string {
		var st ast.Expr
		switch v := field.(type) {
		case typed:
			st = v.AST
		case arg:
			st = v.AST
		case *deepparser.StField:
			st = v.AST.Type
		default:
			panic("ts?")
		}

		if python.isBytes(st) {
			return "decodebytes(" + param + ")"
		}
		subType := deepparser.RebuildTypeNameWithoutPackage(st)
		if !python.isDefined(subType) {
			if isArray(st) {
				return param + " or []"
			}
			return param
		}
		if isArray(st) {
			return "[" + subType + ".from_json(x) for x in (" + param + " or [])]"
		}
		return subType + ".from_json(" + param + ")"
	}
	fm["definitions"] = func() []*deepparser.Definition {
		if python.Ordered == nil {
			for _, mth := range result.UsedMethods {
				for _, lt := range mth.LocalTypes(result.Import) {
					python.Add(lt.Inspect)
				}
			}
		}
		var ans []*deepparser.Definition
		for _, d := range python.Ordered {
			if len(d.StructFields()) > 0 {
				ans = append(ans, d)
			}
		}
		return ans
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
	txt := buffer.String()

	var imps []string
	for imp, mods := range python.Modules {
		var names []string
		for name := range mods {
			names = append(names, name)
		}
		if len(names) > 0 {
			imps = append(imps, "from "+imp+" import "+strings.Join(names, ", "))
		}
	}
	return strings.Replace(txt, "# imports", strings.Join(imps, "\n"), 1)
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
	if py.isBytes(t) {
		py.AddModule("base64", "encodebytes", "decodebytes")
		return "bytes"
	}
	if v, ok := t.(*ast.Ident); ok {
		if py.isDefined(v.Name) {
			return v.Name
		}
		return py.mapBase(v.Name)
	}
	if acc, ok := t.(*ast.SelectorExpr); ok {
		py.AddModule("typing", "Any")
		if py.isDefined(acc.Sel.Name) {
			return acc.Sel.Name
		}
		return "Any"
	}
	if ptr, ok := t.(*ast.StarExpr); ok {
		return py.MapType(ptr.X)
	}

	if arr, ok := t.(*ast.ArrayType); ok {
		py.AddModule("typing", "List", "Optional")
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

func (py *Python) isDefined(name string) bool {
	for _, p := range py.Ordered {
		if p.TypeName == name && len(p.StructFields()) > 0 {
			return true
		}
	}
	return false
}

func (py *Python) isBytes(tp ast.Expr) bool {
	if s, ok := tp.(*ast.ArrayType); ok {
		if e, ok := s.Elt.(*ast.Ident); ok {
			return e.Name == "byte"
		}
	}
	return false
}

func isArray(tp ast.Expr) bool {
	_, ok := tp.(*ast.ArrayType)
	return ok
}
