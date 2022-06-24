class Solution {
    public List<String> printVertically(String s) {
        String[] words = s.split(" ");
        List<String> res = new ArrayList<String>();
        int maxLen = 0;
        for (String word : words) maxLen = Math.max(maxLen, word.length());
        for (int i = 0; i < maxLen; i++) {
            StringBuilder str = new StringBuilder();
            for (int j = 0; j < words.length; j++) {
                if (i >= words[j].length()) str.append(" ");
                else str.append(words[j].charAt(i));
            }
            res.add(trimTrailSpace(str).toString());
        }
        return res;
    }

    private StringBuilder trimTrailSpace(StringBuilder str) {
        int i = str.length() - 1;
        for (; i >= 0 && str.charAt(i) == ' '; i--);
        str.delete(i + 1, str.length());
        return str;
    }
}