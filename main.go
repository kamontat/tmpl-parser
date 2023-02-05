package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kc-workspace/go-lib/commandline"
	"github.com/kc-workspace/go-lib/commandline/commands"
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/commandline/models"
	"github.com/kc-workspace/go-lib/commandline/plugins"
	"github.com/kc-workspace/go-lib/configs"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

// assign from goreleaser
var (
	short   string = "tp"
	name    string = "tmpl-parser"
	version string = "dev"
	commit  string = "none"
	date    string = "unknown"
	builtBy string = "manually"
)

func main() {
	var err = commandline.New(&models.Metadata{
		Name:    name,
		Version: version,
		Commit:  commit,
		Date:    date,
		BuiltBy: builtBy,
		Usage: fmt.Sprintf(`
Value Syntax:
  [<name>:]<template_path>[=<output_path>]
  - name: name of the template
      any string are accepted.
      default to template_path if not provided.
  - template: template file
      relative or absolute are accepted.
      relative will resolve by using --cwd option.
      must have %s extension; otherwise will throw error.
output:
	relative or absolute are accepted.
	directory or filename are accepted as well.
	default to template file without template extension,
	if only directory provided, use template file as output file.

Example
  '/tmp/dir/gh.txt.tmpl'
	name: /tmp/dir/gh.txt.tmpl
	template: /tmp/dir/gh.txt.tmpl
	output: /tmp/dir/gh.txt
  '/tmp/content.md.tmpl=/tmp/readme.txt'
	name: /tmp/content.md.tmpl
	template: /tmp/content.md.tmpl
	output: /tmp/readme.txt
  './file.json.tmpl=./untitled.json'
	name: ./file.json.tmpl
	template: /$CWD/file.json.tmpl
	output: /$CWD/untitled.json
  'default:config.yaml.tmpl'
	name: default
	template: /$CWD/config.yaml.tmpl
	output: /$CWD/config.yaml
  'custom:values.yaml.tmpl=./output/values.yaml'
	name: custom
	template: /$CWD/values.yaml.tmpl
	output: /$CWD/output/values.yaml
  'custom:values.yaml.tmpl=output'
	name: custom
	template: /$CWD/values.yaml.tmpl
	output: /$CWD/output/values.yaml
`, TEMPLATE_EXTENSION),
	}).
		Plugin(plugins.SupportHelp()).
		Plugin(plugins.SupportVersion()).
		Plugin(plugins.SupportLogLevel(logger.INFO)).
		Plugin(plugins.SupportConfig(short, []string{"{{.current}}/configs"})).
		Plugin(plugins.SupportDotEnv(false)).
		Plugin(plugins.SupportVar()).
		Flag(flags.Array{
			Name:    "secrets",
			Aliases: []string{"s"},
			Default: []string{},
			Usage:   "add secrets to templates",
			Action: func(data []string) mapper.Mapper {
				var m = mapper.New()
				for _, d := range data {
					var key, value, ok = configs.ParseOverride(d)
					if ok {
						m.Set(fmt.Sprintf("secrets.%s", key), value)
					}
				}
				return m
			},
		}).
		Command(&commands.Command{
			Name:  commands.DEFAULT,
			Usage: `load and parse template to output`,
			Executor: func(p *commands.ExecutorParameter) error {
				var variable = p.Config.Mi("variables")
				var raw = p.Config.Ar("values")

				var values = NewValues()
				for i, r := range raw {
					if m, ok := mapper.ToMapper(r); ok {
						var err = values.Load(m, variable)
						if err != nil {
							p.Logger.Warn("error on loading value[%d]: %v", i, err)
							continue
						}
					}
				}

				p.Logger.Info("parsing %d values", len(values.values))
				return values.Parse(p.Config)
			},
		}).
		Start(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
