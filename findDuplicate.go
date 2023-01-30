package main

import (
	"fmt"
	"sort"
)

func findDuplicationNumber(numbers []int) {
	MapNumbers := map[int]int{}
	DuplicateNum := []int{}

	for _, v := range numbers {
		MapNumbers[v]++
	}

	for k, v := range MapNumbers {
		if v > 1 {
			DuplicateNum = append(DuplicateNum, k)
		}

	}

	sort.SliceStable(DuplicateNum, func(i, j int) bool {
		return DuplicateNum[i] < DuplicateNum[j]
	})
	fmt.Println(DuplicateNum)

}

func main() {
	findDuplicationNumber([]int{1, 1})                            //[1]
	findDuplicationNumber([]int{1, 2, 3, 4})                      //[]
	findDuplicationNumber([]int{1, 2, 1, 2, 2, 3, 4, 5, 5, 5, 5}) //[1,2,5]
}
