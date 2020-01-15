package internal

import (
	"github.com/dave/jennifer/jen"
	"os"
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
	f := jen.NewFile("xyz")
	f.Add(result.Code)
	err = f.Render(os.Stdout)
	if err != nil {
		panic(err)
	}

	// Output:
	//package xyz
	//
	//import (
	//	"encoding/json"
	//	jsonrpc2 "github.com/reddec/jsonrpc2"
	//)
	//
	//func RegisterUser(router *jsonrpc2.Router, wrap User) []string {
	//	router.RegisterFunc("User.Profile", func(params json.RawMessage, positional bool) (interface{}, error) {
	//		var args struct {
	//			Arg0 string `json:"token"`
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
	//		return wrap.Profile(args.Arg0)
	//	})
	//
	//	return []string{"User.Profile"}
	//}
}
