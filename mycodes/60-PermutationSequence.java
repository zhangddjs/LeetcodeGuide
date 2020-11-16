class Solution {
    public String getPermutation(int n, int k) {
        if (n == 1) return "1";
        StringBuilder sb = new StringBuilder();
        List<Character> lst = new LinkedList<>();
        int mult = 1;
        lst.add('1');
        for (int i = 2; i < n; ++i) {
            mult *= i;
            lst.add((char)(i + '0'));
        }
        lst.add((char)(n + '0'));
        n--;
        k--;
        while (mult != 1) {
            int cur = k / mult;
            sb.append(lst.get(cur));
            lst.remove(cur);
            k %= mult;
            mult /= n--;
        }
        sb.append(lst.get(k));
        sb.append(lst.get(1 - k));
        return sb.toString();
    }
}