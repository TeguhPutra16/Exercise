package main

import "fmt"

func binarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)

		if arr[mid] < x {
			low = mid + 1
			fmt.Println("low", low)

		} else if arr[mid] > x {
			high = mid - 1
			fmt.Println("high", high)
		} else if arr[mid] == x {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	x := 4

	result := binarySearch(arr, x)

	if result == -1 {
		fmt.Printf("Element %d not found\n", x)
	} else {
		fmt.Printf("Element %d found at index %d\n", x, result)
	}
}
