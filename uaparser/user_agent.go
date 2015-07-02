package uaparser

import (
	"regexp"
)

type UserAgent struct {
	Family string
	Major  string
	Minor  string
	Patch  string
}

type UserAgentPattern struct {
	Regexp            *regexp.Regexp
	Regex             string
	RegexFlag         string
	FamilyReplacement string
	V1Replacement     string
	V2Replacement     string
}

func (uaPattern *UserAgentPattern) Match(line string, ua *UserAgent) {
	matches := uaPattern.Regexp.FindStringSubmatch(line)
	if len(matches) > 0 {
		groupCount := uaPattern.Regexp.NumSubexp()

		if len(uaPattern.FamilyReplacement) > 0 {
			ua.Family = singleMatchReplacement(uaPattern.FamilyReplacement, matches, 1)
		} else if groupCount >= 1 {
			ua.Family = matches[1]
		}

		if len(uaPattern.V1Replacement) > 0 {
			ua.Major = singleMatchReplacement(uaPattern.V1Replacement, matches, 2)
		} else if groupCount >= 2 {
			ua.Major = matches[2]
		}

		if len(uaPattern.V2Replacement) > 0 {
			ua.Minor = singleMatchReplacement(uaPattern.V2Replacement, matches, 3)
		} else if groupCount >= 3 {
			ua.Minor = matches[3]
			if groupCount >= 4 {
				ua.Patch = matches[4]
			}
		}
	}
}

func (ua *UserAgent) ToString() string {
	var str string
	if ua.Family != "" {
		str += ua.Family
	}
	version := ua.ToVersionString()
	if version != "" {
		str += " " + version
	}
	return str
}

func (ua *UserAgent) ToVersionString() string {
	var version string
	if ua.Major != "" {
		version += ua.Major
	}
	if ua.Minor != "" {
		version += "." + ua.Minor
	}
	if ua.Patch != "" {
		version += "." + ua.Patch
	}
	return version
}
