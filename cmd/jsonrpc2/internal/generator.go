package internal

import (
	"bytes"
	"github.com/dave/jennifer/jen"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strconv"
)

const Import = "github.com/reddec/jsonrpc2"

type WrapperGenerator struct {
	TypeName  string
	FuncName  string
	Namespace string
}

func (wg *WrapperGenerator) Generate(filename string) (jen.Code, error) {
	fs := token.NewFileSet()
	p, err := parser.ParseFile(fs, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	info, err := CollectInfo(p, fs)
	if err != nil {
		return nil, err
	}

	return wg.generateFunction(info, fs, p), nil
}

func (wg *WrapperGenerator) generateFunction(info *Interface, fs *token.FileSet, file *ast.File) jen.Code {
	return jen.Func().Id(wg.FuncName).Params(jen.Id("router").Op("*").Qual(Import, "Router"), jen.Id("wrap").Id(wg.TypeName)).BlockFunc(func(group *jen.Group) {
		for _, method := range info.Methods {
			if ast.IsExported(method.Name) {
				group.Id("router").Dot("RegisterFunc").Call(jen.Lit(method.Qual(wg.Namespace)), wg.generateLambda(method, fs, file)).Line()
			}
		}
	})
}

func (wg *WrapperGenerator) generateLambda(method *Method, fs *token.FileSet, file *ast.File) jen.Code {
	return jen.Func().Params(jen.Id("params").Qual("encoding/json", "RawMessage"), jen.Id("positional").Bool()).Call(jen.Interface(), jen.Error()).BlockFunc(func(group *jen.Group) {
		var argNames []string
		group.Var().Id("args").StructFunc(func(st *jen.Group) {
			if method.Type.Params == nil {
				return
			}
			argNames = make([]string, 0, len(method.Type.Params.List))
			for _, arg := range method.Type.Params.List {
				for _, argName := range arg.Names {
					name := "Arg" + strconv.Itoa(len(argNames))
					argNames = append(argNames, name)
					st.Id(name).Id(astPrint(arg.Type, fs)).Tag(map[string]string{
						"json": argName.Name,
					})
				}
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
		group.Return().Id("wrap").Dot(method.Name).CallFunc(func(params *jen.Group) {
			for _, arg := range argNames {
				params.Id("args").Dot(arg)
			}
		})
	})
}

func astPrint(t ast.Node, fs *token.FileSet) string {
	var buf bytes.Buffer
	printer.Fprint(&buf, fs, t)
	return buf.String()
}
