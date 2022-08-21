class Solution {
    public String smallestNumber(String pattern) {
        StringBuilder res = new StringBuilder();
        StringBuilder temp = new StringBuilder();
        char ch = '1';
        for (int i = 0; i < pattern.length(); i++) {
            char cur = pattern.charAt(i);
            temp.append(ch++);
            if (cur == 'I') {
                res.append(temp.reverse());
                temp = new StringBuilder();
            }
        }
        temp.append(ch++);
        res.append(temp.reverse());
        return res.toString();
    }
}