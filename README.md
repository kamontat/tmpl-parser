# Usage

```
$ tmpl [-config=<path>] [-debug] value [values...]

value syntax: [<name>:]<template_path>[=<output_path>]
template syntax: <file_name>.<file_extension>.tmpl
	- /tmp/template.yaml.tmpl
	- /tmp/template.yaml.tmpl=/tmp/output.yaml
	- ./template.yaml.tmpl=./output.yaml
	- template:./template.yaml.tmpl
	- template:./template.yaml.tmpl=./output.yaml

By default if no output specify, it will output to template directory.

  -cwd string
    	current directory for relative path resolve to (default "/Users/natcha/Desktop/tmpl")
  -data value
    	data files, either yaml or json (you can pass more than 1 times)
  -debug
    	enable debug information
  -raw value
    	raw data in format <key>=<value> (you can pass more than 1 times)
```

## Development

```bash
go run github.com/kamontat/tmpl/cli
```
