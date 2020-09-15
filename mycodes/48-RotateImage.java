//顺时针链式替换
class Solution {
    public void rotate(int[][] matrix) {
        int N = matrix.length;
        for (int i = 0; i < N / 2; ++i) {
            for(int j = i; j < N - i - 1; ++j) {
                int next = matrix[j][N - 1 - i], cur = matrix[i][j];
                matrix[j][N - 1 - i] = cur;
                cur = next;
                next = matrix[N - 1 - i][N - 1 - j];
                matrix[N - 1 - i][N - 1 - j] = cur;
                cur = next;
                next = matrix[N - 1 - j][i];
                matrix[N - 1 - j][i] = cur;
                cur = next;
                next = matrix[i][j];
                matrix[i][j] = cur;
                cur = next;
            }
        }
    }
}

//逆时针链式替换
class Solution2 {
    public void rotate(int[][] matrix) {
        int N = matrix.length;
        for (int i = 0; i < N / 2; ++i) {
            for(int j = i; j < N - i - 1; ++j) {
                int pre = matrix[i][j];
                matrix[i][j] = matrix[N - 1 - j][i];
                matrix[N - 1 - j][i] = matrix[N - 1 - i][N - 1 - j];
                matrix[N - 1 - i][N - 1 - j] = matrix[j][N - 1 - i];
                matrix[j][N - 1 - i] = pre;
            }
        }
    }
}