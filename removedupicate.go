package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(remove(nums))

}

func remove(nums []int) int {

	arrMap := make(map[int]int)
	res := []int{}

	for _, v := range nums {

		arrMap[v] = +1

	}
	for k, _ := range arrMap {
		res = append(res, k)
	}
	sort.SliceStable(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	fmt.Println(res)
	return len(res)

}
