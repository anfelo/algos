# InsertionSort

With insertion sort, you treat the first part of your list as sorted and the second part of your list as unsorted.

Best case complexity: O(n)
Average/Worst case complexity: O(n^2)

# Pseudocode

* item in index 0 is sorted
* loop from index 1 to length of array
	* loop in reverse from i to 0
		* if arr[j+1] < arr[j]
			* move arr[j+1] to left of arr[j] (swap em)
* return arr