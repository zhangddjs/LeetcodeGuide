//Time 100%
//Space 98%
class Solution {
    public int[][] spiralMatrixIII(int R, int C, int r0, int c0) {
        //每走半个时钟，长度增加1
        //四个方向，0代表东，1南，2西，3北
        int len = 0, dest = -1, i = 0;
        int [][] res = new int[R * C][2];
        int [] rc = new int[]{r0, c0};
        while (i < R * C) {
            len++;
            dest = (dest + 1) % 4;
            i = traverse(res, len, dest, rc, i, R, C);
            dest = (dest + 1) % 4;
            i = traverse(res, len, dest, rc, i, R, C);
        }
        return res;
    }
    
    public int traverse(int [][] res, int len, int dest, int[] rc, int i, int R, int C) {
        int remain = len;
        while (remain != 0 && i < res.length) {
            //visit
            //walk
            //continue;
            if (rc[0] < R && rc[0] >= 0 && rc[1] < C && rc[1] >= 0) res[i++] = new int[]{rc[0], rc[1]};
            rc[1] += dest == 0 ? 1 : dest == 2 ? -1 : 0;
            rc[0] += dest == 1 ? 1 : dest == 3 ? -1 : 0;
            remain--;
        }
        return i;
    }
}