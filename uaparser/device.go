package uaparser

import "strings"

type Device struct {
	Family string
	Brand  string
	Model  string
}

func (parser *deviceParser) Match(line string, dvc *Device) {
	matches := parser.Reg.FindStringSubmatch(line)
	if len(matches) == 0 {
		return
	}
	groupCount := parser.Reg.NumSubexp()

	if len(parser.DeviceReplacement) > 0 {
		dvc.Family = allMatchesReplacement(parser.DeviceReplacement, matches)
	} else if groupCount >= 1 {
		dvc.Family = matches[1]
	}
	dvc.Family = strings.TrimSpace(dvc.Family)
	if len(parser.BrandReplacement) > 0 {
		if strings.Contains(parser.BrandReplacement, "$") {
			dvc.Brand = allMatchesReplacement(parser.BrandReplacement, matches)
		} else {
			dvc.Brand = parser.BrandReplacement
		}
	}
	dvc.Brand = strings.TrimSpace(dvc.Brand)
	if len(parser.ModelReplacement) > 0 {
		dvc.Model = allMatchesReplacement(parser.ModelReplacement, matches)
	} else if groupCount >= 1 {
		dvc.Model = matches[1]
	}
	dvc.Model = strings.TrimSpace(dvc.Model)
}

func (dvc *Device) ToString() string {
	return dvc.Family
}
