package anagrams

import (
	"regexp"
	"sort"
	"strings"
)

func Anagrams(s1 string, s2 string) bool {
	// Cleanup the strings, remove spaces and special chars
	re := regexp.MustCompile(`[^a-zA-Z]`)
	cs1 := re.ReplaceAllString(s1, "")
	cs2 := re.ReplaceAllString(s2, "")
	// Lowercase the strings
	cs1 = strings.ToLower(cs1)
	cs2 = strings.ToLower((cs2))
	// Sort the strings by chars
	css1 := strings.Split(cs1, "")
	css2 := strings.Split(cs2, "")
	sort.Strings(css1)
	sort.Strings(css2)

	// Compare strings
	return strings.Join(css1, "") == strings.Join(css2, "")
}
