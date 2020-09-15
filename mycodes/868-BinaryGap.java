class Solution {
    public int binaryGap(int N) {
        int res = 0, tmp = 1;
        boolean flag = false;
        while (N != 0) {
            if(N % 2 != 0 && !flag) flag = true;    //first '1' found
            else if (N % 2 != 0 && flag) {
                res = Math.max(tmp, res);
                tmp = 1;
            } else if(N % 2 == 0 && flag) tmp++;    //if cur is '0', second pointer + 1.
            N /= 2;
        }
        return res;
    }
}

