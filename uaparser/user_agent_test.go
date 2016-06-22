package uaparser

import (
	"fmt"
	"testing"
)

var uaDefaultRegexFile string = uapCoreRoot + "/regexes.yaml"
var uaTestParser *Parser = nil

func uaInitParser(regexFile string) {
	if uaTestParser == nil {
		uaTestParser, _ = New(regexFile)
	}
}

func uaTest(c *Client, test map[string]string) bool {
	ua := c.UserAgent
	if ua.Family != test["family"] || ua.Major != test["major"] ||
		ua.Minor != test["minor"] || ua.Patch != test["patch"] {
		fmt.Printf("Expected: %q\vActual: %#v\n", test, ua)
		return false
	}
	return true
}

func TestUserAgent(t *testing.T) {
	if !testHelper(uapCoreRoot+"/tests/test_ua.yaml", uaTest) {
		t.Fail()
	} else {
		fmt.Println("PASS")
	}
}

func TestFirefoxUserAgents(t *testing.T) {
	if !testHelper(uapCoreRoot+"/test_resources/firefox_user_agent_strings.yaml", uaTest) {
		t.Fail()
	} else {
		fmt.Println("PASS")
	}
}

func TestPgtsBrowsersList(t *testing.T) {
	if !testHelper(uapCoreRoot+"/test_resources/pgts_browser_list.yaml", uaTest) {
		t.Fail()
	} else {
		fmt.Println("PASS")
	}
}
