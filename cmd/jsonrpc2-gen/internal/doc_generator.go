package internal

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"net/url"
	"strings"
	"text/template"
)

//go:generate go-bindata -pkg internal template.gotemplate python.gotemplate js.gotemplate ts.gotemplate method_doc.gotemplate ktor.gotemplate
func (result *generationResult) GenerateMarkdown() string {
	fm := sprig.TxtFuncMap()
	fm["firstLine"] = func(text string) string {
		return strings.Split(text, "\n")[0]
	}
	t := template.Must(template.New("").Funcs(fm).Parse(string(MustAsset("template.gotemplate"))))
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, result)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (result *generationResult) WithDocAddress(address string) *generationResult {
	result.DocAddr = address
	return result
}

func (result *generationResult) WsAddr() string {
	u, err := url.Parse(result.DocAddr)
	if err != nil {
		return "ws://example.com/api"
	}

	u.Scheme = "ws"
	return u.String()
}

func (result *generationResult) MustRender(templateText string) string {
	t := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).Parse(templateText))
	var out bytes.Buffer
	err := t.Execute(&out, result)
	if err != nil {
		panic(err)
	}
	return out.String()
}
