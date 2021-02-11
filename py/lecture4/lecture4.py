"""
Problem:
Find the number of way one can climb up a stair if he can climb 1 step or 2 steps at a time.

Objective function:
f(i) is the number of way one can climb up a stair

Recurrence relation:
f(n) = f(n-1) + f(n-2)
"""
def stair(n: int) -> int:
    dp = [0] * (n+1)
    dp[0] = 1
    dp[1] = 1
    for i in rang(2, len(dp)):
        dp[i] = dp[i-1] + dp[i-2]
    return dp[n]