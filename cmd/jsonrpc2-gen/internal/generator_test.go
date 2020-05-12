package internal

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"os"
	"testing"
)

func ExampleGenerate() {
	gen := WrapperGenerator{
		TypeName:  "User",
		FuncName:  "RegisterUser",
		Namespace: "User",
	}

	result, err := gen.Generate("../../../example/gen.go")
	if err != nil {
		panic(err)
	}
	f := jen.NewFilePathName("github.com/reddec/jsonrpc2/example", "example")
	f.Add(result.Code)
	err = f.Render(os.Stdout)
	if err != nil {
		panic(err)
	}

	// Output:
	//package example
	//
	//import (
	//	"encoding/json"
	//	jsonrpc2 "github.com/reddec/jsonrpc2"
	//	"math/big"
	//	"time"
	//)
	//
	//func RegisterUser(router *jsonrpc2.Router, wrap User) []string {
	//	router.RegisterFunc("User.Profile", func(params json.RawMessage, positional bool) (interface{}, error) {
	//		var args struct {
	//			Arg0 string    `json:"token"`
	//			Arg1 time.Time `json:"at"`
	//			Arg2 *big.Int  `json:"val"`
	//		}
	//		var err error
	//		if positional {
	//			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1, &args.Arg2)
	//		} else {
	//			err = json.Unmarshal(params, &args)
	//		}
	//		if err != nil {
	//			return nil, err
	//		}
	//		return wrap.Profile(args.Arg0, args.Arg1, args.Arg2)
	//	})
	//
	//	router.RegisterFunc("User.Latest", func(params json.RawMessage, positional bool) (interface{}, error) {
	//		var args struct {
	//			Arg0 []*time.Time `json:"times"`
	//		}
	//		var err error
	//		if positional {
	//			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
	//		} else {
	//			err = json.Unmarshal(params, &args)
	//		}
	//		if err != nil {
	//			return nil, err
	//		}
	//		return wrap.Latest(args.Arg0)
	//	})
	//
	//	return []string{"User.Profile", "User.Latest"}
	//}

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
