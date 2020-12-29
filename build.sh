#!/usr/bin/env bash

git submodule init; git submodule update

# Strip out empty lines and comments for conciseness:
yaml=`cat uap-core/regexes.yaml | sed '/\s*#/d' | sed '/^\s*$/d'`

# Build and format a Go file including our sources:
echo "package uaparser
var DefinitionYaml = []byte(\`$yaml\`)" | gofmt > uaparser/yaml.go

rm -rf top-user-agents; git clone https://github.com/Kikobeats/top-user-agents --depth 1
yaml=`go run cmd/optimize/main.go < top-user-agents/index.json`

# Rebuild the sorted ua list
echo "package uaparser
var DefinitionYaml = []byte(\`$yaml\`)" | gofmt > uaparser/yaml.go
