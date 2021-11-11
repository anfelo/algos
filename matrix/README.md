# Directions

Write a function that accepts an integer N and returns a NxN spiral matrix.

# Examples

- matrix(2)

```
[[1, 2],
[4, 3]]
```

- matrix(3)

```
[[1, 2, 3],
[8, 9, 4],
[7, 6, 5]]
```

- matrix(4)

```
[[1, 2, 3, 4],
[12, 13, 14, 5],
[11, 16, 15, 6],
[10, 9, 8, 7]]
```

# Solution 1: (double loop)

## Pseudocode

- create an array of size n
- append n arrays of size n
- init rowstart = 0, rowend = n, colstart = 0, colend = n
- init counter = 1
- loop while rowstart < rowend and colstart < colend
  - loop from colstart to colend, +step
    - assing to matrix in rowstart, i the counter value
    - add 1 to counter
  - add 1 to rowstart
  - loop from rowstart to rowend, +step
    - assing to matrix in i, colend the counter value
    - add 1 to counter
  - subtract 1 to colend
  - loop from colend to colstart, -step
    - assing to matrix in rowend, i the counter value
    - add 1 to counter
  - subtract 1 to rowend
  - loop from rowend to rowstart, -step
    - assing to matrix in i, colstart the counter value
    - add 1 to counter
  - add 1 to colstart
- return matrix

```go
func Matrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	rs := 0
	re := n
	cs := 0
	ce := n
	c := 1

	for rs < re && cs < ce {
		for i := cs; i < ce; i++ {
			matrix[rs][i] = c
			c++
		}
		rs++
		for i := rs; i < re; i++ {
			matrix[i][ce-1] = c
			c++
		}
		ce--
		for i := ce - 1; i >= cs; i-- {
			matrix[re-1][i] = c
			c++
		}
		re--
		for i := re - 1; i >= rs; i-- {
			matrix[i][cs] = c
			c++
		}
		cs++
	}
	return matrix
}
```
