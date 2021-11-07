# Directions

Given a string, return a new string with the reverse order of characters

## Examples

* reverse("apple") == "elppa"
* reverse("hello") == "olleh"
* reverse("Greetings!") == "!sgniteerG"

# Pseudocode

* Create an empty string called 'reversed'
* for each character in provided string
    * Take the character and add it to the start of 'reversed'
* Return 'reversed'

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

# Solution 2: (using runes swap - only golang)

```go
func Reverse(s string) string {
	rns := []rune(s)
	// This type of for loops are prone to error. Many places to make typos
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
```

# Solution 3: (appending to the begining)

```go
func Reverse(s string) string {
    var result string
    for _, v := range s {
        result = string(v) + result
    }
    return result
}
```