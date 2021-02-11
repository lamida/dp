def stair(n: int, k: int) -> int:
    dp = [0] * (n+1)
    dp[0] = 1
    for i in range(1, len(dp)):
        for j in range(1, k):
            if i - j < 0:
                continue
            dp[i % k] += dp[(i-j) % k]
    return dp[n % k]

if __name__ == "__main__":
    pass
