class Solution {
    public int equalPairs(int[][] grid) {
        int cnt = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][0] != grid[0][j]) continue;
                int k = 0;
                for (k = 0; k < grid.length; k++) if (grid[i][k] != grid[k][j]) break;
                cnt += k == grid.length ? 1 : 0;
            }
        }
        return cnt;
    }
}