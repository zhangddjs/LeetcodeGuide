class Solution {
    public int largestMagicSquare(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int max = 1;
        int[][] prefixSumRow = new int[m][n];
        int[][] prefixSumCol = new int[n][m];
        for (int i = 0; i < m; i++) {
            prefixSumRow[i][0] = grid[i][0];
            for (int j = 1; j < n; j++) prefixSumRow[i][j] = prefixSumRow[i][j - 1] + grid[i][j];
        }
        for (int i = 0; i < n; i++) {
            prefixSumCol[i][0] = grid[0][i];
            for (int j = 1; j < m; j++) prefixSumCol[i][j] = prefixSumCol[i][j - 1] + grid[j][i];
        }
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                for (int size = Math.min(m - i, n - j); size > 1; size--) {
                    if (isMagicSquare(grid, i, j, size, prefixSumRow, prefixSumCol)) {
                        max = Math.max(max, size);
                        break;
                    }
                }
            }
        }
        return max;
    }

    private boolean isMagicSquare(int[][] grid, int row, int col, int size, int[][] prefixSumRow, int[][] prefixSumCol) {
        //check row and col
        int sumRow = col == 0 ? prefixSumRow[row][size - 1] : prefixSumRow[row][col + size - 1] - prefixSumRow[row][col - 1];
        int sumCol = row == 0 ? prefixSumCol[col][size - 1] : prefixSumCol[col][row + size - 1] - prefixSumCol[col][row - 1];
        if (sumRow != sumCol) return false;
        for (int i = row + 1; i < row + size; i++) {
            int minus = col == 0 ? 0 : prefixSumRow[i][col - 1];
            if (prefixSumRow[i][col + size - 1] - minus != sumRow) return false;
        }
        for (int i = col + 1; i < col + size; i++) {
            int minus = row == 0 ? 0 : prefixSumCol[i][row - 1];
            if (prefixSumCol[i][row + size - 1] - minus != sumCol) return false;
        }
        
        //check diag
        int sumMain = 0, sumSub = 0;
        for (int i = 0; i < size; i++) {
            sumMain += grid[row + i][col + i];
            sumSub += grid[row + i][col + size - 1 - i];
        }

        return sumMain == sumSub && sumMain == sumRow;
    }
}