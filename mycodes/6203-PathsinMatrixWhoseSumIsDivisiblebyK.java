class Solution {
    int cnt = 0, m, n;
    int accuH[][], accuV[][];


    public int numberOfPaths(int[][] grid, int k) {
        m = grid.length, n = grid[0].length;
        accuH = new int[m][n];
        accuV = new int[m][n];

        //init dp
        for (int i = 0; i < m; i++) accuH[i][0] = grid[i][0];
        for (int i = 0; i < n; i++) accuV[0][i] = grid[0][i];
        for (int i = 0; i < m; i++) {
            for (int j = 1; j < n; j++) {
                accuH[i][j] = grid[i][j] + accuH[i][j - 1];
            }
        }
        for (int i = 0; i < n; i++) {
            for (int j = 1; j < m; j++) {
                accuV[j][i] = grid[j][i] + accuV[j - 1][i];
            }
        }

        
    }

    private void dfs(int i, int j, int sum, boolean dir, int k) {
        if (i == m - 1) {
            cnt += (sum + accuH[i][j] - accuH[i][j - 1]) % k == 0 ? 1 : 0;
            return
        }
        
    }
}