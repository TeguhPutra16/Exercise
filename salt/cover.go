package main

import "fmt"

// import "fmt"

// func solution(n int) int {

// /* Cover The Floor
// Description
// You have a floor that is a grid of size 2 x N, consisting of 2 rows and N columns. Then, you want to cover
// the floor with ceramics. You have 2 types of ceramics: ceramics with size 1 x 1 and ceramics with size 1 x 2.

// Note that, you can rotate the type of ceramic that you have, so for ceramics of size 1 x 2 you can rotate them to ceramics of size 2 x 1. Count how many ways you can cover the floor with the types of available ceramics!

// Input Format
// You are given an integer N where N is a positive integer as previously described in the descriptions.

// Output Format
// Output 1 integer number that is the number of ways you can cover your floor with the available ceramic types. Since the answers can be very large, output the answer after modulo with 1,000,000,007.*/

// }
func solution(n int) int {
	const mod = 1000000007
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]*2) % mod
	}
	return dp[n]
}

// example :
// input n = 2
// outPut = 7
// there are 7 ways cover the floor

func main() {
	x := 2
	fmt.Println(solution(x))

}

/* Cover The Floor
Description
You have a floor that is a grid of size 2 x N, consisting of 2 rows and N columns. Then, you want to cover
the floor with ceramics. You have 2 types of ceramics: ceramics with size 1 x 1 and ceramics with size 1 x 2.

Note that, you can rotate the type of ceramic that you have, so for ceramics of size 1 x 2 you can rotate them to ceramics of size 2 x 1. Count how many ways you can cover the floor with the types of available ceramics!

Input Format
You are given an integer N where N is a positive integer as previously described in the descriptions.

Output Format
Output 1 integer number that is the number of ways you can cover your floor with the available ceramic types. Since the answers can be very large, output the answer after modulo with 1,000,000,007.*/
