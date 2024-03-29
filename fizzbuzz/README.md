# Directions

Write a program that console logs the numbers from 1 to n. But for multiples of three print "fizz" instead of the number and for the multiples of five print "buzz". For numbers which are multiples of both three and five print "fizzbuzz"

# Example

- fizzbuzz(5)

```
// 1
// 2
// fizz
// 4
// buzz
```

# Solution

## Pseudocode

- loop from 0 to the number
  - if i is multiple of 3
    - print "fizz"
  - if i is multiple of 5
    - print "buzz"
  - if i is both multiple of 3 and 5
    - print "fizzbuzz"
  - else
    - print i

```go
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
```
