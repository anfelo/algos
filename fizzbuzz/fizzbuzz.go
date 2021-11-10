package fizzbuzz

import "strconv"

func Fizzbuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		if i%3 == 0 && i%5 == 0 {
			// Is the number multiple of 3 and 5?
			s = "fizzbuzz"
		} else if i%3 == 0 {
			// Is the number multiple of 3?
			s = "fizz"
		} else if i%5 == 0 {
			// Is the number multiple of 5?
			s = "buzz"
		}
		result = append(result, s)
	}
	return result
}
