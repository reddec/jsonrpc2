package main

import (
	"github.com/dave/jennifer/jen"
	"github.com/jessevdk/go-flags"
	"github.com/reddec/jsonrpc2/cmd/jsonrpc2-gen/internal"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const version = "dev"

type Config struct {
	File        string `short:"i" long:"file" env:"GOFILE" description:"File to scan" required:"yes"`
	Interface   string `short:"I" long:"interface" env:"INTERFACE" description:"Interface to wrap" required:"yes"`
	Namespace   string `long:"namespace" env:"NAMESPACE" description:"Custom namespace for functions. If not defined - interface name will be used" default:""`
	Wrapper     string `short:"w" long:"wrapper" env:"WRAPPER" description:"Wrapper function name. If not defined - Register<interface> name will be used" default:""`
	Output      string `short:"o" long:"output" env:"OUTPUT" description:"Generated output destination (- means STDOUT)" default:"-"`
	Package     string `short:"p" long:"package" env:"PACKAGE" description:"Package name (can be override by output dir)" default:"events"`
	Doc         string `short:"d" long:"doc" env:"DOC" description:"Generate markdown documentation"`
	Case        string `short:"c" long:"case" env:"CASE" description:"Method name case style" default:"keep" choice:"keep" choice:"camel" choice:"pascal" choice:"snake" choice:"kebab"`
	URL         string `long:"url" env:"URL" description:"URL for examples in documentation" default:"https://example.com/api"`
	Interceptor bool   `short:"C" long:"interceptor" env:"INTERCEPTOR" description:"add interceptor for each method"`
}

func (c Config) GetCase() internal.Case {
	switch c.Case {
	case "camel":
		return internal.Camel
	case "pascal":
		return internal.Pascal
	case "snake":
		return internal.Snake
	case "kebab":
		return internal.Kebab
	case "keep":
		return internal.Keep
	default:
		return internal.Keep
	}
}

func main() {
	var config Config
	parser := flags.NewParser(&config, flags.Default)
	parser.LongDescription = "Generate tiny wrapper for JSON-RPC router\nAuthor: Baryshnikov Aleksandr <dev@baryshnikov.net>\nVersion: " + version
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	if config.Namespace == "" {
		config.Namespace = config.Interface
	}
	if config.Wrapper == "" {
		config.Wrapper = "Register" + config.Interface
	}

	var out *jen.File

	ev := internal.WrapperGenerator{
		TypeName:    config.Interface,
		FuncName:    config.Wrapper,
		Namespace:   config.Namespace,
		Case:        config.GetCase(),
		Interceptor: config.Interceptor,
	}
	result, err := ev.Generate(config.File)
	if err != nil {
		log.Fatal(err)
	}
	if config.Output != "-" {
		pkgImport, _ := internal.FindPackage(config.Output)
		out = jen.NewFilePathName(pkgImport, config.Package)
	} else {
		out = jen.NewFile(config.Package)
	}
	out.Add(result.Code)
	var output = os.Stdout
	if config.Output != "-" {
		output, err = os.Create(config.Output)
		if err != nil {
			panic(err)
		}
		defer output.Close()
	}
	_, _ = output.WriteString("// Code generated by jsonrpc2. DO NOT EDIT.\n")
	_, _ = output.WriteString("//go:generate " + strings.Join(os.Args, " ") + "\n")
	err = out.Render(output)
	if err != nil {
		panic(err)
	}

	if config.Doc != "" {
		err = ioutil.WriteFile(config.Doc, []byte(result.WithDocAddress(config.URL).GenerateMarkdown()), 0755)
		if err != nil {
			panic(err)
		}
	}
}
