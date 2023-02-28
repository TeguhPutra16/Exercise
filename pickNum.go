package main

import "fmt"

func pickingNumbers(a []int32) int32 {
	// Write your code here

	//identifikasi jumlah sub array
	//cuma bisa berbeda 1

	arrMap := map[int32]int32{}

	for _, value := range a {
		arrMap[value]++

	}
	max := max(arrMap)
	fmt.Println(arrMap)

	for key, _ := range arrMap {
		if arrMap[key] != 0 && arrMap[key+1] != 0 {
			current := arrMap[key] + arrMap[key+1]
			if current > max {
				max = current
			}
		}
	}

	return max

}

func max(a map[int32]int32) int32 {
	var max int32
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max

}

func main() {
	a := []int32{1, 1, 1, 1, 1}
	res := pickingNumbers(a)
	fmt.Println(res)

}
