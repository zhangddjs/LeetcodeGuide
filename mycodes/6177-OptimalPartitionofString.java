class Solution {
    public int partitionString(String s) {
        Set<Character> set = new HashSet<>();
        int res = 1;
        for (char c : s.toCharArray()) {
            if (set.contains(c)) {
                res++;
                set = new HashSet<>();
            }
            set.add(c);
        }
        return res;
    }
}