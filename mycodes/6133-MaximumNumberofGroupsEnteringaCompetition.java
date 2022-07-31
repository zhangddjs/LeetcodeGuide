class Solution {
    public int maximumGroups(int[] grades) {
        int acc = 0, i = 1;
        for (i = 1; acc + i <= grades.length; i++) acc += i;
        return i - 1;;
    }
}