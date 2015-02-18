package uaparser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func getTestMap(file string) []map[string]string {
	fmt.Printf("File : %s\n", file)
	testFile, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("error reading file: %s\n", err)
		return nil
	}

	testMap := make(map[string][]map[string]string)
	err = yaml.Unmarshal(testFile, &testMap)
	if err != nil {
		fmt.Printf("error parsing yaml: %s\n", err)
		return nil
	}

	return testMap["test_cases"]
}

var uapCoreRoot string = "../uap-core"
var dvcDefaultRegexFile string = uapCoreRoot + "/regexes.yaml"
var dvcParser *Parser = nil

func dvcInitParser(regexFile string) {
	if dvcParser == nil {
		dvcParser, _ = New(regexFile)
	}
}

func deviceTest(c *Client, test map[string]string) bool {
	if c.Device.Family != test["family"] {
		fmt.Printf("Expected: %v\nActual: %v\n", test, c.Device)
		return false
	}
	return true
}

func testHelper(file string, test_fn func(*Client, map[string]string) bool) bool {
	dvcInitParser(dvcDefaultRegexFile)

	tests := getTestMap(file)
	if tests == nil || len(tests) == 0 {
		fmt.Println("FAIL\nCouldn't open test file or no tests found")
		return false
	}

	for _, test := range tests {

		// Other language ports of ua_parser skips js_ua in testing
		if test["js_ua"] != "" {
			continue
		}

		testingString := test["user_agent_string"]
		client := dvcParser.Parse(testingString)

		if !test_fn(client, test) {
			fmt.Println("FAIL")
			return false
		}
	}
	return true
}

func TestDevice(t *testing.T) {
	if !testHelper(uapCoreRoot+"/tests/test_device.yaml", deviceTest) {
		t.Fail()
	} else {
		fmt.Println("PASS")
	}
}
