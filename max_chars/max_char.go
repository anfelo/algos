package maxchar

func MaxChar(s string) string {
	ll := make(map[string]int)
	for _, v := range s {
		_, exists := ll[string(v)]
		if exists {
			ll[string(v)] += 1
		} else {
			ll[string(v)] = 1
		}
	}
	var max string
	for k, v := range ll {
		if max == "" {
			max = k
		}

		if v > ll[max] {
			max = k
		}
	}
	return string(max)
}
