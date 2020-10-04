//TLE
class Solution {
    public int subarrayBitwiseORs(int[] A) {
        int [] num = new int[A.length];
        Set<Integer> set = new HashSet<>();
        for (int len = 0; len < A.length; ++len) {
            for (int i = 0; i < A.length - len; ++i) {
                num[i] |= A[i + len];
                set.add(num[i]);
            }
        }
        return set.size();
    }
}