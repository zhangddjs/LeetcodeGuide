class Solution {
    public int minimumDeleteSum(String s1, String s2) {
        if ((s1 == null || s1.isEmpty()) && (s2 == null || s2.isEmpty())) return 0;
        if (s1 == null || s1.isEmpty()) return (int)s2.charAt(0) + minimumDeleteSum(s1, s2.substring(1));
        else if (s2 == null || s2.isEmpty())return (int)s1.charAt(0) + minimumDeleteSum(s1.substring(1), s2);
        if (s1.charAt(0) == s2.charAt(0)) return minimumDeleteSum(s1.substring(1), s2.substring(1));
        return Math.min((int)s2.charAt(0) + minimumDeleteSum(s1, s2.substring(1)), (int)s1.charAt(0) + minimumDeleteSum(s1.substring(1), s2));
    }
}

class Solution2 {
    int [][] mem;
    public int minimumDeleteSum(String s1, String s2) {
        if ((s1 == null || s1.isEmpty()) && (s2 == null || s2.isEmpty()) || s1.equals(s2)) return 0;
        mem = new int[s1.length() + 1][s2.length() + 1];
        for (int[] elm : mem) Arrays.fill(elm, -1);
        return dfs(s1, s2, 0, 0);
    }

    public int dfs(String s1, String s2, int p1, int p2) {
        if (p1 == s1.length() && p2 == s2.length()) return 0;
        if (p1 == s1.length())
            mem[p1][p2] = mem[p1][p2] == -1 ? 
                (int)s2.charAt(p2) + dfs(s1, s2, p1, p2 + 1) : mem[p1][p2];
        else if (p2 == s2.length())
            mem[p1][p2] = mem[p1][p2] == -1 ? 
                (int)s1.charAt(p1) + dfs(s1, s2, p1 + 1, p2) : mem[p1][p2];
        else if (s1.charAt(p1) == s2.charAt(p2))
            mem[p1][p2] = mem[p1][p2] == -1 ?
                dfs(s1, s2, p1 + 1, p2 + 1) : mem[p1][p2];
        if (mem[p1][p2] == -1)
            mem[p1][p2] = Math.min((int)s2.charAt(p2) + dfs(s1, s2, p1, p2 + 1),
                        (int)s1.charAt(p1) + dfs(s1, s2, p1 + 1, p2));
        return mem[p1][p2];
    }
}

class Solution3 {
    int [][] mem;
    public int minimumDeleteSum(String s1, String s2) {
        if ((s1 == null || s1.isEmpty()) && (s2 == null || s2.isEmpty()) || s1.equals(s2)) return 0;
        int M = s1.length(), N = s2.length();
        mem = new int[M + 1][N + 1];
        //init
        for (int[] elm : mem) Arrays.fill(elm, -1);
        mem[M][N] = 0;
        for (int i = M - 1; i >= 0; --i) mem[i][N] = mem[i + 1][N] + (int)s1.charAt(i);
        for (int j = N - 1; j >= 0; --j) mem[M][j] = mem[M][j + 1] + (int)s2.charAt(j);
        //iterate
        for (int i = M - 1; i >= 0; --i) {
            for (int j = N - 1; j >= 0; --j) {
                if (s1.charAt(i) == s2.charAt(j)) mem[i][j] = mem[i + 1][j + 1];
                else mem[i][j] = Math.min(s1.charAt(i) + mem[i + 1][j], s2.charAt(j) + mem[i][j + 1]);
            }
        }
        return mem[0][0];
    }
}