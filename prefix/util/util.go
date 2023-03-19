package util

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func GetMD5(inputString string) string {
	// returns md5 string of an input string
	return fmt.Sprintf("%x", md5.Sum([]byte(inputString)))
}

func AppendCSV(originalStr string, newStr string) string {
	return originalStr + ", " + newStr
}

func Csv2List(input string) ([]string, bool) {
	// checks if the input string is comma seperated or not.
	// if yes, then returns the individual values as a list and true
	// else returns the single value and false
	values := strings.Split(input, ", ")
	if len(values) > 1 {
		return values, true
	}
	return values, false
}

func IsSubstring(segment string, text string) bool {
	// this function checks if text starts with segment
	for i, char := range segment {
		if rune(text[i]) != char {
			return false
		}
	}
	return true
}
