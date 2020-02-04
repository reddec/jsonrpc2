package internal

import (
	"bytes"
	"text/template"
)

//go:generate go-bindata -pkg internal template.gotemplate
func (result *generationResult) GenerateMarkdown() string {
	t := template.Must(template.New("").Parse(string(MustAsset("template.gotemplate"))))
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
