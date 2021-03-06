package internal

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

const Import = "github.com/reddec/jsonrpc2"

type Case int

const (
	Keep Case = iota
	Camel
	Pascal
	Snake
	Kebab
)

func (c Case) Convert(text string) string {
	switch c {
	case Camel:
		return strcase.ToLowerCamel(text)
	case Pascal:
		return strcase.ToCamel(text)
	case Snake:
		return strcase.ToSnake(text)
	case Kebab:
		return strcase.ToKebab(text)
	case Keep:
		return text
	default:
		return text
	}
}

type WrapperGenerator struct {
	TypeName                  string
	FuncName                  string
	Namespace                 string
	CustomHandlerMethodPrefix string
	CustomHandlers            []string // custom handlers for types (*import@id)
	Case                      Case
	Interceptor               bool
}

func (wg *WrapperGenerator) Qual(mg *Method) string {
	name := wg.Case.Convert(mg.Name)
	if wg.Namespace != "" {
		return wg.Namespace + "." + name
	}
	return name
}

func (wg *WrapperGenerator) Name() string {
	if wg.Namespace != "" {
		return wg.Namespace
	}
	return wg.TypeName
}

type generationResult struct {
	Code        jen.Code
	Generator   WrapperGenerator
	Service     *Interface
	UsedMethods []*Method
	Import      string
	DocAddr     string
}

