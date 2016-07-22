package uaparser

import (
	"io/ioutil"
	"log"
	"testing"

	"gopkg.in/yaml.v2"
)

var testParser *Parser

func init() {
	var err error
	testParser, err = New("../uap-core/regexes.yaml")
	if err != nil {
		log.Fatal(err)
	}
}

func TestOSParsing(t *testing.T) {
	var YAMLTestCases struct {
		Cases []struct {
			Input    string `yaml:"user_agent_string"`
			Expected Os     `yaml:",inline"`
		} `yaml:"test_cases"`
	}

	for _, filepath := range []string{"../uap-core/test_resources/additional_os_tests.yaml", "../uap-core/tests/test_os.yaml"} {
		unmarshalResourceTestFile(filepath, &YAMLTestCases)

		for _, c := range YAMLTestCases.Cases {
			actual := testParser.ParseOs(c.Input)
			if got, want := *actual, c.Expected; got != want {
				t.Fatalf("got %#v\n want %#v\n", got, want)
			}
		}
	}
}

func TestReadsInteralYAML(t *testing.T) {
	_ = NewFromSaved() // should not panic
}

func TestUAParsing(t *testing.T) {
	t.Parallel()

	var YAMLTestCases struct {
		Cases []struct {
			Input    string    `yaml:"user_agent_string"`
			Expected UserAgent `yaml:",inline"`
		} `yaml:"test_cases"`
	}

	for _, filepath := range []string{"../uap-core/test_resources/pgts_browser_list.yaml", "../uap-core/test_resources/firefox_user_agent_strings.yaml", "../uap-core/tests/test_ua.yaml"} {
		unmarshalResourceTestFile(filepath, &YAMLTestCases)

		for _, c := range YAMLTestCases.Cases {
			actual := testParser.ParseUserAgent(c.Input)
			if got, want := *actual, c.Expected; got != want {
				t.Fatalf("got %#v\n want %#v\n", got, want)
			}
		}
	}
}

func TestDeviceParsing(t *testing.T) {
	t.Parallel()

	var YAMLTestCases struct {
		Cases []struct {
			Input    string `yaml:"user_agent_string"`
			Expected Device `yaml:",inline"`
		} `yaml:"test_cases"`
	}

	for _, filepath := range []string{"../uap-core/tests/test_device.yaml"} {
		unmarshalResourceTestFile(filepath, &YAMLTestCases)

		for _, c := range YAMLTestCases.Cases {
			actual := testParser.ParseDevice(c.Input)
			if got, want := *actual, c.Expected; got != want {
				t.Fatalf("got %#v\n want %#v\n", got, want)
			}
		}
	}
}

func TestGenericParseMethodConcurrency(t *testing.T) { // go test -race -run=Concurrency
	var YAMLTestCases struct {
		Cases []struct {
			Input    string `yaml:"user_agent_string"`
			Expected Os     `yaml:",inline"`
		} `yaml:"test_cases"`
	}

	for _, filepath := range []string{"../uap-core/tests/test_os.yaml"} {
		unmarshalResourceTestFile(filepath, &YAMLTestCases)

		for _, c := range YAMLTestCases.Cases {
			actual := testParser.Parse(c.Input).Os
			if got, want := *actual, c.Expected; got != want {
				t.Fatalf("got %#v\n want %#v\n", got, want)
			}
		}
	}
}

func unmarshalResourceTestFile(filepath string, v interface{}) {
	file, err := ioutil.ReadFile(filepath)
	if nil != err {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(file, v); err != nil {
		log.Fatal(err)
	}
}
