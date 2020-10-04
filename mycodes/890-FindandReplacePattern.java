//Time 60%
//Space 97%

class Solution {
    public List<String> findAndReplacePattern(String[] words, String pattern) {
        Map<Character, Character> map = new HashMap<>();
        Set<Character> set = new HashSet<>();
        List<String> res = new ArrayList<>();
        for (String word : words) {
            int i = 0;
            for (i = 0; i < word.length(); ++i) {
                if (map.get(word.charAt(i)) == null) {
                    if (set.contains(pattern.charAt(i))) break;
                    map.put(word.charAt(i), pattern.charAt(i));
                    set.add(pattern.charAt(i));
                }
                if (map.get(word.charAt(i)) != pattern.charAt(i)) break;
            }
            if (i == word.length()) res.add(word);
            map.clear();
            set.clear();
        }
        return res;
    }
}