class Solution {
    public int myAtoi(String str) {
        int i = 0;
        long buf = 0, sign = 1;
        while (i < str.length() && str.charAt(i) == ' ') i++;
        if(i == str.length()) return 0;
        else if(str.charAt(i) == '+' || str.charAt(i) == '-') sign = str.charAt(i++) == '+' ? 1 : -1;
        else if(!Character.isDigit(str.charAt(i))) return 0;
        while (i != str.length() && Character.isDigit(str.charAt(i)) && 
               sign * buf < Integer.MAX_VALUE && sign * buf > Integer.MIN_VALUE)
            buf = buf * 10 + (str.charAt(i++) - '0');
        if (sign * buf < Integer.MAX_VALUE && sign * buf > Integer.MIN_VALUE)
            return new Long(buf * sign).intValue();
        return sign * buf >= Integer.MAX_VALUE ? Integer.MAX_VALUE : Integer.MIN_VALUE;
    }
}