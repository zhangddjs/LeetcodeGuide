class Solution {
    public List<List<Integer>> removeInterval(int[][] intervals, int[] toBeRemoved) {
        List<List<Integer>> res = new ArrayList<>();
        if (intervals.length == 0 || toBeRemoved.length != 2) return res;
        for (int i = 0; i < intervals.length; i++) {
            int begin = intervals[i][0], end = intervals[i][1];
            if (end <= toBeRemoved[0] || begin >= toBeRemoved[1]) res.add(IntStream.of(intervals[i]).boxed().collect(Collectors.toList()));
            else {
                if (begin < toBeRemoved[0]) res.add(Arrays.asList(intervals[i][0], toBeRemoved[0]));
                if (end > toBeRemoved[1]) res.add(Arrays.asList(toBeRemoved[1], intervals[i][1]));
            }
        }
        return res;
    }
}