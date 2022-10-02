class Solution {
    public int maxSum(int[][] grid) {
        int m = grid.length, n = grid[0].length, max = 0;
        for (int i = 0; i <= m - 3; i++)
            for (int j = 0; j <= n - 3; j++)
                max = Math.max(max, getSum(grid, i, j));
        return max;
    }

    private int getSum(int[][] grid, int i, int j) {
        int sum = 0;
        for (int k = 0; k < 3; k++) {
            sum += grid[i][j + k];
            sum += grid[i + 2][j + k];
        }
        sum += grid[i + 1][j + 1];
        return sum;
    }
}