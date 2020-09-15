class Solution {
    public int numDistinctIslands(int[][] grid) {
        if (grid.length == 0 || grid[0].length == 0) return 0;
        List<List<Integer[]>> islands = new ArrayList<>();
        for (int i = 0; i < grid.length; ++i)
            for (int j = 0; j < grid[0].length; ++j)
                if (grid[i][j] == 1)
                    islands.add(dfs(grid, i, j, new ArrayList<Integer[]>()));
        int res = islands.size();
        for (int i = 0; i < islands.size(); ++i)
            for (int j = i + 1; j < islands.size(); ++j)
                if (judgeTwoIslands(grid, islands.get(i), islands.get(j))){
                    res--;
                    break;
                }
        return res;
    }

    public List dfs(int[][] grid, int i, int j, List<Integer[]> island) {
        if (i < 0 || j < 0 || i >= grid.length || j >= grid[0].length || grid[i][j] != 1) return island;
        grid[i][j] = 2;
        island.add(new Integer[]{i, j});
        dfs(grid, i - 1, j, island);
        dfs(grid, i, j - 1, island);
        dfs(grid, i, j + 1, island);
        return dfs(grid, i + 1, j, island);
    }

    public boolean judgeTwoIslands(int[][] grid, List<Integer[]> A, List<Integer[]> B) {
        if (A.size() != B.size()) return false;
        int [] offset = new int[]{B.get(0)[0] - A.get(0)[0], B.get(0)[1] - A.get(0)[1]};
        for (int i = 1; i < A.size(); ++i)
            if (B.get(i)[0] - A.get(i)[0] != offset[0] || 
                B.get(i)[1] - A.get(i)[1] != offset[1]) return false;
        return true;
    }
}