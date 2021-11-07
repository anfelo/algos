package reversestring

import (
	"strings"
)

func Reverse(s string) string {
	ll := strings.Split(s, "")
	rl := make([]string, len(ll))
	for i, v := range ll {
		rl[len(ll)-1-i] = v
	}
	return strings.Join(rl, "")
}

func Reverse1(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

// Preferred
func Reverse2(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
