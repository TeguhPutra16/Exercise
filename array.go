package main

import "fmt"

func main() {
	a := []int{1, 2, 2, 3}

	x := removeDuplicates(a)
	fmt.Println(x)

}

func remove(arr []int) []int {

	mapArr := map[int]bool{}
	result := []int{}

	for _, val := range arr {
		if _, ok := mapArr[val]; !ok {
			mapArr[val] = true
			result = append(result, val)
		}

	}
	return result

}

func removeDuplicates(nums []int) []int {
	result := []int{} //1,2

	// Loop through the array and add numbers to the result slice only if they're not already in it
	for _, num := range nums {
		found := false
		for _, r := range result {
			if r == num {
				found = true
				break
			}
		}
		if !found {
			result = append(result, num)
		}
	}

	return result
}
