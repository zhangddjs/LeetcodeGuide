class Solution {
    public String nextClosestTime(String time) {
        TreeSet<Character> set = new TreeSet<>();
        for (int i = 0; i < 5; i++) if (i != 2) set.add(time.charAt(i));
        char smallest = set.iterator().next();
        for (char c : set) if (time.charAt(4) < c) return time.substring(0, 4) + c; //judge second
        for (char c : set) if (c < '6' && c > time.charAt(3)) return time.substring(0, 3) + c + smallest;   //judge minute
        for (char c : set) {
            if ((time.charAt(0) != '2' || (time.charAt(0) == '2' && c < '4')) && time.charAt(1) < c) 
                return time.substring(0,1) + c + ':' + smallest + smallest;
        }
        for (char c : set) if (time.charAt(0) < c && c <= '2') return "" + c + smallest + ':' + smallest + smallest;
        return "" + smallest + smallest + ':' + smallest + smallest;
    }
}