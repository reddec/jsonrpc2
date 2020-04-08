package internal

import (
	"bytes"
	"text/template"
)

func (result *generationResult) GenerateJS() string {
	t := template.Must(template.New("").Parse(string(MustAsset("js.gotemplate"))))
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, result)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
