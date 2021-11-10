# Directions

Given a string, return the character that is most commonly used in the string.

## Examples

- maxchar("abbaa") == "a"
- maxchar("abcdefg 12311") == "1"

## Common String Questions

- What is the most common character in the string?
- Does string A have the same characters as string B (is it an anagram)?
- Does the given string have any repeated characters in it?

# Pseudocode

- create an empty map (object with unique keys)
- loop over all the chars in string
  - if the key is not yet in the map
    - add the key to the map with value 1
  - else sum the value by 1
- create a reference to the first key
- loop over the map keys
  - if the value is greater than the prev stored
    store the new key
- return the key

# Solution 1: char map

```go
func MaxChar(s string) string {
	ll := make(map[string]int)
	for _, v := range s {
		_, exists := ll[string(v)]
		if exists {
			ll[string(v)] += 1
		} else {
			ll[string(v)] = 1
		}
	}
	var max string
	for k, v := range ll {
		if max == "" {
			max = k
		}

		if v > ll[max] {
			max = k
		}
	}
	return string(max)
}
```
