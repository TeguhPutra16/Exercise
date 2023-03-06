package main

import "fmt"

func Shipping(N int, h []int) []int {

	x := []int{}

	for i := 0; i < N; i++ {
		min := Max(h) + N
		for j := 0; j < N; j++ {
			val := h[j] + abs(j-i)
			// arrMin = append(arrMin, val)
			if val < min {
				min = val
			}
		}
		// min := Min(arrMin)
		x = append(x, min)
	}
	return x
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}

	return a

}

func Max(arr []int) int {
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v

		}
	}
	return max

}

func main() {

	arr := []int{1, 4, 2, 4}
	fmt.Println(Shipping(4, arr))

}
