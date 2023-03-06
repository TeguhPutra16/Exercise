package main

import (
	"fmt"
)

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	MapRanked, ArrRank := Ranked(ranked)
	MyLeaderBoard := []int32{}

	for i := 0; i < len(player); i++ {
		PlayerValue := player[i]

		if val, ok := MapRanked[PlayerValue]; ok {
			MyLeaderBoard = append(MyLeaderBoard, val)

		} else {
			if PlayerValue < ArrRank[len(ArrRank)-1] {

				MyLeaderBoard = append(MyLeaderBoard, MapRanked[ArrRank[len(ArrRank)-1]]+1)
			}
			if PlayerValue > ArrRank[0] {

				MyLeaderBoard = append(MyLeaderBoard, MapRanked[ArrRank[0]])
			}

			for j := 0; j < len(ArrRank)-1; j++ {

				if PlayerValue < ArrRank[j] && PlayerValue > ArrRank[j+1] {

					MyLeaderBoard = append(MyLeaderBoard, MapRanked[ArrRank[j+1]])
				}

			}
		}
	}
	return MyLeaderBoard

}

func Ranked(r []int32) (map[int32]int32, []int32) {
	arrmap := map[int32]int32{}
	arr := []int32{}
	var rank int32

	for _, val := range r {

		if _, ok := arrmap[val]; ok {
			continue
		}
		rank = rank + 1
		arr = append(arr, val)
		arrmap[val] = rank

	}
	return arrmap, arr

}

func main() {

	ranked := []int32{100, 90, 90, 80, 75, 60}
	player := []int32{50, 65, 77, 90, 102}
	a := climbingLeaderboard(ranked, player)
	fmt.Println(a)

}
