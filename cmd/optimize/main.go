package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ua-parser/uap-go/uaparser"
	"gopkg.in/yaml.v2"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	must(err)

	var uas []string
	err = json.Unmarshal(data, &uas)
	must(err)

	parser := uaparser.NewFromSaved()
	parser.UseSort = true

	for _, ua := range uas {
		// Hack to ensure we actually sort...
		parser.UserAgentMisses = 10000000
		parser.Parse(ua)
	}

	data, err = yaml.Marshal(parser.RegexesDefinitions)
	must(err)

	_, err = os.Stdout.Write(data)
	must(err)
}
