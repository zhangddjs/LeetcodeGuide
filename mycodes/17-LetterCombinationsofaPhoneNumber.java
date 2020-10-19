class Solution {
    Map<Character, String> map = new HashMap<>();
    public List<String> letterCombinations(String digits) {
        init();
        List<String> res = new ArrayList<>();
        helper(digits, "", res);
        return res;
    }

    public void init() {
        map.put('2', "abc");
        map.put('3', "def");
        map.put('4', "ghi");
        map.put('5', "jkl");
        map.put('6', "mno");
        map.put('7', "pqrs");
        map.put('8', "tuv");
        map.put('9', "wxyz");
    }

    public void helper(String digits, String path, List<String> res) {
        if (digits == null || digits.equals("")) {
            if (!path.equals("")) res.add(path);
            return;
        }
        String nextDigits = digits.substring(1);
        char cur = digits.charAt(0);
        helper(nextDigits, path + map.get(cur).charAt(0), res);
        helper(nextDigits, path + map.get(cur).charAt(1), res);
        helper(nextDigits, path + map.get(cur).charAt(2), res);
        if (cur == '7' || cur == '9')
            helper(nextDigits, path + map.get(cur).charAt(3), res);
    }
}