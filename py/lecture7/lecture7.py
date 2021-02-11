from typing import List
"""
Problem:
Paid Staircase

You are climbing a paid staircase. It takes n steps to reach to the top and you have to pay
p[i] to step on the i-th stair. Each time you can climb 1 or 2 steps.
What's the cheapest amount you  have to pay to get to the top of the staircase?
"""
def paidStair(n: int: p: List[int]):
    dp = [0] * (n+1)
    dp[0], dp[1] = 1, 1
    for i in range(2, len(dp)):
        dp[i] = min(dp[i-1], dp[i-2]) * p[i]
    return dp[n]