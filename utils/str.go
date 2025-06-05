package utils

import (
	"strings"
)

func AnswerTrim(str string) string {
	strSlice := strings.Split(str, "</think>")
	if len(strSlice) != 2 {
		return str
	}

	return strings.ReplaceAll(strSlice[1], "\n", "")
}
