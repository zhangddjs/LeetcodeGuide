class Solution {
    public String toGoatLatin(String S) {
        String [] words = S.split(" ");
        Set<Character> vowels = new HashSet<>();
        Collections.addAll(vowels, new Character[]{'a','e','i','o','u', 'A', 'E', 'I', 'O', 'U'});
        StringBuilder suffix = new StringBuilder("ma");
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < words.length; ++i) {
            suffix.append("a");
            StringBuilder wordBuilder = new StringBuilder(words[i]);
            if (!vowels.contains(words[i].charAt(0))) {
                wordBuilder.deleteCharAt(0);
                wordBuilder.append(words[i].charAt(0));
            }
            wordBuilder.append(suffix);
            res.append(wordBuilder);
            if(i != words.length - 1) res.append(" ");
        }
        return res.toString();
    }
}

class Solution2 {
    public String toGoatLatin(String S) {
        String [] words = S.split(" ");
        Set<Character> vowels = new HashSet<>();
        Collections.addAll(vowels, new Character[]{'a','e','i','o','u', 'A', 'E', 'I', 'O', 'U'});
        String suffix = "ma";
        String res = "";
        for (int i = 0; i < words.length; ++i) {
            suffix += "a";
            if (!vowels.contains(words[i].charAt(0)))
                words[i] = words[i].substring(1, words[i].length()) + words[i].charAt(0);
            words[i] += suffix;
        }
        for (String word : words)
            res += word + " ";
        return res.substring(0, res.length() - 1);
    }
}

class Solution3 {
    public String toGoatLatin(String S) {
        Set<Character> vowels = new HashSet<>();
        Collections.addAll(vowels, new Character[]{'a','e','i','o','u', 'A', 'E', 'I', 'O', 'U'});
        StringBuilder suffix = new StringBuilder("ma");
        StringBuilder res = new StringBuilder();
        StringBuilder word = new StringBuilder();
        S += " ";
        for (char c : S.toCharArray()) {
            if (Character.isLetter(c)){
                word.append(c);
                continue;
            }
            suffix.append('a');
            if (!vowels.contains(word.charAt(0)))
                word.append(word.charAt(0)).deleteCharAt(0);
            res.append(word).append(suffix).append(" ");
            word = new StringBuilder();
        }
        return res.deleteCharAt(res.length() - 1).toString();
    }
}