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
	var reservedKeywords = []string{
		"as", "class", "break", "continue", "do", "else", "for", "fun", "false", "if",
		"in", "interface", "super", "return", "object", "package", "null", "is", "try",
		"throw", "true", "this", "typeof", "typealias", "when", "while", "val", "var",
	}

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
			return typer.MapTyped(v)
		case arg:
			return typer.MapTyped(v.typed)
		case *ast.TypeSpec:
			return typer.MapType(v.Type)
		case *deepparser.StField:
			return typer.MapField(v)
		default:
			panic("ktor?")
		}
	}
	fm["kotlinString"] = func(t string) string {
		if len(t) == 0 {
			return `""`
		}
		if t[0] == '"' {
			return t
		}
		return `"` + t + `"`
	}
	fm["escape"] = func(name string) string {
		for _, k := range reservedKeywords {
			if k == name {
				return name + "_"
			}
		}
		return name
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
			return typer.DefaultField(v)
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
	typer.Shim(typeShim{
		Qual:       "time@Time",
		Content:    "java.time.LocalDateTime",
		Initialize: "java.time.LocalDateTime.now()",
	})
	typer.Shim(typeShim{
		Qual:       "math/big@Int",
		Content:    "java.math.BigInteger",
		Initialize: "java.math.BigInteger.ZERO",
	})
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
	typesShim
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
	if tsg.isDefined(typeName) {
		return typeName
	}
	return "Any"
}

func (tsg *ktorGenerator) Default(t ast.Expr) string {
	if _, isRef := t.(*ast.StarExpr); isRef {
		return "null"
	}
	if isArray(t) {
		return "listOf<" + tsg.MapType(t.(*ast.ArrayType).Elt) + ">()"
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

func (tsg *ktorGenerator) DefaultField(v *deepparser.StField) string {
	if v.Definition != nil {
		shim := tsg.FindShim(v.Definition.Import.Path + "@" + v.Definition.TypeName)
		if shim != nil {
			return shim.Initialize
		}
	}
	return tsg.Default(v.AST.Type)
}

func (tsg *ktorGenerator) MapType(t ast.Expr) string {
	if v, ok := t.(*ast.Ident); ok {
		return tsg.mapBase(v.Name)
	}
	if acc, ok := t.(*ast.SelectorExpr); ok {
		return tsg.MapType(acc.Sel)
	}
	if ptr, ok := t.(*ast.StarExpr); ok {
		return tsg.MapType(ptr.X) + "?"
	}

	if arr, ok := t.(*ast.ArrayType); ok {
		return "List<" + tsg.MapType(arr.Elt) + ">"
	}

	return "Any"
}

func (tsg *ktorGenerator) MapTyped(t typed) string {
	if shim := tsg.FindShim(t.localQual()); shim != nil {
		return shim.Content
	}
	return tsg.MapType(t.AST)
}

func (tsg *ktorGenerator) MapField(st *deepparser.StField) string {
	var tp string
	if st.Definition != nil {
		shim := tsg.FindShim(st.Definition.Import.Path + "@" + st.Definition.TypeName)
		if shim != nil {
			tp = shim.Content
		}
	}
	if tp == "" {
		tp = tsg.MapType(st.AST.Type)
	}

	if st.Omitempty {
		return tp + "?"
	}
	return tp
}

func (tsg *ktorGenerator) isDefined(name string) bool {
	for _, p := range tsg.Ordered {
		if p.TypeName == name && shouldBeDefined(p) {
			return true
		}
	}
	return false
}
