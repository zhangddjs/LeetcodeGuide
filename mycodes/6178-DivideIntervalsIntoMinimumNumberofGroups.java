class Solution {
    public int minGroups(int[][] intervals) {
        int [] count = new int[1000002];
        int max = 1;
        for (int[] interval : intervals) {
            count[interval[0]]++;
            count[interval[1] + 1]--;
        }
        for (int i = 1; i < count.length; i++) {
            count[i] += count[i - 1];
            max = Math.max(max, count[i]);
        }
        return max;
    }
}