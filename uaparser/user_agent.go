package uaparser

type UserAgent struct {
	Family string
	Major  string
	Minor  string
	Patch  string
}

func (parser *uaParser) Match(line string, ua *UserAgent) {
	matches := parser.Reg.FindStringSubmatch(line)
	if len(matches) > 0 {
		groupCount := parser.Reg.NumSubexp()

		if len(parser.FamilyReplacement) > 0 {
			ua.Family = singleMatchReplacement(parser.FamilyReplacement, matches, 1)
		} else if groupCount >= 1 {
			ua.Family = matches[1]
		}

		if len(parser.V1Replacement) > 0 {
			ua.Major = singleMatchReplacement(parser.V1Replacement, matches, 2)
		} else if groupCount >= 2 {
			ua.Major = matches[2]
		}

		if len(parser.V2Replacement) > 0 {
			ua.Minor = singleMatchReplacement(parser.V2Replacement, matches, 3)
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
