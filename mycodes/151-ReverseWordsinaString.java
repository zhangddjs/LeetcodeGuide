class Solution {
    public String reverseWords(String s) {
        String[] words = s.trim().split("\\s+");
        StringBuilder res = new StringBuilder();
        for (int i = words.length - 1; i >= 0; --i) res.append(words[i] + " ");
        return res.toString().trim();
    }
}

class Solution2 {
    public String reverseWords(String s) {
        StringBuilder res = new StringBuilder();
        for (String word : s.trim().split("\\s+")) res.insert(0, " " + word);
        return res.toString().trim();
    }
}