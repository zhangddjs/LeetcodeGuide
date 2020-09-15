class Solution {
    public void setZeroes(int[][] matrix) {
        Set<Integer> set1 = new HashSet<>();
        Set<Integer> set2 = new HashSet<>();
        for (int i = 0; i < matrix.length; ++i)
            for (int j = 0; j < matrix[0].length; ++j)
                if (matrix[i][j] == 0) {
                    set1.add(i);
                    set2.add(j);
                }
        for (int row : set1)
            Arrays.fill(matrix[row], 0);
        for (int col : set2)
            for (int i = 0; i < matrix.length; ++i) matrix[i][col] = 0;
    }
}

class Solution2 {
    public void setZeroes(int[][] matrix) {
        boolean pre = false, cur = false;
        for (int i = 0; i < matrix.length; ++i) {
            pre = cur;
            cur = false;
            for (int j = 0; j < matrix[0].length; ++j) {
                if (matrix[i][j] == 0) {
                    cur = true;
                    for (int k = i - 1; k >= 0; --k) matrix[k][j] = 0;
                } else if (i > 0 && matrix[i - 1][j] == 0) matrix[i][j] = 0;
            }
            if (pre) for (int k = 0; k < matrix[0].length; ++k) matrix[i - 1][k] = 0;
        }
        if (cur) for (int k = 0; k < matrix[0].length; ++k) matrix[matrix.length - 1][k] = 0;
    }
}