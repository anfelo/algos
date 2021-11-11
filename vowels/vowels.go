package vowels

import (
	"regexp"
	"strings"
)

func Vowels(s string) int {
	re := regexp.MustCompile(`[^aeiouAEIOU]`)
	v := re.ReplaceAllString(s, "")

	return len(v)
}

func Vowels1(s string) int {
	vv := []string{"a", "e", "i", "o", "u"}
	counter := 0
	for _, v := range strings.Split(s, "") {
		if contains(vv, v) {
			counter++
		}
	}
	return counter
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == strings.ToLower(e) {
			return true
		}
	}
	return false
}
