# https://www.youtube.com/watch?v=3hHmUszRXjw&list=PLVrpF4r7WIhTT1hJqZmjP10nxsmrbRvlf&index=8&ab_channel=AndreyGrehov
# This is still wrong
from typing import List
def paidStair(n: int, p: List[int]) -> int:
    dp = [0] * (n+1)
    fr = [0] * (n+1)
    dp[0] = 1
    dp[1] = 1
    for i in range(2, len(dp)):
        dp[i] = min(dp[i-1], dp[i-2]) + p[i]
        if dp[i-1] < dp[i-2]:
            fr[i] = i-1
        else:
            fr[i] = i-2
    path = []
    current = n
    while current > 0:
        path.append(current)
        current = fr[current]
    path.append(0)
    return path[::-1]

if __name__ == "__main__":
    n = 8
    p = [0, 3, 2, 4, 6, 1, 1, 5, 3]
    want = [0, 2, 3, 5, 6, 8]
    res = paidStair(n, p)
    print(res, want)
