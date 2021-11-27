# QuickSort

It's another divide-and-conquer, recursive algorithm but it takes a slightly different approach. The basic gist is that you take the last element in the list and call that the pivot. Everything that's smaller than the pivot gets put into a "left" list and everything that's greater get's put in a "right" list. You then call quick sort on the left and right lists independently (hence the recursion.) After those two sorts come back, you concatenate the sorted left list, the pivot, and then the right list (in that order.) The base case is when you have a list of length 1 or 0, where you just return the list given to you.

Best case runtime: O(n^2) --> when items are already sorted
Average case runtime: O(n*log(n))
Worst case runtime: O(n*log(n)) --> when randomly sorted list

## Pseudocode

* if array length is 0 or 1 return array (base case)
* select the pivot (last element in the array)
* create two lists, left and right
* loop over the elements in the array minus the pivot
	* if current item is less than the pivot
		* add item to the left list
	* else add item to the right
* recurse with left, recurse with right
* concat left, pivot, right

```go
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	p := arr[len(arr)-1]
	var l []int
	var r []int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < p {
			l = append(l, arr[i])
		} else {
			r = append(r, arr[i])
		}
	}

	result := QuickSort(l)
	result = append(result, p)
	result = append(result, QuickSort(r)...)
	return result
}
```