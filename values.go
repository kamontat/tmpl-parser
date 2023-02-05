package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtemplates"
)

func NewValues() *Values {
	return &Values{
		values: make(map[string]*Value),
		logger: logger.Get("values"),
	}
}

type Value struct {
	Name   string
	Input  string
	Output string
}

type Values struct {
	values map[string]*Value
	logger *logger.Logger
}

func (v *Values) Add(value *Value) {
	v.values[value.Name] = value
}

func (v *Values) Load(m mapper.Mapper, variable mapper.Mapper) error {
	var name, in, out, err = GetValues(m, variable)
	if err != nil {
		return err
	}

	var isInFile = IsTmpl(in)
	var isOutFile = IsFile(out)

	v.logger.Debug("Is input file: %t", isInFile)
	v.logger.Debug("Is output file: %t", isOutFile)

	if isInFile && isOutFile {
		v.Add(&Value{
			Name:   name,
			Input:  in,
			Output: out,
		})
		return nil
	}

	if isInFile && !isOutFile {
		var outFile = GetOutFile(in)
		v.Add(&Value{
			Name:   name,
			Input:  in,
			Output: JoinPath(out, outFile),
		})
		return nil
	}

	if !isInFile && !isOutFile {
		var ins = GetInFiles(in)
		for _, inFile := range ins {
			var outFile = GetOutFile(inFile)
			v.Add(&Value{
				Name:   JoinPath(name, inFile),
				Input:  inFile,
				Output: JoinPath(out, outFile),
			})
		}

		return nil
	}

	return fmt.Errorf(
		"cannot parse input %s and output %s",
		in, out,
	)
}

func (v *Values) Parse(variable mapper.Mapper) error {
	for key, value := range v.values {
		v.logger.Debug("%#v\n", value)

		// Load template
		var output, err = xtemplates.
			File(key, value.Input)
		if err != nil {
			return err
		}

		// Create output file
		outFile, err := CreateOutFile(value.Output)
		if err != nil {
			return err
		}

		// Parse template
		err = output.ExecuteTemplate(
			outFile,
			filepath.Base(value.Input),
			variable,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *Values) String() string {
	var builder = strings.Builder{}

	for _, value := range v.values {
		builder.WriteString(fmt.Sprintf("- %s:\n", value.Name))
		builder.WriteString(fmt.Sprintf("  Input: %s\n", value.Input))
		builder.WriteString(fmt.Sprintf("  Output: %s\n", value.Output))
	}

	return builder.String()
}
