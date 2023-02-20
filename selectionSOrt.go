package main

import "fmt"

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		fmt.Println("i", i)
		for j := i + 1; j < len(arr); j++ {
			fmt.Println("minidx", arr[minIdx])
			fmt.Println("j", arr[j])
			if arr[j] < arr[minIdx] {
				fmt.Println("masuk")
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
		fmt.Println(arr)
	}
	return arr
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	sortedArr := selectionSort(arr)
	fmt.Println(sortedArr)
}
