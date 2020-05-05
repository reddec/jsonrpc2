package internal

import (
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"strings"
)

type typeShim struct {
	Qual    string `yaml:"qual"` // import@type
	Content string `yaml:"content"`
}

func (ts *typeShim) TypeName() string {
	return strings.Split(ts.Qual, "@")[1]
}

type typesShim struct {
	shims map[string]*typeShim
}

func (ts *typesShim) ShimFromYamlFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return ts.ShimFromYaml(f)
}

func (ts *typesShim) ShimFromYaml(reader io.Reader) error {
	dec := yaml.NewDecoder(reader)
	for {
		var shim typeShim
		err := dec.Decode(&shim)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ts.Shim(shim)
	}
	return nil
}

func (ts *typesShim) Shim(shim typeShim) {
	if ts.shims == nil {
		ts.shims = make(map[string]*typeShim)
	}
	ts.shims[shim.Qual] = &shim
}

func (ts *typesShim) FindShim(qual string) *typeShim {
	return ts.shims[qual]
}