func (wg *WrapperGenerator) Generate(filename string) (*generationResult, error) {
	importPath, err := FindPackage(filepath.Dir(filename))
	if err != nil {
		return nil, fmt.Errorf("detect package for source file %s: %v", filename, err)
	}
	fs := token.NewFileSet()
	p, err := parser.ParseFile(fs, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	info, err := CollectInfo(wg.TypeName, p, fs, filename)
	if err != nil {
		return nil, err
	}
	code, methods := wg.generateFunction(info, fs, p, importPath)
	return &generationResult{
		Code:        code,
		Generator:   *wg,
		Service:     info,
		Import:      importPath,
		UsedMethods: methods,
	}, nil
}

func (wg *WrapperGenerator) generateFunction(info *Interface, fs *token.FileSet, file *ast.File, importPath string) (jen.Code, []*Method) {
	qual := jen.Id(wg.TypeName)
	if importPath != "" {
		qual = jen.Qual(importPath, wg.TypeName)
	}

	var indexedCustomHandlers = make(map[string]typed)
	for _, typeID := range wg.CustomHandlers {
		var ops string
		if typeID[0] == '*' {
			ops = "*"
			typeID = typeID[1:]
		}
		imptp := strings.SplitN(typeID, "@", 2)
		indexedCustomHandlers[typeID] = typed{
			Type:   imptp[1],
			Import: imptp[0],
			Ops:    ops,
		}
	}

	var usedTypes = make(map[string]typed)

	// Detect is custom handler needed here
	for _, method := range info.Methods {
		for _, arg := range method.Args() {
			typeID := arg.globalQual(importPath)
			if tp, ok := indexedCustomHandlers[typeID]; ok && tp.Ops == arg.Ops {
				usedTypes[typeID] = arg.typed
			}
		}
	}

	var customTypeHandler jen.Code
	if len(usedTypes) > 0 {
		// generate type handler
		customTypeHandler = jen.InterfaceFunc(func(group *jen.Group) {
			for _, typed := range usedTypes {
				group.Id(wg.CustomHandlerMethodPrefix+typed.Type).Params(jen.Id("ctx").Qual("context", "Context"), jen.Id("value").Add(typed.Qual(importPath))).Error()
			}
		})
	}

	var usedMethods []*Method
	code := jen.Func().Id(wg.MustRender(wg.FuncName)).ParamsFunc(func(params *jen.Group) {
		params.Id("router").Op("*").Qual(Import, "Router")
		params.Id("wrap").Add(qual)
		if len(usedTypes) > 0 {
			params.Id("typeHandler").Add(customTypeHandler)
		}
		if wg.Interceptor {
			params.Id("interceptor").Func().Params(jen.Id("ctx").Qual("context", "Context"), jen.Id("methodName").String(), jen.Id("params").Index().Interface()).Error()
		}
	}).Index().String().BlockFunc(func(group *jen.Group) {
		for _, method := range info.Methods {
			if ast.IsExported(method.Name) {
				group.Id("router").Dot("RegisterFunc").Call(jen.Lit(wg.Qual(method)), wg.generateLambda(method, fs, file, importPath, usedTypes)).Line()
				usedMethods = append(usedMethods, method)
			}
		}
		group.Return().Index().String().ValuesFunc(func(values *jen.Group) {
			for _, m := range usedMethods {
				values.Lit(wg.Qual(m))
			}
		})
	})
	return code, usedMethods
}

func (wg *WrapperGenerator) generateLambda(method *Method, fs *token.FileSet, file *ast.File, importPath string, usedCustomTypesHandlers map[string]typed) jen.Code {
	return jen.Func().Params(jen.Id("ctx").Qual("context", "Context"), jen.Id("params").Qual("encoding/json", "RawMessage"), jen.Id("positional").Bool()).Call(jen.Interface(), jen.Error()).BlockFunc(func(group *jen.Group) {
		var argNames []string
		args := method.Args()
		if len(args) > 0 {
			group.Var().Id("args").StructFunc(func(st *jen.Group) {
				for _, arg := range method.Args() {
					name := "Arg" + strconv.Itoa(len(argNames))
					argNames = append(argNames, name)
					st.Id(name).Add(arg.Qual(importPath)).Tag(map[string]string{
						"json": arg.Name,
					})
				}
			})

			group.Var().Id("err").Error()
			group.If().Id("positional").BlockFunc(func(pos *jen.Group) {
				pos.Err().Op("=").Qual(Import, "UnmarshalArray").CallFunc(func(params *jen.Group) {
					params.Id("params")
					for _, arg := range argNames {
						params.Op("&").Id("args").Dot(arg)
					}
				})
			}).Else().BlockFunc(func(named *jen.Group) {
				named.Err().Op("=").Qual("encoding/json", "Unmarshal").Call(jen.Id("params"), jen.Op("&").Id("args"))
			})
			group.If().Err().Op("!=").Nil().BlockFunc(func(failed *jen.Group) {
				failed.Return(jen.Nil(), jen.Err())
			})

			for i, arg := range args {
				if handler, ok := usedCustomTypesHandlers[arg.globalQual(importPath)]; ok && handler.Ops == arg.Ops {
					argName := argNames[i]
					group.Err().Op("=").Id("typeHandler").Dot(wg.CustomHandlerMethodPrefix+arg.Type).Call(jen.Id("ctx"), jen.Id("args").Dot(argName))
					group.If().Err().Op("!=").Nil().BlockFunc(func(failed *jen.Group) {
						failed.Return(jen.Nil(), jen.Err())
					})
				}
			}
		}

		if wg.Interceptor {
			group.If(jen.Err().Op(":=").Id("interceptor").Call(jen.Id("ctx"), jen.Lit(wg.Qual(method)), jen.Index().Interface().ValuesFunc(func(params *jen.Group) {
				for _, arg := range argNames {
					params.Id("args").Dot(arg)
				}
			})), jen.Err().Op("!=").Nil()).BlockFunc(func(failed *jen.Group) {
				failed.Return(jen.Nil(), jen.Err())
			})
		}
		group.Return().Id("wrap").Dot(method.Name).CallFunc(func(params *jen.Group) {
			params.Id("ctx")
			for _, arg := range argNames {
				params.Id("args").Dot(arg)
			}
		})
	})
}

func (wg *WrapperGenerator) MustRender(templateText string) string {
	t := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).Parse(templateText))
	var out bytes.Buffer
	err := t.Execute(&out, wg)
	if err != nil {
		panic(err)
	}
	return out.String()
}

func astPrint(t ast.Node, fs *token.FileSet) string {
	var buf bytes.Buffer
	printer.Fprint(&buf, fs, t)
	return buf.String()
}
