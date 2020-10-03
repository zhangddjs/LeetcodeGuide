//暴力法
//MLE
class Solution1 {
    public String decodeAtIndex(String S, int K) {
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < S.length(); ++i) {
            if (!Character.isDigit(S.charAt(i))) res.append(S.charAt(i));
            else {
                String tmp = res.toString();
                for (int j = 0; j < S.charAt(i) - '0' - 1 && res.length() < K; ++j) res.append(tmp);
            }
            if (res.length() >= K) break;
        }
        return String.valueOf(res.charAt(K - 1));
    }
}


//分解子问题，规律
//Time 100%
//Space 87%
class Solution2 {
    public String decodeAtIndex(String S, int K) {
        int len = K;
        int [] lenAndIndex = new int[]{0,0};
        do {
            lenAndIndex = getLenAndIndex(S, len);
            len = len - lenAndIndex[0];
        } while (len != 0);
        return String.valueOf(S.charAt(lenAndIndex[1] - 1));
    }
    
    public int[] getLenAndIndex(String S, int K) {
        int len = 0, i;
        for (i = 0; i < S.length() && len < K; ++i) {
            if (!Character.isDigit(S.charAt(i))) len++;
            else if (S.charAt(i) - '0' <= (K - 1) / len) len *= S.charAt(i) - '0';
            else {
                len *= (K - 1) / len;
                break;
            }
        }
        return new int[]{len, i};
    }
}