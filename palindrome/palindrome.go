package palindrome

func Palindrome(s string) bool {
	var r string
	for _, v := range s {
		r = string(v) + r
	}
	return r == s
}

func Palindrome1(s string) bool {
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
