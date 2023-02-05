package main

import (
	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtemplates"
)

func ParseTemplate(tmpl string, variable mapper.Mapper) string {
	var content, err = xtemplates.Text(tmpl, variable)
	if err != nil {
		return tmpl
	}

	return content
}
