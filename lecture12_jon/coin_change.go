package lecture12_jon

/*
Given an unlimitted suply of coins of given denominations, find the total number of ways to make a change of size n
e.g coins: 1, 3, 5, 10
1. Objective function f(i) = total number of ways to make a change of size i
2. Base case f(0) = 1 --> do nothing
   f(1) = 1 way (1 cent)
   f(2) = 1 way (1 cent and 1 cent)
   f(3) = f(2) + f(0) = 1 + 1 = 2 ways
   f(n) = f(n - 1) + f(n - 3) + f(n - 5) + f(n - 10)
*/
func change(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if i >= 1 {
			dp[i] += dp[i-1]
		}
		if i >= 3 {
			dp[i] += dp[i-3]
		}
		if i >= 5 {
			dp[i] += dp[i-5]
		}
		if i >= 10 {
			dp[i] += dp[i-10]
		}
	}
	return dp[n]
}
