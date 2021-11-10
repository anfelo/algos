package chunk

import "math"

func Chunk(arr []int, size int) [][]int {
	var chunks [][]int
	for i := 0; i < len(arr); i += size {
		max_size := math.Min(float64(len(arr)), float64(i+size))
		chunks = append(chunks, arr[i:int(max_size)])
	}
	return chunks
}

func Chunk1(arr []int, size int) [][]int {
	chunks := [][]int{{}}
	for i := range arr {
		last_chunk := chunks[len(chunks)-1]
		if len(last_chunk) == size {
			chunks = append(chunks, []int{i})
		} else {
			chunks[len(chunks)-1] = append(last_chunk, i)
		}
	}
	return chunks
}

func Chunk2(arr []int, size int) [][]int {
	num_chunks := math.Ceil(float64(len(arr)) / float64(size))
	chunks := make([][]int, int(num_chunks))
	for i := range chunks {
		max_size := math.Min(float64(len(arr)), float64(size*(i+1)))
		for j := i * size; j < int(max_size); j++ {
			chunks[i] = append(chunks[i], arr[j])
		}
	}
	return chunks
}
