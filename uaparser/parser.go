package uaparser

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sync"

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

func (ua *uaParser) setDefaults() {
	if ua.FamilyReplacement == "" {
		ua.FamilyReplacement = "$1"
	}
	if ua.V1Replacement == "" {
		ua.V1Replacement = "$2"
	}
	if ua.V2Replacement == "" {
		ua.V2Replacement = "$3"
	}
	if ua.V3Replacement == "" {
		ua.V3Replacement = "$4"
	}
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
}

func (os *osParser) setDefaults() {
	if os.OSReplacement == "" {
		os.OSReplacement = "$1"
	}
	if os.V1Replacement == "" {
		os.V1Replacement = "$2"
	}
	if os.V2Replacement == "" {
		os.V2Replacement = "$3"
	}
	if os.V3Replacement == "" {
		os.V3Replacement = "$4"
	}
	if os.V4Replacement == "" {
		os.V4Replacement = "$5"
	}
}

type deviceParser struct {
	Reg               *regexp.Regexp
	Expr              string `yaml:"regex"`
	Flags             string `yaml:"regex_flag"`
	DeviceReplacement string `yaml:"device_replacement"`
	BrandReplacement  string `yaml:"brand_replacement"`
	ModelReplacement  string `yaml:"model_replacement"`
}

func (device *deviceParser) setDefaults() {
	if device.DeviceReplacement == "" {
		device.DeviceReplacement = "$1"
	}
	if device.ModelReplacement == "" {
		device.ModelReplacement = "$1"
	}
}

type Client struct {
	UserAgent *UserAgent
	Os        *Os
	Device    *Device
}

type Parser struct {
	RegexesDefinitions
}

func (parser *Parser) mustCompile() { // until we can use yaml.UnmarshalYAML with embedded pointer struct
	for _, p := range parser.UA {
		p.Reg = compileRegex(p.Flags, p.Expr)
		p.setDefaults()
	}
	for _, p := range parser.OS {
		p.Reg = compileRegex(p.Flags, p.Expr)
		p.setDefaults()
	}
	for _, p := range parser.Device {
		p.Reg = compileRegex(p.Flags, p.Expr)
		p.setDefaults()
	}
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

func (parser *Parser) Parse(line string) *Client {
	cli := new(Client)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		cli.UserAgent = parser.ParseUserAgent(line)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cli.Os = parser.ParseOs(line)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cli.Device = parser.ParseDevice(line)
	}()

	wg.Wait()

	return cli
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

func compileRegex(flags, expr string) *regexp.Regexp {
	if flags == "" {
		return regexp.MustCompile(expr)
	} else {
		return regexp.MustCompile(fmt.Sprintf("(?%s)%s", flags, expr))
	}
}
