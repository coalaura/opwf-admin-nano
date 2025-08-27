package main

import "regexp"

var (
	licenseRgx = regexp.MustCompile(`(?m)^license:[a-f0-9]{40}$`)
	integerRgx = regexp.MustCompile(`(?m)^\d+$`)
)

func IsValidLicense(str string) bool {
	return licenseRgx.MatchString(str)
}

func IsValidInteger(str string) bool {
	return integerRgx.MatchString(str)
}
