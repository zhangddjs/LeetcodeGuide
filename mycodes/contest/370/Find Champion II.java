class Solution {
    public int findChampion(int n, int[][] edges) {
        int[] ingress = new int[n];
        for (int i = 0; i < edges.length; i++) {
            ingress[edges[i][1]]++;
        }
        int champion = -1;
        for (int i = 0; i < n; i++) {
            if (ingress[i] == 0) {
                if (champion != -1) {
                    return -1;
                }
                champion = i;
            }
        }
        return champion;
    }
}