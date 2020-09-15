class Solution {
    public List<List<Integer>> largeGroupPositions(String S) {
        int i = 0, j = 1;
        List<List<Integer>> res = new ArrayList<>();
        while (j < S.length()) {
            if (S.charAt(j) != S.charAt(i)) {
                if(j - i >= 3) res.add(Arrays.asList(i, j - 1));
                i = j;
            }
            j++;
        }
        if(j - i >= 3) res.add(Arrays.asList(i, j - 1));
        return res;
    }
}