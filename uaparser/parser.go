package uaparser

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

type Parser struct {
	UserAgentPatterns []UserAgentPattern
	OsPatterns        []OsPattern
	DevicePatterns    []DevicePattern
}

type Client struct {
	UserAgent *UserAgent
	Os        *Os
	Device    *Device
}

var exportedNameRegex = regexp.MustCompile("[0-9A-Za-z]+")

func GetExportedName(src string) string {
	byteSrc := []byte(src)
	chunks := exportedNameRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}
	return string(bytes.Join(chunks, nil))
}

func ToStruct(interfaceArr []map[string]string, typeInterface interface{}, returnVal *[]interface{}) {
	structArr := make([]interface{}, 0)
	for _, interfaceMap := range interfaceArr {
		structValPtr := reflect.New(reflect.TypeOf(typeInterface))
		structVal := structValPtr.Elem()
		for key, value := range interfaceMap {
			structVal.FieldByName(GetExportedName(key)).SetString(value)
		}
		structArr = append(structArr, structVal.Interface())
	}
	*returnVal = structArr
}

func New(regexFile string) (*Parser, error) {
	parser := new(Parser)

	data, err := ioutil.ReadFile(regexFile)
	if nil != err {
		return nil, err
	}

	return parser.newFromBytes(data)
}

func NewFromBytes(regexBytes []byte) (*Parser, error) {
	parser := new(Parser)

	return parser.newFromBytes(regexBytes)
}

func (parser *Parser) newFromBytes(data []byte) (*Parser, error) {
	m := make(map[string][]map[string]string)
	err := yaml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	uaPatternType := new(UserAgentPattern)
	var uaInterfaces []interface{}
	var uaPatterns []UserAgentPattern

	wg.Add(1)
	go func() {
		ToStruct(m["user_agent_parsers"], *uaPatternType, &uaInterfaces)
		uaPatterns = make([]UserAgentPattern, len(uaInterfaces))
		for i, inter := range uaInterfaces {
			uaPatterns[i] = inter.(UserAgentPattern)
			uaPatterns[i].Regexp = regexp.MustCompile(uaPatterns[i].Regex)
		}
		wg.Done()
	}()

	osPatternType := new(OsPattern)
	var osInterfaces []interface{}
	var osPatterns []OsPattern

	wg.Add(1)
	go func() {
		ToStruct(m["os_parsers"], *osPatternType, &osInterfaces)
		osPatterns = make([]OsPattern, len(osInterfaces))
		for i, inter := range osInterfaces {
			osPatterns[i] = inter.(OsPattern)
			osPatterns[i].Regexp = regexp.MustCompile(osPatterns[i].Regex)
		}
		wg.Done()
	}()

	dvcPatternType := new(DevicePattern)
	var dvcInterfaces []interface{}
	var dvcPatterns []DevicePattern

	wg.Add(1)
	go func() {
		ToStruct(m["device_parsers"], *dvcPatternType, &dvcInterfaces)
		dvcPatterns = make([]DevicePattern, len(dvcInterfaces))
		for i, inter := range dvcInterfaces {
			dvcPatterns[i] = inter.(DevicePattern)
			flags := ""
			if strings.Contains(dvcPatterns[i].RegexFlag, "i") {
				flags = "(?i)"
			}
			regexString := fmt.Sprintf("%s%s", flags, dvcPatterns[i].Regex)
			dvcPatterns[i].Regexp = regexp.MustCompile(regexString)
		}
		wg.Done()
	}()

	wg.Wait()

	parser.UserAgentPatterns = uaPatterns
	parser.OsPatterns = osPatterns
	parser.DevicePatterns = dvcPatterns

	return parser, nil
}

func (parser *Parser) ParseUserAgent(line string) *UserAgent {
	ua := new(UserAgent)
	found := false
	for _, uaPattern := range parser.UserAgentPatterns {
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
	for _, osPattern := range parser.OsPatterns {
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
	for _, dvcPattern := range parser.DevicePatterns {
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
