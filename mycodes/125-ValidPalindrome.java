class Solution {
    public boolean isPalindrome(String s) {
        int p1 = 0, p2 = s.length() - 1;
        while (p1 < p2) {
            while(p1 < p2 && !(Character.isLetter(s.charAt(p1)) || Character.isDigit(s.charAt(p1)))) p1++;  //Character.isLetterOrDigit
            while(p1 < p2 && !(Character.isLetter(s.charAt(p2)) || Character.isDigit(s.charAt(p2)))) p2--;
            if (Character.toLowerCase(s.charAt(p1)) == Character.toLowerCase(s.charAt(p2))) {
                p1++;
                p2--;
            } else return false;
        }
        return true;
    }
}