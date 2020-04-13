package internal

import (
	"bufio"
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/reddec/godetector/deepparser"
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Interface struct {
	Name    string
	Comment string
	Methods []*Method
	Imports []*ast.ImportSpec
}

func (iface *Interface) LookupForImport(pkg string) string {
	if pkg == "" {
		return ""
	}
	// dummy lookup - expect that directory is equal to package name
	// priority to aliases
	for _, imp := range iface.Imports {
		if imp.Name != nil {
			if imp.Name.Name == pkg {
				path, _ := strconv.Unquote(imp.Path.Value)
				return path
			}
		}
	}
	for _, imp := range iface.Imports {
		path, _ := strconv.Unquote(imp.Path.Value)
		if filepath.Base(path) == pkg {
			return path
		}
	}
	return ""
}

type Method struct {
	Name       string
	Definition *ast.Field
	Type       *ast.FuncType
	Interface  *Interface
	fs         *token.FileSet
	file       *ast.File
	fileName   string
}

func (mg *Method) Comment() string {
	return strings.TrimSpace(mg.Definition.Doc.Text())
}

func (mg *Method) ReturnType() string {
	return astPrint(mg.Type.Results.List[0].Type, mg.fs)
}

func (mg *Method) LocalTypes(parentImportPath string) []LocalType {
	var usedTypes = map[string]typed{}
	// collect types from arguments
	for _, arg := range mg.Args() {
		if !isBuiltin(arg.Type) {
			usedTypes[arg.localQual()] = arg.typed
		}
	}
	// do not forget return type
	retType := mg.Return()
	if !isBuiltin(retType.Type) {
		usedTypes[retType.localQual()] = retType
	}

	// collect types definitions
	var ans = make([]LocalType, 0, len(usedTypes))
	for _, typeDef := range usedTypes {
		definition := deepparser.FindDefinitionFromAst(typeDef.Type, typeDef.Alias, mg.file, filepath.Dir(mg.fileName))
		if definition == nil {
			continue
		}
		definition.RemoveJSONIgnoredFields()
		fields := definition.StructFields()
		ans = append(ans, LocalType{
			Type:         definition.TypeName,
			Definition:   astPrint(definition.Type, definition.FS),
			IsStruct:     definition.IsStruct() && len(fields) > 0,
			StructFields: fields,
			Inspect:      definition,
		})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Type < ans[j].Type
	})
	return ans
}

type LocalType struct {
	Type         string
	Definition   string
	IsStruct     bool
	StructFields []*deepparser.StField
	Inspect      *deepparser.Definition
}
type typed struct {
	Type   string
	Import string
	Alias  string
	Ops    string
	AST    ast.Expr
}

func (a typed) localQual() string {
	return a.Import + "@" + a.Type
}

type arg struct {
	Name string
	typed
}

func (a arg) Qual(parentImportPath string) jen.Code {
	if a.Import == "" {
		if parentImportPath == "" || !ast.IsExported(a.Type) {
			return jen.Op(a.Ops).Id(a.Type)
		}
		return jen.Op(a.Ops).Qual(parentImportPath, a.Type)
	}
	return jen.Op(a.Ops).Qual(a.Import, a.Type)
}

func (mg *Method) Args() []arg {
	if mg.Type.Params == nil || len(mg.Type.Params.List) == 0 {
		return nil
	}
	var args []arg
	for _, t := range mg.Type.Params.List {
		alias := detectPackageInType(t.Type)
		var importPath = mg.Interface.LookupForImport(alias)
		for _, name := range t.Names {
			args = append(args, arg{
				Name: name.Name,
				typed: typed{
					Ops:    rebuildOps(t.Type),
					Type:   rebuildTypeNameWithoutPackage(t.Type),
					Import: importPath,
					Alias:  alias,
					AST:    t.Type,
				},
			})
		}
	}
	return args
}

func (mg *Method) Return() typed {
	if mg.Type.Results == nil || len(mg.Type.Results.List) == 0 {
		return typed{}
	}
	retType := mg.Type.Results.List[0].Type
	alias := detectPackageInType(retType)
	var importPath = mg.Interface.LookupForImport(alias)
	return typed{
		Ops:    rebuildOps(retType),
		Type:   rebuildTypeNameWithoutPackage(retType),
		Import: importPath,
		Alias:  alias,
		AST:    retType,
	}
}

