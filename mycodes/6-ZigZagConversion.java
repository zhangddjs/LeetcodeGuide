// 76% Time
// 86%

class Solution {
    public String convert(String s, int numRows) {
        List<StringBuilder> list = new ArrayList<>();
        int i = 0;
        for (int j = 0; j < numRows; ++j) {
            list.add(new StringBuilder());
        }
        while (i < s.length()) {
            for (int j = 0; j < numRows && i < s.length(); ++j) {
                list.get(j).append(s.charAt(i++));
            }
            for (int j = numRows - 2; j > 0 && i < s.length(); --j) {
                list.get(j).append(s.charAt(i++));
            }
        }
        StringBuilder res = new StringBuilder();
        for (StringBuilder elm : list) {
            res.append(elm);
        }
        return res.toString();
    }
}