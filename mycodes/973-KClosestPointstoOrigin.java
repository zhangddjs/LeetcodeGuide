//Time 63%
//Space 8%
class Solution {
    public int[][] kClosest(int[][] points, int K) {
        if (K == points.length) return points;
        Arrays.sort(points, (a, b) -> (distance(a) - distance(b)));
        return Arrays.copyOf(points, K);
    }
    
    private int distance(int[] a) {
        return a[0] * a[0] + a[1] * a[1];
    }
}