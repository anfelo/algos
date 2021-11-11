# Directions

Write a function that accepts a positive number N. The function should console log a step shape with N levels using the # character. Make sure the step has spaces on the right hand side! (str repeat).

# Examples

- steps(2)
  "#-"
  "##"
- steps(3)
  "#--"
  "##-"
  "###"
- steps(4)
  "#---"
  "##--"
  "###-"
  "####"

# Solution 1: (double loop)

## Pseudocode

- create empty array
  - loop over the rows size n
    - create and append empty row
    - loop over the columns size n
      - if column > row append a space to the row
      - else append a "#" to the row
- return array

```go
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
```

# Solution 2: (loop string repeat)

## Pseudocode

- create empty array
- loop until n
  - create step string repeating "#" i times
  - create space string repeating " " n - i times
  - append to the array step + space
- return array

```go
func Steps(n int) []string {
	steps := []string{}
	for i := 1; i <= n; i++ {
		step := strings.Repeat("#", i)
		spaces := strings.Repeat(" ", n-i)
		steps = append(steps, step+spaces)
	}
	return steps
}
```

# Solution 3: (recursive + string repeat)

## Pseudocode

- if the iteration number is equal to n return
- print the step usign string repeat
- call the function incrementing the interation number

```go
func Steps(n int, row int) {
	// Base case
	if n == row {
		return
	}
	// Do something
	fmt.Println(strings.Repeat("#", row+1) + strings.Repeat("*", (n-row-1)))
	// Call next iteration
	Steps2(n, row+1)
}
```
