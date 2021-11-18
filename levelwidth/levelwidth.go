package levelwidth

func LevelWidth(r *Node) []int {
	c := make([]int, 1)
	arr := []*Node{r, NewNode(-999)} // Stop symbol is -999
	for len(arr) > 1 {
		var curr *Node
		curr, arr = arr[0], arr[1:]
		if curr.Data == -999 {
			arr = append(arr, curr)
			c = append(c, 0)
		} else {
			c[len(c)-1]++
			arr = append(arr, curr.Children...)
		}
	}
	return c
}
