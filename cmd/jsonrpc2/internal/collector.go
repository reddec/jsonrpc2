package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"strings"
)

type Interface struct {
	Name    string
	Comment string
	Methods []*Method
}

type Method struct {
	Name       string
	Definition *ast.Field
	Type       *ast.FuncType
	fs         *token.FileSet
}

func (mg *Method) Qual(namespace string) string {
	if namespace != "" {
		return namespace + "." + mg.Name
	}
	return mg.Name
}

func CollectInfo(file *ast.File, fs *token.FileSet) (*Interface, error) {
	var (
		name        string
		comment     string
		prevComment string
	)
	var srv *Interface
	for _, def := range file.Decls {
		ast.Inspect(def, func(node ast.Node) bool {
			switch v := node.(type) {
			case *ast.CommentGroup:
				prevComment = v.Text()
			case *ast.TypeSpec:
				name = v.Name.Name
				comment = strings.TrimSpace(prevComment)
				prevComment = ""
			case *ast.InterfaceType:
				srv = &Interface{
					Name:    name,
					Comment: comment,
					Methods: nil,
				}

				for _, fn := range v.Methods.List {
					tp := fn.Type.(*ast.FuncType)
					if !isMethodValid(tp) {
						log.Println("method", fn.Names[0].Name, "has unsupported signature")
						continue
					}
					srv.Methods = append(srv.Methods, &Method{
						Name:       fn.Names[0].Name,
						Definition: fn,
						Type:       tp,
						fs:         fs,
					})
				}

				return false
			}
			return true
		})
	}
	if srv == nil {
		return nil, fmt.Errorf("interface %v not found", name)
	}
	return srv, nil
}

func isMethodValid(tp *ast.FuncType) bool {
	if tp.Results == nil {
		return false
	}
	list := tp.Results.List
	if len(list) != 2 {
		// payload + error
		return false
	}
	return true
}