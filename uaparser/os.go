package uaparser

import (
	"regexp"
)

type Os struct {
	Family     string
	Major      string
	Minor      string
	Patch      string
	PatchMinor string
}

type OsPattern struct {
	Regexp          *regexp.Regexp
	Regex           string
	OsReplacement   string
	OsV1Replacement string
	OsV2Replacement string
	OsV3Replacement string
}

func (osPattern *OsPattern) Match(line string, os *Os) {
	matches := osPattern.Regexp.FindStringSubmatch(line)
	if len(matches) > 0 {
		groupCount := osPattern.Regexp.NumSubexp()

		if len(osPattern.OsReplacement) > 0 {
			os.Family = singleMatchReplacement(osPattern.OsReplacement, matches, 1)
		} else if groupCount >= 1 {
			os.Family = matches[1]
		}

		if len(osPattern.OsV1Replacement) > 0 {
			os.Major = singleMatchReplacement(osPattern.OsV1Replacement, matches, 2)
		} else if groupCount >= 2 {
			os.Major = matches[2]
		}

		if len(osPattern.OsV2Replacement) > 0 {
			os.Minor = singleMatchReplacement(osPattern.OsV2Replacement, matches, 3)
		} else if groupCount >= 3 {
			os.Minor = matches[3]
		}

		if len(osPattern.OsV3Replacement) > 0 {
			os.Patch = singleMatchReplacement(osPattern.OsV3Replacement, matches, 4)
		} else if groupCount >= 4 {
			os.Patch = matches[4]
		}

		if groupCount >= 5 {
			os.PatchMinor = matches[5]
		}
	}
}

func (os *Os) ToString() string {
	var str string
	if os.Family != "" {
		str += os.Family
	}
	version := os.ToVersionString()
	if version != "" {
		str += " " + version
	}
	return str
}

func (os *Os) ToVersionString() string {
	var version string
	if os.Major != "" {
		version += os.Major
	}
	if os.Minor != "" {
		version += "." + os.Minor
	}
	if os.Patch != "" {
		version += "." + os.Patch
	}
	if os.PatchMinor != "" {
		version += "." + os.PatchMinor
	}
	return version
}
