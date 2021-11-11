package matrix

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
