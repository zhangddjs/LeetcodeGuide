/**
 * O(n)
 * S(n)
 */

class Solution {
    public boolean areOccurrencesEqual(String s) {
        Map<Character, Integer> map = new HashMap<Character, Integer>();
        for (int i = 0; i < s.length(); i++) {
            map.put(s.charAt(i), map.getOrDefault(s.charAt(i), 0) + 1);
        }
        int val = -1;
        for (char k : map.keySet()) {
            if (val == -1) {
                val = map.get(k);
            } else if (val != map.get(k)) {
                return false;
            }
        }
        return true;
    }
}