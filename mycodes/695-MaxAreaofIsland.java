class Solution {
    public int maxAreaOfIsland(int[][] grid) {
        if (grid.length == 0 || grid[0].length == 0) return 0;
        int res = 0;
        for (int i = 0; i < grid.length; ++i)
            for (int j = 0; j < grid[0].length; ++j)
                res = Math.max(dfs(grid, i, j), res);
        return res;
    }

    public int dfs(int [][] grid, int i, int j) {
        if (i < 0 || j < 0 || i >= grid.length || j >= grid[0].length || grid[i][j] != 1) return 0;
        int tmp = 1;
        grid[i][j] = 2;
        tmp += dfs(grid, i - 1, j);
        tmp += dfs(grid, i + 1, j);
        tmp += dfs(grid, i, j - 1);
        tmp += dfs(grid, i, j + 1);
        return tmp;
    }
}