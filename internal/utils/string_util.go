package utils

import (
	"strings"
)

func TrimAll(s string) (trimmed string) {
	trimmed = s
	trimmed = strings.Replace(trimmed, "\r", "", -1)
	trimmed = strings.Replace(trimmed, "\n", "", -1)
	trimmed = strings.Replace(trimmed, "\t", "", -1)
	trimmed = strings.Replace(trimmed, " ", "", -1)

	return trimmed
}
