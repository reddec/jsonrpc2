package internal

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	gen := WrapperGenerator{
		TypeName:    "User",
		FuncName:    "RegisterUser",
		Namespace:   "User",
		Interceptor: true,
	}

	result, err := gen.Generate("../../../example/gen.go")
	if err != nil {
		t.Fatal(err)
	}
	f := jen.NewFilePathName("github.com/reddec/jsonrpc2/example", "example")
	f.Add(result.Code)
	err = f.Render(os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerationResult_GenerateGo(t *testing.T) {
	gen := WrapperGenerator{
		TypeName: "User",
		FuncName: "Register",
	}
	file := "../../../example/gen.go"
	result, err := gen.Generate(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.WithDocAddress("http://example.com/api").GenerateGo("client", true, "Def"))
}

func TestGenerateKtor(t *testing.T) {
	gen := WrapperGenerator{
		TypeName: "User",
		FuncName: "Register",
	}
	file := "../../../example/gen.go"
	if v := os.Getenv("KTOR_FILE"); v != "" {
		file = v
	}
	if t := os.Getenv("KTOR_TYPE"); t != "" {
		gen.TypeName = t
	}
	if ns := os.Getenv("KTOR_NS"); ns != "" {
		gen.Namespace = ns
	}
	result, err := gen.Generate(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.WithDocAddress("http://example.com/api").GenerateKtor())

}
