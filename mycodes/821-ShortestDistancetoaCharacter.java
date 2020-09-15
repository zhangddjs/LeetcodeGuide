class Solution {
    public int[] shortestToChar(String S, char C) {
        int j = S.indexOf(C), i = -j;
        int [] res = new int[S.length()];
        for(int k = 0; k < S.length(); k++){
            res[k] = Math.min(j - k, k - i);
            if(k == j){
                i = j;
                j = S.indexOf(C, j + 1);
                j = j < 0 ? Integer.MAX_VALUE : j;
            }
        }
        return res;
    }
}