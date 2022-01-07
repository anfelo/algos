package maxchar

func MaxChar(s string) string {
	currentMaxChar := ""
	ll := map[string]int{currentMaxChar: 0}
	for _, v := range s {
		ll[string(v)] += 1

		if ll[string(v)] > ll[currentMaxChar] {
			currentMaxChar = string(v)
		}
	}

	return currentMaxChar
}
