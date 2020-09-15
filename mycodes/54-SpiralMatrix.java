class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> res = new ArrayList<>();
        if (matrix.length == 0 || matrix[0].length == 0) return res;
        int M = matrix.length, N = matrix[0].length;
        for (int i = 0; i <= (Math.min(M, N) + 1) / 2 - 1; ++i) {
            for (int j = i; j <= N - 1 - i; ++j) res.add(matrix[i][j]);
            for (int j = i + 1; j <= M - 1 - i; ++j) res.add(matrix[j][N - 1 - i]);
            for (int j = N - 2 - i; i != M - 1 - i && j > i; --j) res.add(matrix[M - 1 - i][j]);
            for (int j = M - 1 - i; i != N - 1 - i && j > i; --j) res.add(matrix[j][i]);
        }
        return res;
    }
}