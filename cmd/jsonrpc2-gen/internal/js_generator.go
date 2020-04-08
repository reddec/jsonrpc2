package internal

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"strings"
	"text/template"
)

func (result *generationResult) GenerateJS() string {
	fm := sprig.TxtFuncMap()
	fm["firstLine"] = func(text string) string {
		return strings.Split(text, "\n")[0]
	}
	t := template.Must(template.New("").Funcs(fm).Parse(string(MustAsset("js.gotemplate"))))
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, result)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
