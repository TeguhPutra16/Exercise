package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	arr := []int{3, 2, 1, 3}
	res := Birthdaycakecandlesn1(arr)
	fmt.Println(res)

}
func Birthdaycakecandlesn1(candles []int) int {

	MapCandles := map[int]int{}
	Max := 0
	go func() {
		for _, v := range candles {
			MapCandles[v]++
			if v > Max {
				Max = v
			}
		}

	}()
	time.Sleep(1 * time.Nanosecond)

	return MapCandles[Max]
}

func Birthdaycakecandlesn(candles []int) int {
	// Your code starts here.
	// Max := max(candles)
	sort.SliceStable(candles, func(i, j int) bool {
		return candles[i] > candles[j]
	})
	var result int
	// if len(candles)==0{
	//   return
	// }

	for _, v := range candles {
		if candles[0] == v {
			result += 1
		}
	}
	return result
}

func max(arr []int) int {

	max := 0

	for _, v := range arr {
		if v > max {
			max = v
		}

	}

	return max

}
