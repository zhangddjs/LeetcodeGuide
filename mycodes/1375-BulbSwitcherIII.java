//13% Time
//8% Space
//存在S(1)解
class Solution {
    public int numTimesAllBlue(int[] light) {
        if (light == null || light.length == 0) return 0;
        Set<Integer> set = new HashSet<>();
        int p = 1, cnt = 0;
        for (int v : light) {
            set.add(v);
            while (set.contains(p)) {
                set.remove(p++);
            }
            if (set.isEmpty()) cnt++;
        }
        return cnt;
    }
}

//100% Time
//39% Space
//S(1)
class Solution2 {
    public int numTimesAllBlue(int[] light) {
        if (light == null || light.length == 0) return 0;
        int max = 0x80000000, cnt = 0;
        for (int i = 0; i < light.length; ++i) {
            max = Math.max(max, light[i]);
            if (i + 1 == max) cnt++;
        }
        return cnt;
    }
}