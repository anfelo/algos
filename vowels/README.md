# Directions

Write a function that returns the number of vowels used in a string. Vowels are the characters 'a', 'e', 'i', 'o', and 'u'.

# Examples

- vowels("Hi There!") --> 3
- vowels("Why do you asdk?") --> 4
- vowels("Why?") --> 0

# Solution 1: (loop + contains)

## Pseudocode

- create a counter starting in 0
- create an array of the vowels ["a", "e", "i", "o", "u"]
- loop over the string chars
  - if vowels array contains the current char(lowercased)
    - add 1 to counter
- return counter

```go
func Vowels(s string) int {
	vv := []string{"a", "e", "i", "o", "u"}
	counter := 0
	for _, v := range strings.Split(s, "") {
		if contains(vv, v) {
			counter++
		}
	}
	return counter
}

// Helper contains method (other languages have built-in)
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == strings.ToLower(e) {
			return true
		}
	}
	return false
}
```

# Solution 2: (regexp)

## Pseudocode

- create a regexp with the vowels
- match only the vowels in string
- return the length of the match

```go
func Vowels(s string) int {
	re := regexp.MustCompile(`[^aeiouAEIOU]`)
	v := re.ReplaceAllString(s, "")

	return len(v)
}
```
