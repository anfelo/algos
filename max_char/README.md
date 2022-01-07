# Directions

Given a string, return the character that is most commonly used in the string.

## Examples

- maxchar("abbaa") == "a"
- maxchar("abcdefg 12311") == "1"

## Common String Questions

- What is the most common character in the string?
- Does string A have the same characters as string B (is it an anagram)?
- Does the given string have any repeated characters in it?

# Solution 1: char map

## Pseudocode

- create an empty map (object with unique keys)
- initialize map with a maxChar key equals to 0
- keep a reference to the maxChar string
- loop over all the chars in string
  - if the key is not yet in the map
    - add the key to the map with value 1
  - else sum the value by 1
  - if the current char count is greater than the maxChar
    - set the maxChar to the current char

```go
func MaxChar(s string) string {
	currentMaxChar := ""
	ll := map[string]int{currentMaxChar: 0}
	for _, v := range s {
		ll[string(v)] += 1

		if ll[string(v)] > ll[currentMaxChar] {
			currentMaxChar = string(v)
		}
	}

	return currentMaxChar
}
```
