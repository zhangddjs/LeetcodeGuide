class Solution {
    class MaxPoint {
        int max;
        int idx;
        MaxPoint(int max, int idx) {
            this.max = max;
            this.idx = idx;
        }
    }
    public int edgeScore(int[] edges) {
        int[] weight = new int[edges.length];
        MaxPoint p = new MaxPoint(0, 0);
        for (int i = 0; i < edges.length; i++) {
            weight[edges[i]] += i;
            if (weight[edges[i]] > p.max) p = new MaxPoint(weight[edges[i]], edges[i]);
            else if (weight[edges[i]] == p.max && edges[i] < p.idx) p.idx = edges[i];
        }
        return p.idx;
    }
}