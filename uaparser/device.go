package uaparser

import (
	"regexp"
	"strings"
)

type Device struct {
	Family string
	Brand  string
	Model  string
}

type DevicePattern struct {
	Regexp            *regexp.Regexp
	Regex             string
	RegexFlag         string
	BrandReplacement  string
	DeviceReplacement string
	ModelReplacement  string
}

func (dvcPattern *DevicePattern) Match(line string, dvc *Device) {
	matches := dvcPattern.Regexp.FindStringSubmatch(line)
	if len(matches) == 0 {
		return
	}
	groupCount := dvcPattern.Regexp.NumSubexp()

	if len(dvcPattern.DeviceReplacement) > 0 {
		dvc.Family = allMatchesReplacement(dvcPattern.DeviceReplacement, matches)
	} else if groupCount >= 1 {
		dvc.Family = matches[1]
	}
	dvc.Family = strings.TrimSpace(dvc.Family)
	if len(dvcPattern.BrandReplacement) > 0 {
		if strings.Contains(dvcPattern.BrandReplacement, "$") {
			dvc.Brand = allMatchesReplacement(dvcPattern.BrandReplacement, matches)
		} else {
			dvc.Brand = dvcPattern.BrandReplacement
		}
	}
	dvc.Brand = strings.TrimSpace(dvc.Brand)
	if len(dvcPattern.ModelReplacement) > 0 {
		dvc.Model = allMatchesReplacement(dvcPattern.ModelReplacement, matches)
	} else if groupCount >= 1 {
		dvc.Model = matches[1]
	}
	dvc.Model = strings.TrimSpace(dvc.Model)
}

func (dvc *Device) ToString() string {
	return dvc.Family
}
