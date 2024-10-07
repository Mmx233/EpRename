package match

import "strings"

func Ext(path string) string {
	split := strings.Split(path, ".")
	if len(split) == 1 {
		return ""
	}
	return "." + strings.Join(split[1:], ".")
}
