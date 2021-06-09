package layout

import (
	"couture/internal/pkg/couture"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Default ...
const Default = "default"

// Registry is the registry of layout names to their structs.
var Registry = map[string]Layout{
	Default: mustLoad(Default),
}

type (
	padding struct {
		Left  uint `yaml:"left,omitempty"`
		Right uint `yaml:"right,omitempty"`
	}

	// ColumnLayout ...
	ColumnLayout struct {
		Align   string  `yaml:"align,omitempty"`
		Width   uint    `yaml:"width,omitempty"`
		Padding padding `yaml:"padding,omitempty"`
		Sigil   string  `yaml:"sigil,omitempty"`
	}

	// Layout ...
	Layout struct {
		Application ColumnLayout `yaml:"application"`
		Caller      ColumnLayout `yaml:"caller"`
		Level       ColumnLayout `yaml:"level"`
		Message     ColumnLayout `yaml:"message"`
		Source      ColumnLayout `yaml:"source"`
		Context     ColumnLayout `yaml:"context"`
		Timestamp   ColumnLayout `yaml:"timestamp"`
	}
)

func load(name string) (*Layout, error) {
	f, err := couture.Open("/layouts/" + name + ".yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var layout Layout
	err = yaml.Unmarshal(b, &layout)
	if err != nil {
		return nil, err
	}
	return &layout, nil
}

func mustLoad(name string) Layout {
	layout, err := load(name)
	if err != nil {
		panic(err)
	}
	return *layout
}
