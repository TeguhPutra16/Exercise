package main

import "fmt"

func maxProfit(prices []int, i int) int {
	// 8, 2, 5, 80, 12, 23
	if len(prices) < 2 || i < 1 {
		return 0
	}

	n := len(prices)
	profit := make([]int, n)

	for k := 1; k <= i; k++ {
		minPrice := prices[0]
		fmt.Println("k", k)
		for j := 1; j < n; j++ {
			profit[j] = max(profit[j-1], prices[j]-minPrice)
			minPrice = min(minPrice, prices[j]-profit[j-1])

		}
	}

	return profit[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	arr := []int{8, 2, 5, 80, 12, 23}
	res := maxProfit(arr, 3)
	fmt.Println(res)

}
