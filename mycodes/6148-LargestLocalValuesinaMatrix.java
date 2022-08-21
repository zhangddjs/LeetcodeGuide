class Solution {
    public int[][] largestLocal(int[][] grid) {
        int[][] res = new int[grid.length - 2][grid.length - 2];
        for (int i = 0; i < grid.length - 2; i++)
            for (int j = 0; j < grid.length - 2; j++)
                res[i][j] = getMax(grid, i, j);
        return res;
    }

    private int getMax(int[][] grid, int row, int col) {
        int max = grid[row][col];
        for (int i = row; i < row + 3; i++)
            for (int j = col; j < col + 3; j++)
                max = Math.max(grid[i][j], max);
        return max;
    }
}