#!/usr/bin/env bash

git submodule update --recursive --remote

# Strip out empty lines and comments for conciseness:
yaml=`cat uap-core/regexes.yaml | sed '/\s*#/d' | sed '/^\s*$/d'`

# Build and format a Go file including our sources:
echo "package uaparser
var DefinitionYaml = []byte(\`$yaml\`)" | gofmt > uaparser/yaml.go
