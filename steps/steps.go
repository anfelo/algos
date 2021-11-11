package steps

import (
	"fmt"
	"strings"
)

func Steps(n int) []string {
	steps := []string{}
	for i := 0; i < n; i++ {
		steps = append(steps, "")
		for j := 0; j < n; j++ {
			if j > i {
				steps[i] = steps[i] + " "
			} else {
				steps[i] = steps[i] + "#"
			}
		}
	}
	return steps
}

func Steps1(n int) []string {
	steps := []string{}
	for i := 1; i <= n; i++ {
		step := strings.Repeat("#", i)
		spaces := strings.Repeat(" ", n-i)
		steps = append(steps, step+spaces)
	}
	return steps
}

func Steps2(n int, row int) {
	// Base case
	if n == row {
		return
	}

	fmt.Println(strings.Repeat("#", row+1) + strings.Repeat("*", (n-row-1)))
	Steps2(n, row+1)
}
