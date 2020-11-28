//Time 5%
//Space 5%
class Solution {
    public String minWindow(String s, String t) {
        if (t.length() == 1) return s.contains(t) ? t : "";
        TreeSet<Integer> indexSet = new TreeSet<>();
        Map<Character, TreeSet<Integer>> map = new HashMap<>();
        Map<Character, Integer> charCnt = new HashMap<>();
        for (char c : t.toCharArray()) {
            charCnt.put(c, charCnt.getOrDefault(c, 0) + 1);
            map.put(c, new TreeSet<>());
        }
        int cnt = 0, start = 0, end = 0, minLen = s.length() + 1;
        for (int i = 0; i < s.length(); ++i) {
            if (charCnt.get(s.charAt(i)) == null) continue;
            if (charCnt.get(s.charAt(i)) > 0) {
                cnt++;
                charCnt.put(s.charAt(i), charCnt.get(s.charAt(i)) - 1);
                map.get(s.charAt(i)).add(i);
                indexSet.add(i);
            } else {
                int rmvIdx = map.get(s.charAt(i)).first();
                map.get(s.charAt(i)).remove(rmvIdx);
                indexSet.remove(rmvIdx);
                map.get(s.charAt(i)).add(i);
                indexSet.add(i);
            }
            if (cnt == t.length()) {
                if (minLen > i - indexSet.first() + 1) {
                    minLen = i - indexSet.first() + 1;
                    start = indexSet.first();
                    end = i + 1;
                }
            }
        }
        return s.substring(start, end);
    }
}