class Solution {
    public int numMagicSquaresInside(int[][] grid) {
        if(grid.length < 3 || grid[0].length < 3) return 0;
        int res = 0;
        for (int i = 0; i < grid.length - 2; i++)
            for (int j = 0; j < grid[i].length - 2; j++)
                res += isMagic(grid, i, j);
        return res;
    }
    public int isMagic(int[][] grid, int i, int j) {
        int sum = grid[i][j] + grid[i][j + 1] + grid[i][j + 2];
        Set<Integer> set = new HashSet<>();
        int [] diagtemps = new int[2];
        for (int k = 0; k < 3; k++) {
            int rowtemp = 0, coltemp = 0;
            for (int h = 0; h < 3; h++) {
                if (set.remove(grid[i + k][j + h]) ||
                    grid[i + k][j + h] > 9 ||
                    grid[i + k][j + h] < 1) return 0;
                set.add(grid[i + k][j + h]);
                rowtemp += grid[i + k][j + h];
                coltemp += grid[i + h][j + k];
                if(h == k) diagtemps[0] += grid[i + k][j + h];
                if(h + k == 2) diagtemps[1] += grid[i + k][j + h];
            }
            if (rowtemp != sum || coltemp != sum) return 0;
        }
        if(diagtemps[0] != sum || diagtemps[1] != sum) return 0;
        return 1;
    }
}