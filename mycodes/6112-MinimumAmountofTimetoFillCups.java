class Solution {
    public int fillCups(int[] amount) {
        Arrays.sort(amount);
        if (amount[2] >= amount[0] + amount[1]) return amount[2];
        int res = 0;
        res += amount[1] - amount[0];
        amount[2] -= amount[1] - amount[0];
        amount[1] = amount[0];
        int tmp = amount[2] - amount[1];
        res += tmp * 2;
        amount[0] -= tmp;
        res += amount[0] % 2 == 0 ? 3 * amount[0] / 2 : 3 * (amount[0] / 2) + 2;
        return res;
    }
}