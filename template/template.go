package template

import (
	"html/template"

	"github.com/mitchellh/mapstructure"
)

type TemplateConfig struct {
	Path string
}

func Creator(cfg interface{}) (interface{}, error) {
	var tc TemplateConfig
	err := mapstructure.WeakDecode(cfg, &tc)
	if err != nil {
		return nil, err
	}
	tpl, err := template.ParseGlob(tc.Path)
	if err != nil {
		return nil, err
	}
	return tpl, nil
}
