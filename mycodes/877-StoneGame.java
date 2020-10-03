//Time 40%
//Space 51%

class Solution1 {
    public boolean stoneGame(int[] piles) {
        //每次取两端中的一端
        //如果不能赢，回溯
        //如果无论怎么取都不能赢，返回false
        //否则当某种情况能赢，则返回true
        //用动态规划或带备忘录的递归，将每种情况下Alex的优势进行缓存。
        //基本Case1: 如果只有1组石头，Lee永远赢，每种情况dp[i][i]对应Alex优势为负数，即-piles[i]
        //基本Case2: 如果只有2组石头，选择个数多的赢，Alex优势为Math.max(dp[0][0] + piles[1], dp[1][1] + piles[0])
        //通常Case: 假设当前情况为[i, j]，则下一情况为[i + 1, j]、[i, j - 1]，选择对各自优势更大的那个，并进行优势累加或累减(i 和 j 距离为偶数时，表示有奇数组石头，Lee取，取累减优势后优势更小的，否则表示Alex取，取累加优势后优势更大的)
        int N = piles.length;
        int [][] dp = new int[N][N];
        //init
        for (int i = 0; i < N; ++i) dp[i][i] = -piles[i];
        
        //dp
        for (int dist = 1; dist < N; ++dist) {
            for (int i = 0; i < N - dist; ++i) {
                if (dist % 2 != 0)
                    dp[i][i + dist] = Math.max(piles[i] + dp[i + 1][i + dist],
                                               piles[i + dist] + dp[i][i + dist - 1]);
                else
                    dp[i][i + dist] = Math.min(dp[i + 1][i + dist] - piles[i],
                                               dp[i][i + dist - 1] - piles[i + dist]);
            }
        }
        
        return dp[0][piles.length - 1] > 0;
    }
}

//Time 100%
//Space 88%

class Solution2 {
    public boolean stoneGame(int[] piles) {
        return true;
    }
}
