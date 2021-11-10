# Directions

Check to see if two provided strings are anagrams of each other. One string is an anagram of another if it uses the same characters in the same quantity. Only consider characters, not spaces or punctuation. Consider capital letters to be the same as lower case.

# Examples

- anagrams("rail safety", "fairy tales") --> true
- anagrmas("RAIL! SAFETY!", "fairy tales") --> true
- anagrams("Hi there!", "Bye there") --> false

# Solution: clean, sort and compare

## Pseudocode

- filter spaces and special characters from both strings
- lowercase both strings
- sort the letters in both strings
- compare both strings
- if equal
  return true
- else return false

```go
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
```
