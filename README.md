# Usage

```
Usage of tmpl-parser:

Value Syntax:
  [<name>:]<template_path>[=<output_path>]
  - name: name of the template
      any string are accepted.
      default to template_path if not provided.
  - template: template file
      relative or absolute are accepted.
      relative will resolve by using --cwd option.
      must have .tmpl extension; otherwise will throw error.
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

Commands:
  - default    : load and parse template to output
  - help       : show application help
  - version    : show current application version
  - config     : list all possible config user can set
    --data         <bool:false> 
       show config value as well
    --all          <bool:false> 
       show all configuration, including internal

Options:
  --secrets      <[]>                                               
     [-s] add secrets to templates
  --help         <bool:false>                                       
     [-h] show application help
  --version      <bool:false>                                       
     [-v] show current application version
  --log-level    <int:3>                                            
     [-l] setup log level; 0 is silent and 4 is verbose
  --debug        <bool:false>                                       
     [-D] mark current log to debug mode
  --pwd          <str:/Users/kamontat/Desktop/Personal/tmpl-parser> 
     current directory
  --configs      <[{{.current}}/configs]>                           
     configuration file/directory. both must be either json or yaml
  --envs         <[]>                                               
     environment file/directory. each file must following .env regulation
  --no-env-file  <bool:false>                                       
     disabled loading .env files completely
  --var          <[]>                                               
     add data to variables config
```

## Development

```bash
go run github.com/kamontat/tmpl/cli
```
