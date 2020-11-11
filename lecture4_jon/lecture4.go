package lecture4_jon

/*
* f(i) Objective function: distinct number way to reach the i-th of stair
* base cases:
  f(2) = 2
  f(1) = 1
  f(0) = 1 --> do nothing

  How to come to the top of the stair n, if we can just 1 or 2 steps?
  * it will be the sum of f(n - 1) + f(n - 2)
*/
func climb(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
