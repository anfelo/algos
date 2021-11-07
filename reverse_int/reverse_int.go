package reverseint

import "strconv"

func Reverse(num int) int {
	sign := ""
	// Convert to string
	strNum := strconv.Itoa(num)
	// Store sign
	if num < 0 {
		sign = "-"
		strNum = strNum[1:]
	}
	// reverse string
	var r string
	for _, v := range strNum {
		r = string(v) + r
	}
	// Add back sign
	r = sign + r
	// Convert to int
	i, _ := strconv.Atoi(r)
	return i
}
