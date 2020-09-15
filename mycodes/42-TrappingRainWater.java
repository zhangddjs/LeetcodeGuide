class Solution {
    public int trap(int[] height) {
        if (height == null || height.length == 0) return 0;
        int max = height[0], maxIndex = 0;
        for (int i = 1; i < height.length; ++i) {
            if (max < height[i]) {
                max = height[i];
                maxIndex = i;
            }
        }
        int maximal = 0, res = 0;
        for (int i = 0; i < maxIndex; ++i) {
            if (height[i] < maximal) res += maximal - height[i];
            else maximal = height[i];
        }
        maximal = 0;
        for (int i = height.length - 1; i > maxIndex; --i) {
            if (height[i] < maximal) res += maximal - height[i];
            else maximal = height[i];
        }
        return res;
    }
}