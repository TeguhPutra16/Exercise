package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var x []int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			} else if nums[i]+nums[j] == target {
				x = append(x, i)
				x = append(x, j)
				return x
			}

		}
	}
	return nil

}

func twoSum2(nums []int, target int) []int {

	var record = make(map[int]int)
	//result := make([]int,0,len(nums))
	var result []int

	for key, value := range nums {
		if val, ok := record[target-value]; ok {
			result = append(result, key)
			result = append(result, val)
			break
		}

		record[value] = key
	}

	return result
}

func twoSum3(nums []int, target int) []int {

	record := make(map[int]int)
	result := []int{}

	for k, v := range nums {
		fmt.Println("k,v", k, v)
		fmt.Println("map", record)
		if val, ok := record[v]; ok {
			//map map[2:2 3:1 5:0]
			//record[2]
			result = append(result, k)
			result = append(result, val)
			fmt.Println("k,val", k, val, ok)

			fmt.Println(record)

			return result
		}

		record[target-v] = k
	}

	return result
}

func main() {

	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum2(nums, target))

}
