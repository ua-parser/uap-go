package uaparser

type Os struct {
	Family     string
	Major      string
	Minor      string
	Patch      string
	PatchMinor string
}

func (parser *osParser) Match(line string, os *Os) {
	matches := parser.Reg.FindStringSubmatch(line)
	if len(matches) > 0 {
		groupCount := parser.Reg.NumSubexp()

		if len(parser.OSReplacement) > 0 {
			os.Family = singleMatchReplacement(parser.OSReplacement, matches, 1)
		} else if groupCount >= 1 {
			os.Family = matches[1]
		}

		if len(parser.V1Replacement) > 0 {
			os.Major = singleMatchReplacement(parser.V1Replacement, matches, 2)
		} else if groupCount >= 2 {
			os.Major = matches[2]
		}

		if len(parser.V2Replacement) > 0 {
			os.Minor = singleMatchReplacement(parser.V2Replacement, matches, 3)
		} else if groupCount >= 3 {
			os.Minor = matches[3]
		}

		if len(parser.V3Replacement) > 0 {
			os.Patch = singleMatchReplacement(parser.V3Replacement, matches, 4)
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
