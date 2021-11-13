package fib

// Iterative solutiion
func Fib(n int) int {
	fs := []int{0, 1}
	for i := 2; i <= n; i++ {
		fs = append(fs, fs[i-2]+fs[i-1])
	}
	return fs[n]
}

// Recursive solution
func FibRec(n int) int {
	if n < 2 {
		return n
	}

	return FibRec(n-2) + FibRec(n-1)
}

func Mem(f func(int) int) func(int) int {
	mem := make(map[int]int)
	inner := func(n int) int {
		if v, exists := mem[n]; exists {
			return v
		}
		result := f(n)
		mem[n] = result
		return result
	}
	return inner
}