func detectPackageInType(t ast.Expr) string {
	if acc, ok := t.(*ast.SelectorExpr); ok {
		return acc.X.(*ast.Ident).Name
	} else if ptr, ok := t.(*ast.StarExpr); ok {
		return detectPackageInType(ptr.X)
	} else if arr, ok := t.(*ast.ArrayType); ok {
		return detectPackageInType(arr.Elt)
	}
	return ""
}

func rebuildOps(t ast.Expr) string {
	if ptr, ok := t.(*ast.StarExpr); ok {
		return "*" + rebuildOps(ptr.X)
	}
	if arr, ok := t.(*ast.ArrayType); ok {
		return "[]" + rebuildOps(arr.Elt)
	}
	return ""
}

func rebuildTypeNameWithoutPackage(t ast.Expr) string {
	if v, ok := t.(*ast.Ident); ok {
		return v.Name
	}
	if ptr, ok := t.(*ast.StarExpr); ok {
		return rebuildTypeNameWithoutPackage(ptr.X)
	}
	if acc, ok := t.(*ast.SelectorExpr); ok {
		return acc.Sel.Name
	}
	if arr, ok := t.(*ast.ArrayType); ok {
		return rebuildTypeNameWithoutPackage(arr.Elt)
	}
	return ""
}

func CollectInfo(search string, file *ast.File, fs *token.FileSet, fileName string) (*Interface, error) {
	var (
		name        string
		comment     string
		prevComment string
		imports     []*ast.ImportSpec
	)
	var srv *Interface
	for _, def := range file.Decls {
		ast.Inspect(def, func(node ast.Node) bool {
			switch v := node.(type) {
			case *ast.ImportSpec:
				imports = append(imports, v)
			case *ast.CommentGroup:
				prevComment = v.Text()
			case *ast.TypeSpec:
				name = v.Name.Name
				comment = strings.TrimSpace(prevComment)
				prevComment = ""
			case *ast.InterfaceType:
				if name != search {
					comment = ""
					return true
				}
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
						Interface:  srv,
						fs:         fs,
						file:       file,
						fileName:   fileName,
					})
				}

				return false
			}
			return true
		})
	}
	if srv == nil {
		return nil, fmt.Errorf("interface %v not found", search)
	}
	srv.Imports = imports
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

func FindPackage(dir string) (string, error) {
	const Vendor = "vendor/"
	if strings.HasPrefix(dir, Vendor) {
		return dir[len(Vendor):], nil
	}
	dir, _ = filepath.Abs(dir)
	return findPackage(dir)
}

func findPackage(dir string) (string, error) {
	if dir == "" {
		return "", os.ErrNotExist
	}
	if isRootPackage(dir) {
		return "", nil
	}
	pkg, ok := isVendorPackage(dir)
	if ok {
		return pkg, nil
	}
	mod := filepath.Base(dir)
	top, err := findPackage(filepath.Dir(dir))
	if err != nil {
		return "", err
	}
	if top != "" {
		return top + "/" + mod, nil
	}
	return mod, nil
}

func isVendorPackage(path string) (string, bool) {
	path = filepath.Join(path, "go.mod")
	if fs, err := os.Stat(path); err != nil {
		return "", false
	} else if fs.IsDir() {
		return "", false
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	if !(scanner.Scan() && scanner.Scan()) {
		return "", false
	}
	pkg := scanner.Text()
	return pkg, true
}

func isRootPackage(path string) bool {
	GOPATH := filepath.Join(os.Getenv("GOPATH"), "src")
	GOROOT := filepath.Join(os.Getenv("GOROOT"), "src")
	return isRootOf(path, GOPATH) || isRootOf(path, GOROOT)
}

func isRootOf(path, root string) bool {
	root, _ = filepath.Abs(root)
	path, _ = filepath.Abs(path)
	return root == path
}

func isBuiltin(name string) bool {
	for _, k := range types.Typ {
		if k.Name() == name {
			return true
		}
	}
	if name == "error" {
		return true
	}
	return false
}
