class MyCalendarTwo {
    List<Integer[]> intervals1;
    List<Integer[]> intervals2;
    public MyCalendarTwo() {
        intervals1 = new ArrayList<>();
        intervals2 = new ArrayList<>();
    }
    
    public boolean book(int start, int end) {
        for (Integer[] interval2 : intervals2) if (isOverlap(start, end, interval2)) return false;
        for (Integer[] interval1 : intervals1) {
            if (isOverlap(start, end, interval1)) {
                Integer[] overlap = getOverlapInterval(start, end, interval1);
                intervals2.add(overlap);
            }
        }
        intervals1.add(new Integer[]{start, end});
        return true;
    }

    private boolean isOverlap(int start, int end, Integer[] interval) {
        return start < interval[1] && end > interval[0];
    }

    private Integer[] getOverlapInterval(int start, int end, Integer[] interval) {
        return new Integer[]{Math.max(start, interval[0]), Math.min(end, interval[1])};
    }
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * MyCalendarTwo obj = new MyCalendarTwo();
 * boolean param_1 = obj.book(start,end);
 */