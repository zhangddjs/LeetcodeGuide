class Solution {
    public boolean buddyStrings(String A, String B) {
        if(A.length() != B.length()) return false;  //case 1
        char difA = '\0', difB = '\0';
        Set<Character> set = new HashSet<>();
        int count = 0;
        boolean dup = false;
        for (int i = 0; i < A.length(); ++i) {
            if(set.remove(A.charAt(i))) dup = true;
            set.add(A.charAt(i));
            if (A.charAt(i) != B.charAt(i)) {
                count++;
                if (count > 2) return false;    //case 2
                else if (count == 1) {
                    difA = A.charAt(i);
                    difB = B.charAt(i);
                } else if (difA != B.charAt(i) || difB != A.charAt(i)) return false;     //case 3
            }
        }
        return count == 2 || (count == 0 && dup);   //case 4
    }
}