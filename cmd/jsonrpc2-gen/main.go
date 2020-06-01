package main

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/dave/jennifer/jen"
	"github.com/jessevdk/go-flags"
	"github.com/reddec/godetector"
	"github.com/reddec/jsonrpc2/cmd/jsonrpc2-gen/internal"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const version = "dev"

type Config struct {
	File                    string   `short:"i" long:"file" env:"GOFILE" description:"File to scan"`
	Interface               []string `short:"I" long:"interface" env:"INTERFACE" description:"Interface to wrap" required:"yes"`
	Namespace               string   `short:"N" long:"namespace" env:"NAMESPACE" description:"Custom namespace for functions. If not defined - interface name will be used" default:""`
	Wrapper                 string   `short:"w" long:"wrapper" env:"WRAPPER" description:"Wrapper function name. If not defined - Register<interface> name will be used" default:""`
	Output                  string   `short:"o" long:"output" env:"OUTPUT" description:"Generated output destination (- means STDOUT)" default:"-"`
	CustomTypeHandlerPrefix string   `yaml:"custom_type_handler_prefix" long:"custom-type-handler-prefix" env:"CUSTOM_TYPE_HANDLER_PREFIX" description:"Custom prefix for methods for custom handlers" default:"Validate"`
	CustomTypeHandler       []string `yaml:"custom_type_handler" short:"T" long:"custom-type-handler" env:"CUSTOM_TYPE_HANDLER" description:"Handlers for custom types"`
	Package                 string   `short:"p" long:"package" env:"PACKAGE" description:"Package name (can be override by output dir)" default:"events"`
	Doc                     string   `short:"d" long:"doc" env:"DOC" description:"Generate markdown documentation"`
	DocShimFile             string   `long:"doc-shim-file" yaml:"doc_shim_file" env:"DOC_SHIM_FILE" description:"File for shim for markdown"`
	Python                  string   `short:"P" long:"python" env:"PYTHON" description:"Generate Python client" `
	JS                      string   `long:"js" env:"JS" description:"Generate JS client"`
	TS                      string   `long:"ts" env:"TS" description:"Generate TypeScript client"`
	TSShimFile              string   `long:"ts-shim-file" yaml:"ts_shim_file" env:"TS_SHIM_FILE" description:"Typescript shim file"`
	GO                      string   `long:"go" env:"GO" description:"Generate independent Golang client"`
	GoPackage               string   `long:"go-package" env:"GO_PACKAGE" description:"Destination go package" default:"client" yaml:"go_package"`
	GoDefault               string   `long:"go-default" env:"GO_DEFAULT" description:"Name for default constructor for Go clients" default:"Default" yaml:"go_default"`
	GoLinked                bool     `long:"go-linked" env:"GO_LINKED" description:"Link Go types instead of copy" yaml:"go_linked"`
	Ktor                    string   `long:"ktor" env:"KTOR" description:"KTOR (kotlin) client"`
	KtorShimFile            string   `long:"ktor-shim-file" env:"KTOR_SHIM_FILE" description:"KTOR shims" yaml:"ktor_shim_file"`
	Postman                 string   `long:"postman" env:"POSTMAN" description:"Generate Postman collection"`
	Case                    string   `short:"c" long:"case" env:"CASE" description:"Method name case style" default:"keep" choice:"keep" choice:"camel" choice:"pascal" choice:"snake" choice:"kebab"`
	URL                     string   `long:"url" env:"URL" description:"URL for examples in documentation" default:"https://example.com/api"`
	Interceptor             bool     `short:"C" long:"interceptor" env:"INTERCEPTOR" description:"add interceptor for each method"`
	Config                  string   `short:"f" long:"config" env:"CONFIG" description:"Location to configuration file"`
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
	if config.Config != "" {
		err = config.ApplyConfigFile()
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(config.Interface) == 0 {
		log.Fatal("at least one interface should be specified")
	}
	for i, interfaceName := range config.Interface {
		processInterface(config, interfaceName, i == 0)
	}
}

func processInterface(config Config, interfaceName string, appendCli bool) {
	if config.Namespace == "" {
		config.Namespace = interfaceName
	}
	if config.Wrapper == "" {
		config.Wrapper = "Register" + interfaceName
	}

	var out *jen.File

	ev := internal.WrapperGenerator{
		TypeName:                  interfaceName,
		FuncName:                  config.Wrapper,
		Namespace:                 config.MustRender(config.Namespace, interfaceName),
		Case:                      config.GetCase(),
		Interceptor:               config.Interceptor,
		CustomHandlers:            config.CustomTypeHandler,
		CustomHandlerMethodPrefix: config.CustomTypeHandlerPrefix,
	}
	result, err := ev.Generate(config.File)
	if err != nil {
		log.Fatal(err)
	}
	if config.Output != "-" {
		info, err := godetector.InspectDirectory(filepath.Dir(config.Output))
		if err != nil {
			log.Fatalln(err)
		}
		packageName := godetector.FindPackageNameByDir(info.ImportDir)
		out = jen.NewFilePathName(info.Import, packageName)
	} else {
		out = jen.NewFile(config.Package)
	}
	out.Add(result.Code)
	var output = os.Stdout
	if config.Output != "-" {
		output, err = os.Create(result.MustRender(config.Output))
		if err != nil {
			panic(err)
		}
		defer output.Close()
	}
	_, _ = output.WriteString("// Code generated by jsonrpc2. DO NOT EDIT.\n")
	if appendCli {
		_, _ = output.WriteString("//go:generate " + strings.Join(os.Args, " ") + "\n")
	}
	err = out.Render(output)
	if err != nil {
		panic(err)
	}

	if config.Doc != "" {
		err = writeFile(result.MustRender(config.Doc), []byte(result.WithDocAddress(config.URL).GenerateMarkdown(config.DocShimFile)), 0755)
		if err != nil {
			panic(err)
		}
	}
	if config.Python != "" {
		err = writeFile(result.MustRender(config.Python), []byte(result.WithDocAddress(config.URL).GeneratePython()), 0755)
		if err != nil {
			panic(err)
		}
	}
	if config.JS != "" {
		err = writeFile(result.MustRender(config.JS), []byte(result.WithDocAddress(config.URL).GenerateJS()), 0755)
		if err != nil {
			panic(err)
		}
	}
	if config.TS != "" {
		err = writeFile(result.MustRender(config.TS), []byte(result.WithDocAddress(config.URL).GenerateTS(config.TSShimFile)), 0755)
		if err != nil {
			panic(err)
		}
	}
	if config.Postman != "" {
		err = writeFile(result.MustRender(config.Postman), []byte(result.WithDocAddress(config.URL).GeneratePostman()), 0755)
		if err != nil {
			panic(err)
		}
	}
	if config.GO != "" {
		err = writeFile(result.MustRender(config.GO), []byte(result.WithDocAddress(config.URL).GenerateGo(result.MustRender(config.GoPackage), config.GoLinked, result.MustRender(config.GoDefault))), 0755)
		if err != nil {
			panic(err)
		}
	}
	if config.Ktor != "" {
		err = writeFile(result.MustRender(config.Ktor), []byte(result.WithDocAddress(config.URL).GenerateKtor(config.KtorShimFile)), 0755)
		if err != nil {
			panic(err)
		}
	}
}

func (c *Config) ApplyConfigFile() error {
	abs, err := filepath.Abs(c.Config)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(c.Config)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}
	root := filepath.Dir(abs)
	if c.Output != "-" {
		c.Output = resolvePath(root, c.Output)
	}
	c.Doc = resolvePath(root, c.Doc)
	c.DocShimFile = resolvePath(root, c.DocShimFile)
	c.Python = resolvePath(root, c.Python)
	c.JS = resolvePath(root, c.JS)
	c.TS = resolvePath(root, c.TS)
	c.TSShimFile = resolvePath(root, c.TSShimFile)
	c.GO = resolvePath(root, c.GO)
	c.Postman = resolvePath(root, c.Postman)
	c.Ktor = resolvePath(root, c.Ktor)
	c.KtorShimFile = resolvePath(root, c.KtorShimFile)
	return nil
}

func resolvePath(root, path string) string {
	if path == "" {
		return path
	}
	if !filepath.IsAbs(path) && !strings.HasPrefix(path, "./") {
		return filepath.Join(root, path)
	}
	return path
}

func writeFile(path string, content []byte, perm os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), perm); err != nil {
		return err
	}
	return ioutil.WriteFile(path, content, perm)
}

func (c Config) MustRender(templateText, interfaceName string) string {
	t := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).Parse(templateText))
	var out bytes.Buffer
	err := t.Execute(&out, map[string]interface{}{
		"Config":    c,
		"Interface": interfaceName,
	})
	if err != nil {
		panic(err)
	}
	return out.String()
}
