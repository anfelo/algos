# Directions

Given a string, return true if a string is a palindrome or false if it is not.
Palindromes are strings that form the same word when reversed. _Do_ include spaces
and punctuation in determining is the string is a palindrome.

## Examples

- palindrome("abba") == true
- palindrome("abcdefg") == false

# Solution 1: (string reversal appending)

## Pseudocode

- Reverse the string passed as an argument
- Compare the original string and the reversed

```go
func Palindrome(s string) bool {
	var r string
	for _, v := range s {
		r = string(v) + r
	}
	return r == s
}
```

# Solution 2: (checking char by oposing char)

## Pseudocode

- for every character in the string
  - Check if the current char is equal to the char in the oposite location
  - If false break and return false
- return true

_Note: A problem with this solution is that we end up with 1/2 redundancy check_

```go
func Palindrome(s string) bool {
	rns := []rune(s)
	isPalindrome := true
	for i, v := range s {
		if v != rns[len(rns)-1-i] {
			isPalindrome = false
			break
		}
	}
	return isPalindrome
}
```
