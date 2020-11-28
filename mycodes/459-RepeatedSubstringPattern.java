//11% Time
//94% Space
//O(n^2)
class Solution {
    public boolean repeatedSubstringPattern(String s) {
        if (s == null || s.length() < 2) return false;
        int patternLen = 1;
        for (int i = 1; i < s.length(); ++i) {
            if (s.charAt(i) != s.charAt(i % patternLen)) {
                patternLen++;
                if (patternLen > (s.length() >>> 1)) return false;
                i = patternLen - 1;
            }
        }
        return s.length() % patternLen == 0;
    }
}

//12% Time
//14% Space
class Solution2 {
    public boolean repeatedSubstringPattern(String s) {
        if (s == null || s.length() < 2) return false;
        for (int i = 1; i <= (s.length() >>> 1); ++i) {
            if (s.equals(s.substring(i) + s.substring(0, i))) return true;
        }
        return false;
    }
}

