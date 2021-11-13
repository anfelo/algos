# Directions

Print out the n-th entry in the fibonacci series. The fibonacci series is an ordering of numbers where each number is the sum of the preceeding two.

# Examples

* Series: [0, 1, 1, 2, 3, 5, 8, 13, 21, 34...]
* fib(4) === 3

# Solution 1: (iterative)

## Pseudocode

* create an array of int starting with 0 and 1 as first items
* loop from 2 to n (included)
  * append to the array the value of (i - 2) + (i - 1) of the array
* return value n of the array

```go
func Fib(n int) int {
	fs := []int{0, 1}
	for i := 2; i <= n; i++ {
		fs = append(fs, fs[i-2]+fs[i-1])
	}
	return fs[n]
}
```

# Solution 2: (recursion)

## Pseudocode

* if n is less than n
  * return n
* return the call to fib(n-1) + fib(n-2)

```go
func FibRec(n int) int {
	if n < 2 {
		return n
	}

	return FibRec(n-2) + FibRec(n-1)
}
```

# Solution 3: (recursion + memoization)

**Memoization:** Store arguments of each function call along with the result. If the function is called again with the same arguments, return the precomputed result, rather than running the function again

## Pseudocode

* create a function that accepts as an argument a function and return a new function
* create a cache object (map)
* create a new funtion
  * if arg is in the cache return value in cache
  * else call the initial function with the arg and store it in cache
* return the newly created function

```go
// Recursive solution
func FibRec(n int) int {
	if n < 2 {
		return n
	}

	return FibRec(n-2) + FibRec(n-1)
}

func Mem(f func(int) int) func(int) int {
	mem := make(map[int]int)
	inner := func(n int) int {
		if v, exists := mem[n]; exists {
			return v
		}
		result := f(n)
		mem[n] = result
		return result
	}
	return inner
}

func main() {
	fastFib := Mem(FibRec)
	fastFib(5) // 5
}
```