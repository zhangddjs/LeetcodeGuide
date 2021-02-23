//23% Time
// 6% Space
class Solution {
    public String intToRoman(int num) {
        Character [] roman = new Character[]{'I', 'V', 'X', 'L', 'C', 'D', 'M'};
        int cur = 0;
        StringBuilder res = new StringBuilder();
        while (num != 0) {
            StringBuilder tmpstr = new StringBuilder();
            int tmp = num % 10;
            if (tmp == 9 || tmp == 4) {
                int offset = tmp == 4 ? 1 : 2;
                tmpstr.append(roman[cur]).append(roman[cur + offset]);
            } else {
                if (tmp >= 5) {
                    tmpstr.append(roman[cur + 1]);
                    tmp -= 5;
                }
                while (tmp-- > 0) {
                    tmpstr.append(roman[cur]);
                }
            }
            cur += 2;
            num /= 10;
            res.append(tmpstr.reverse());
        }
        return res.reverse().toString();
    }
}