# Directions

Write a function that accepts a string. The function should capitalize the first letter of each word in the string then return the capitalized string.

# Examples

- capitalize("hi there, how is it going?") --> "Hi There, How Is It Going?"
- capitalize("i love breakfast at bill miller bbq") --> "I Love Breakfast At Bill Miller Bbq"

# Solution: (split words capitalize first char)

```go
func Capitalize(s string) string {
	w := strings.Split(s, " ")
	for i, v := range w {
		cw := strings.ToUpper(string(v[0])) + v[1:]
		w[i] = cw
	}
	return strings.Join(w, " ")
}
```
