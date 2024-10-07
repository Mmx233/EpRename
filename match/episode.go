package match

import (
	"fmt"
	"regexp"
)

func GetEpisodeSerial(filename string) (string, error) {
	reg := regexp.MustCompile(`(\[([0-9]+)])`)
	match := reg.FindStringSubmatch(filename)
	if len(match) == 0 {
		return "", fmt.Errorf("match episode serial failed: %s", filename)
	}
	return match[2], nil
}

func GetEpisodeSerialWithRegexp(reg *regexp.Regexp, filename string) (string, error) {
	match := reg.FindStringSubmatch(filename)
	if len(match) == 0 {
		return "", fmt.Errorf("match episode serial failed: %s", filename)
	}
	return match[1], nil
}
