package match

import (
	"errors"
	"regexp"
	"strings"
)

func GetSeasonSerial(pathname string) (string, error) {
	pathname = strings.ReplaceAll(pathname, `\`, `/`)
	pathname = strings.ToLower(pathname)

	reg := regexp.MustCompile(`^(season\s([0-9]+))|(s([0-9]+))$`)
	splitPath := strings.Split(pathname, "/")
	for i := len(splitPath) - 1; i >= 0; i-- {
		match := reg.FindStringSubmatch(splitPath[i])
		if len(match) > 0 {
			return match[2], nil
		}
	}
	return "", errors.New("match season serial failed")
}
