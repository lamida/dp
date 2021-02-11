"""
Problem: 
Find the sum of the first N numbers.

Objective function:
f(i) is the sum of first i elements.

Recurrence relation:
f(n) = f(n-1) + n
"""
def nSum(n: int) -> int:
    dp = [0] * (n+1)
    dp[0] = 0
    for i in range(1, len(dp)):
        dp[i] = i + dp[i-1]
    return dp[n]

if __name__ == "__main__":
    for i in range(10):
        print(i, nSum(i))
