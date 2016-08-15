#!/usr/bin/env bash

git submodule update

# Strip out empty lines and comments for conciseness:
yaml=`cat uap-core/regexes.yaml | sed '/\s*#/d' | sed '/^\s*$/d'`

# Build and format a Go file including our sources:
echo "package uaparser
var definitionYaml = []byte(\`$yaml\`)" | gofmt > uaparser/yaml.go
