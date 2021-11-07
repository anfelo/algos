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
