# Directions

Given a string, return true if a string is a palindrome or false if it is not.
Palindromes are strings that form the same word when reversed. *Do* include spaces
and punctuation in determining is the string is a palindrome.

## Examples

* palindrome("abba") == true
* palindrome("abcdefg") == false

# Pseudocode

## Option 1:

* Reverse the string passed as an argument
* Compare the original string and the reversed

## Option 2:

* for every character in the string
	* Check if the current char is equal to the char in the oposite location
	* If false break and return false
* return true



# Solution 1: (string reversal appending)

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

*Note: A problem with this solution is that we end up with 1/2 redundancy check*

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