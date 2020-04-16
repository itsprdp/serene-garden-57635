package main

import (
	"fmt"
)

// i. Write a function that solves this problem by performing a linear scan.
func linearScan(array []int32, value int32) int {
	count := 0

	for _, num := range array {
		if num == value {
			count += 1
		}
	}

	return count
}

// modified binary search to return first and last occurance of a value
// for a given sorted array.
func binarySearch(array []int32, value int32, searchFirstOccurance bool) int {
	eleIdx, low, high := -1, 0, len(array)-1

	for low <= high {
		mid := (low + high) / 2
		switch {
		case array[mid] == value:
			eleIdx = mid

			// instead of returning the index of the occurance
			// find the lowest or highest occurance
			if searchFirstOccurance {
				high = mid - 1
			} else {
				low = mid + 1
			}
		case array[mid] < value:
			low = mid + 1
		case array[mid] > value:
			high = mid - 1
		}
	}

	return eleIdx
}

// ii. Next, write a function that solves this problem by performing two binary searches.
func twoBinarySearchScan(array []int32, value int32) int {
	firstOccuranceIdx := binarySearch(array, value, true)
	lastOccuranceIdx := binarySearch(array, value, false)

	// since the elements are sorted in the given array
	// the difference of the last and first occurance of value
	// would give away the total occurances of the value
	return ((lastOccuranceIdx - firstOccuranceIdx) + 1)
}

func main() {
	// Example:
	// if x=5 and a=[1,1,2,4,5,5,7,9], then the count is 2.
	a := []int32{1, 1, 2, 4, 5, 5, 7, 9}
	x := int32(5)
	fmt.Println("LinearScan: 5 appears", linearScan(a, x), "times")

	fmt.Println("TwoBinarySearchScan: 5 appears", twoBinarySearchScan(a, x), "times")
}
