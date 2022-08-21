class Solution {
    int INF = 2147483647, m, n;

    public void wallsAndGates(int[][] rooms) {
        m = rooms.length;
        n = rooms[0].length;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (rooms[i][j] == 0) {
                    dfs(rooms, i - 1, j, 1);
                    dfs(rooms, i + 1, j, 1);
                    dfs(rooms, i, j - 1, 1);
                    dfs(rooms, i, j + 1, 1);
                }
            }
        }
    }

    private void dfs(int[][] rooms, int i, int j, int dis) {
        if (i < 0 || i >= m || j < 0 || j >= n || rooms[i][j] == -1 ||
            rooms[i][j] == 0 || (rooms[i][j] != INF && rooms[i][j] <= dis)) {
            return;
        }
        rooms[i][j] = dis;
        dfs(rooms, i - 1, j, dis + 1);
        dfs(rooms, i + 1, j, dis + 1);
        dfs(rooms, i, j - 1, dis + 1);
        dfs(rooms, i, j + 1, dis + 1);
    }
}