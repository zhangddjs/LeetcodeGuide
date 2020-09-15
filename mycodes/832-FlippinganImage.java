class Solution {
    public int[][] flipAndInvertImage(int[][] A) {
        for (int i = 0; i < A.length; ++i) {
            for (int j = 0; j < A[i].length / 2; ++j){
                A[i][j] = A[i][j] + A[i][A[i].length - j - 1];
                A[i][A[i].length - j - 1] = A[i][j] - A[i][A[i].length - j - 1];
                A[i][j] -= A[i][A[i].length - j - 1];
                A[i][j] = 1 - A[i][j];
                A[i][A[i].length - j - 1] = 1 - A[i][A[i].length - j - 1];
            }
            if (A[i].length % 2 == 1) A[i][A[i].length / 2] = 1 - A[i][A[i].length / 2];
        }
        return A;
    }
}