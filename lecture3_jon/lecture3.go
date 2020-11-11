package lecture3_jon

/*
In this function we are using recursion in getting the sum
*/
func nSum(n int) int {
	if n <= 1 {
		return n
	} else {
		return nSum(n-1) + n
	}
}

/*
Use dynamic programming to sum the n
*/
func nSumDp(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i
	}
	return dp[n]
}
