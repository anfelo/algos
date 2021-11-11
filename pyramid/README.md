# Directions

Write a function that accepts a positive number N. The function should console log a pyramid shape with N levels using the # character. Make sure the pyramid has spaces on both the left _and_ right hand sides

# Examples

- pyramid(1)

```
"#"
```

- pyramid(2)

```
" # "
"###"
```

- pyramid(3)

```
"  #  "
" ### "
"#####"
```

- pyramid(4)

```
"   #   "
"  ###  "
" ##### "
"#######"
```

# Solution 1: (double loop)

## Pseudocode

- create empty array
  - loop over the rows size n
    - create and append empty row
    - loop over the columns size 2\*n - 1
      - if column is between n - row and n + row append a "#" to the row
      - else append a " " to the row
- return array

```go
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
```

# Solution 2: (loop + string repeat)

## Pseudocode

- create empty array
- loop until n
  - create space string repeating " " n - i - 1 times
  - create right step string repeating "#" i + 1 times and adding right spaces
  - create left step string repeating "#" i times and adding left spaces
  - append to the array left step + right step
- return array

```go
func Pyramid(n int) []string {
	steps := []string{}
	for i := 0; i < n; i++ {
		spaces := strings.Repeat("-", n-i-1)
		rstep := strings.Repeat("#", i+1) + spaces
		lstep := spaces + strings.Repeat("#", i)
		steps = append(steps, lstep+rstep)
	}
	return steps
}
```

# Solution 3: (recursive)

## Pseudocode

- (base case) if the row is equal to n **return**
- if length of row string is equal to 2\*n-1
  - print row
  - call function adding 1 to the row
  - **return (important)**
- if length of row string is between [n - row - 1 and n + row) append a "#" to the row string
- else append a space to row string
- call function passing new row string

```go
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
```
