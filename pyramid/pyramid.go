package pyramid

import (
	"fmt"
	"strings"
)

func Pyramid(n int) []string {
	steps := []string{}
	for i := 0; i < n; i++ {
		steps = append(steps, "")
		for j := 0; j < 2*n-1; j++ {
			if j >= n-i-1 && j < n+i {
				steps[i] = steps[i] + "#"
			} else {
				steps[i] = steps[i] + "-"
			}
		}
	}
	return steps
}

func Pyramid1(n int) []string {
	steps := []string{}
	for i := 0; i < n; i++ {
		spaces := strings.Repeat("-", n-i-1)
		rstep := strings.Repeat("#", i+1) + spaces
		lstep := spaces + strings.Repeat("#", i)
		steps = append(steps, lstep+rstep)
	}
	return steps
}

func Pyramid2(n int, s string, r int) {
	// Exit
	if n == r {
		return
	}

	// New row
	if len(s) == 2*n-1 {
		fmt.Println(s)
		Pyramid2(n, "", r+1)
		return
	}

	if len(s) >= n-r-1 && len(s) < n+r {
		s += "#"
	} else {
		s += "-"
	}
	Pyramid2(n, s, r)
}
