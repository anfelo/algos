# SelectionSort

**(Prove me wrong algo)** Assume current index is the lowest in the array, compare it with all. If there is a lower value, move it to the beginning

## Pseudocode

* From i = 0 < array length
	* assume the element at i is the least in the array, assign i to indexOfMin
	* loop from i + 1 to end of array
		* see if there is an element with lower value
			* if there is, record its index
	* if the index of the current element and the index of the 'lowest' element is not the same, swap them
* return array