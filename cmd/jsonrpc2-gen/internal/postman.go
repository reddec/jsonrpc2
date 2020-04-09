package internal

import (
	"bytes"
	"encoding/json"
	"github.com/Masterminds/sprig"
	"github.com/reddec/jsonrpc2"
	"net/url"
	"strings"
	"text/template"
)

func (result *generationResult) GeneratePostman() string {
	var root postman

	root.Info.ID = result.Import + "@" + result.Service.Name
	root.Info.Name = result.Service.Name
	root.Info.Description = "# " + result.Service.Name + "\n\n" + result.Service.Comment
	root.Info.Schema = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"

	for _, method := range result.Service.Methods {
		var item postmanItem

		item.Name = method.Name
		item.Request.Method = "POST"
		item.Request.Header = []postmanHeader{
			{Key: "Content-Type", Value: "application/json", Type: "text"},
		}
		item.Request.Body.Mode = "raw"
		data, _ := json.MarshalIndent(jsonrpc2.Request{
			Version: "2.0",
			Method:  result.Generator.Qual(method),
			ID:      json.RawMessage("1"),
			Params:  json.RawMessage("{}"),
		}, "", "  ")
		item.Request.Body.Raw = string(data)
		item.Request.Body.Options.Raw.Language = "json"

		if u, err := url.Parse(result.DocAddr); err == nil && result.DocAddr != "" {
			item.Request.URL.Raw = result.DocAddr
			item.Request.URL.Protocol = u.Scheme
			item.Request.URL.Host = strings.Split(u.Hostname(), ".")
			item.Request.URL.Port = u.Port()
			item.Request.URL.Path = strings.Split(u.Path, "/")
		}

		item.Request.Description = result.renderDocMethod(method)

		root.Item = append(root.Item, item)
	}

	data, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (result *generationResult) renderDocMethod(method *Method) string {
	fm := sprig.TxtFuncMap()
	fm["firstLine"] = func(text string) string {
		return strings.Split(text, "\n")[0]
	}
	t := template.Must(template.New("").Funcs(fm).Parse(string(MustAsset("method_doc.gotemplate"))))
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, map[string]interface{}{
		"Global": result,
		"Method": method,
	})
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

type postman struct {
	Info struct {
		ID          string `json:"_postman_id"`
		Name        string `json:"name"`
		Schema      string `json:"schema"`
		Description string `json:"description"`
	} `json:"info"`

	Item []postmanItem `json:"item"`
}

type postmanHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type postmanItem struct {
	Name    string `json:"name"`
	Request struct {
		Method string          `json:"method"`
		Header []postmanHeader `json:"header"`
		Body   struct {
			Mode    string `json:"mode"`
			Raw     string `json:"raw"`
			Options struct {
				Raw struct {
					Language string `json:"language"`
				} `json:"raw"`
			} `json:"options"`
		} `json:"body"`
		URL struct {
			Raw      string   `json:"raw"`
			Protocol string   `json:"protocol"`
			Host     []string `json:"host"`
			Port     string   `json:"port"`
			Path     []string `json:"path"`
		} `json:"url"`
		Description string `json:"description"`
	} `json:"request"`
}
