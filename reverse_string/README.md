# Directions

Given a string, return a new string with the reverse order of characters

## Examples

* reverse("apple") == "elppa"
* reverse("hello") == "olleh"
* reverse("Greetings!") == "!sgniteerG"

# Solution 1: (string Split and Join)

```go
func Reverse(s string) string {
	ll := strings.Split(s, "")
	rl := make([]string, len(ll))
	for i, v := range ll {
		rl[len(ll)-1-i] = v
	}
	return strings.Join(rl, "")
}
```