package main

import "fmt"

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			fmt.Println("a", arr[j])
			fmt.Println("b", arr[j+1])
			fmt.Println("c", key)
			arr[j+1] = arr[j]

			j--
		}
		arr[j+1] = key
		fmt.Println("d", arr[j+1])

	}
	return arr
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	sortedArr := insertionSort(arr)
	fmt.Println(sortedArr)
}
