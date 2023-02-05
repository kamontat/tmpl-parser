package main

import (
	"errors"

	"github.com/kc-workspace/go-lib/mapper"
)

func GetValues(m, variable mapper.Mapper) (string, string, string, error) {
	var in, err = m.Se("input")
	if err != nil {
		return "", "", "", errors.New("input field is requires")
	}

	in = ParseTemplate(in, variable)

	var name = in
	if m.Has("name") {
		name = m.Si("name")
	}

	var out = in
	if m.Has("output") {
		out = m.Si("output")
		out = ParseTemplate(out, variable)
	}

	return name, in, out, nil
}
