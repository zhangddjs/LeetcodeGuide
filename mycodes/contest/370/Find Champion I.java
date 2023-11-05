class Solution {
    public int findChampion(int[][] grid) {
        for (int i = 0; i < grid.length; i++) {
            int cnt = 0;
            for (int j = 0; j < grid[i].length; j++) {
                cnt += grid[i][j];
            }
            if (cnt == grid.length - 1) {
                return i;
            }
        }
        return 0;
    }
}