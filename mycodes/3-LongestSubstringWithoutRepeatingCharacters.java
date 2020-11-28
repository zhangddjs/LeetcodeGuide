//Time 90.72%
//Space 5.69%
//slide window

class Solution {
    public int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> map = new HashMap<>();
        int res = 0, left = 0;
        for (int i = 0; i < s.length(); ++i) {
            if (map.containsKey(s.charAt(i))) {
                res = Math.max(i - left, res);
                left = Math.max(map.get(s.charAt(i)) + 1, left);
            }
            map.put(s.charAt(i), i);
        }
        return Math.max(res, s.length() - left);
    }
}

//Time 90%
//Space 84%
class Solution1122 {
    public int lengthOfLongestSubstring(String s) {
        if (s == null || s.length() == 0) return 0;
        Map<Character, Integer> map = new HashMap<>();
        int left = 0, max = 0, i;
        for (i = 0; i < s.length(); ++i) {
            char c = s.charAt(i);
            if (map.getOrDefault(c, -1) < left) map.put(c, i);
            else {
                max = Math.max(max, i - left);
                left = map.get(c) + 1;
                map.put(c, i);
            }
        }
        return Math.max(max, i - left);
    }
}