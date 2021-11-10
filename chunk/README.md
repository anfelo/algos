# Directions

Given an array and a chunk size, divide the array into many subarrays where each subarray is of length size

# Examples

- chunk([1,2,3,4], 2) --> [[1,2], [3,4]]
- chunk([1,2,3,4,5], 2) --> [[1,2], [3,4], [5]]
- chunk([1,2,3,4,6,7,8], 3) --> [[1,2,3], [4,5,6], [7,8]]
- chunk([1,2,3,4,5], 4) --> [[1,2,3,4], [5]]
- chunk([1,2,3,4,5], 10) --> [[1,2,3,4,5]]

# Solution 1: (using slices syntax)

## Pseudocode

- create an empty chunks array
- loop over all items in the array in steps of size
  - append to chunks a slice from i to i + size
- return chunks

```go
func Chunk(arr []int, size int) [][]int {
	var chunks [][]int
	for i := 0; i < len(arr); i += size {
		max_size := math.Min(float64(len(arr)), float64(i+size))
		chunks = append(chunks, arr[i:int(max_size)])
	}
	return chunks
}
```

# Solution 2: (filling the chunks loop over the array)

## Pseudocode

- create an empty chunks array
- loop over all items in the array
  - if the last chunk size is equal to the specified length
    - create a new empty chunk and append it to the chunks array
  - else
    - append the current item to the last chunk
- return chunks

```go
func Chunk(arr []int, size int) [][]int {
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
```

# Solution 3: (filling the chunks loop over the chunks)

_Note: not the most clean solution_

```go
func Chunk(arr []int, size int) [][]int {
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
```
