package utils

import (
	"strings"
)

func AnswerTrim(str string) string {
	startIndex := strings.Index(str, "</think>")
	if startIndex == -1 {
		return str
	}

	return strings.TrimSpace(str[startIndex:])
}
