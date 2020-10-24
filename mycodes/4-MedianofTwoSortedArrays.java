//Brute Force
//Time 100%
//Space 15%

class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int len1 = nums1.length, len2 = nums2.length, i = 0, j = 0,  pre = 0, cur = 0;
        while (i + j != (len1 + len2) / 2 + 1) {
            pre = cur;
            if (j == len2 || (i != len1 && nums1[i] < nums2[j])) cur = nums1[i++];
            else if (i == len1 || (j != len2 && nums1[i] >= nums2[j])) cur = nums2[j++];
        }
        return ((len1 + len2) & 1) == 0 ? (pre + cur) / 2.0 : cur;
    }
}