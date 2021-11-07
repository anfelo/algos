# Directions

Given an integer, return an integer that is the reverse ordering of numbers

## Examples

* reverse(15) == 51
* reverse(981) == 189
* reverse(500) == 5
* reverse(-15) == -51
* reverse(-90) == -90

# Pseudocode

* convert int to string
* if num is negative
	* store the sign if negative number
	* remove minus sign from the string
* reverse the string
* add back the sign
* convert to int
* return int

# Solution

```go
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
```