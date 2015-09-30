package uaparser

import (
	"regexp"
	"strings"
)

type Device struct {
	Family string
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
}

func (dvc *Device) ToString() string {
	return dvc.Family
}
