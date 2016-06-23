package uaparser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"
)

type RegexesDefinitions struct {
	UA     []*uaParser     `yaml:"user_agent_parsers"`
	OS     []*osParser     `yaml:"os_parsers"`
	Device []*deviceParser `yaml:"device_parsers"`
}

type uaParser struct {
	Reg               *regexp.Regexp
	Expr              string `yaml:"regex"`
	Flags             string `yaml:"regex_flag"`
	FamilyReplacement string `yaml:"family_replacement"`
	V1Replacement     string `yaml:"v1_replacement"`
	V2Replacement     string `yaml:"v2_replacement"`
	V3Replacement     string `yaml:"v3_replacement"`
}

type osParser struct {
	Reg           *regexp.Regexp
	Expr          string `yaml:"regex"`
	Flags         string `yaml:"regex_flag"`
	OSReplacement string `yaml:"os_replacement"`
	V1Replacement string `yaml:"os_v1_replacement"`
	V2Replacement string `yaml:"os_v2_replacement"`
	V3Replacement string `yaml:"os_v3_replacement"`
	V4Replacement string `yaml:"os_v4_replacement"`
	V5Replacement string `yaml:"os_v5_replacement"`
}

type deviceParser struct {
	Reg               *regexp.Regexp
	Expr              string `yaml:"regex"`
	Flags             string `yaml:"regex_flag"`
	DeviceReplacement string `yaml:"device_replacement"`
	BrandReplacement  string `yaml:"brand_replacement"`
	ModelReplacement  string `yaml:"model_replacement"`
}

type Parser struct {
	RegexesDefinitions
}

func (parser *Parser) mustCompile() { // until we can use yaml.UnmarshalYAML with embedded pointer struct
	for _, p := range parser.UA {
		p.Reg = compileRegex(p.Flags, p.Expr)
	}
	for _, p := range parser.OS {
		p.Reg = compileRegex(p.Flags, p.Expr)
	}
	for _, p := range parser.Device {
		p.Reg = compileRegex(p.Flags, p.Expr)
	}
}

func compileRegex(flags, expr string) *regexp.Regexp {
	if flags == "" {
		return regexp.MustCompile(expr)
	} else {
		return regexp.MustCompile(fmt.Sprintf("(?%s)%s", flags, expr))
	}
}

type Client struct {
	UserAgent *UserAgent
	Os        *Os
	Device    *Device
}

func New(regexFile string) (*Parser, error) {
	data, err := ioutil.ReadFile(regexFile)
	if nil != err {
		return nil, err
	}
	return NewFromBytes(data)
}

func NewFromBytes(data []byte) (*Parser, error) {
	var definitions RegexesDefinitions
	if err := yaml.Unmarshal(data, &definitions); err != nil {
		return nil, err
	}

	parser := &Parser{definitions}
	parser.mustCompile()

	return parser, nil
}

func (parser *Parser) ParseUserAgent(line string) *UserAgent {
	ua := new(UserAgent)
	found := false
	for _, uaPattern := range parser.UA {
		uaPattern.Match(line, ua)
		if len(ua.Family) > 0 {
			found = true
			break
		}
	}
	if !found {
		ua.Family = "Other"
	}
	return ua
}

func (parser *Parser) ParseOs(line string) *Os {
	os := new(Os)
	found := false
	for _, osPattern := range parser.OS {
		osPattern.Match(line, os)
		if len(os.Family) > 0 {
			found = true
			break
		}
	}
	if !found {
		os.Family = "Other"
	}
	return os
}

func (parser *Parser) ParseDevice(line string) *Device {
	dvc := new(Device)
	found := false
	for _, dvcPattern := range parser.Device {
		dvcPattern.Match(line, dvc)
		if len(dvc.Family) > 0 {
			found = true
			break
		}
	}
	if !found {
		dvc.Family = "Other"
	}
	return dvc
}

func (parser *Parser) Parse(line string) *Client {
	cli := new(Client)
	cli.UserAgent = parser.ParseUserAgent(line)
	cli.Os = parser.ParseOs(line)
	cli.Device = parser.ParseDevice(line)
	return cli
}

func singleMatchReplacement(replacement string, matches []string, idx int) string {
	token := "$" + strconv.Itoa(idx)
	if strings.Contains(replacement, token) {
		return strings.Replace(replacement, token, matches[idx], -1)
	}
	return replacement
}

// allMatchesReplacement replaces all tokens in format $<digit> (like $1 or $12) with values
// at corresponding indexes (NOT POSITIONS, so $1 will be replaced with v[1], NOT v[0]) in the provided array.
// If array doesn't have value at the index (when array length is less than the value), it remains unchanged in the string
func allMatchesReplacement(pattern string, matches []string) string {
	var output bytes.Buffer
	readingToken := false
	var readToken bytes.Buffer
	writeTokenValue := func() {
		if !readingToken {
			return
		}
		if readToken.Len() == 0 {
			output.WriteRune('$')
			return
		}
		idx, err := strconv.Atoi(readToken.String())
		// index is out of range when value is too big for int or when it's zero (or less) or greater than array length
		indexOutOfRange := (err != nil && err.(*strconv.NumError).Err != strconv.ErrRange) || idx <= 0 || idx >= len(matches)
		if indexOutOfRange {
			output.WriteRune('$')
			output.Write(readToken.Bytes())
			readToken.Reset()
			return
		}
		if err != nil {
			// should never happen
			panic(err)
		}
		output.WriteString(matches[idx])
		readToken.Reset()
	}
	for _, r := range pattern {
		if !readingToken && r == '$' {
			readingToken = true
			continue
		}
		if !readingToken {
			output.WriteRune(r)
			continue
		}
		if unicode.IsDigit(r) {
			readToken.WriteRune(r)
			continue
		}
		writeTokenValue()
		readingToken = (r == '$')
		if !readingToken {
			output.WriteRune(r)
		}
	}
	writeTokenValue()
	return output.String()
}
