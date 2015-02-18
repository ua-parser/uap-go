package uaparser

import (
	"fmt"
	"testing"
)

var osDefaultRegexFile string = uapCoreRoot + "/regexes.yaml"
var osParser *Parser = nil

func osInitParser(regexFile string) {
	if osParser == nil {
		osParser, _ = New(regexFile)
	}
}

func osTest(c *Client, test map[string]string) bool {
	os := c.Os
	if os.Family != test["family"] || os.Major != test["major"] ||
		os.Minor != test["minor"] || os.Patch != test["patch"] ||
		os.PatchMinor != test["patch_minor"] {
		fmt.Printf("Expected: %v\nActual: %v\n", test, os)
		return false
	}
	return true
}

func TestOs(t *testing.T) {
	if !testHelper(uapCoreRoot+"/tests/test_os.yaml", osTest) {
		t.Fail()
	} else {
		fmt.Println("PASS")
	}
}

func TestAdditionalOs(t *testing.T) {
	if !testHelper(uapCoreRoot+"/test_resources/additional_os_tests.yaml", osTest) {
		t.Fail()
	} else {
		fmt.Println("PASS")
	}
}
