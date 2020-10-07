//Time 41%
//Space 96%
class Solution {
    Map<List<Integer>, Integer> map = new HashMap<>();
    Map<List<Integer>, Set<List<Integer>>> neighbors = new HashMap<>();
    public int largestIsland(int[][] grid) {
        //dfs遍历某个岛屿，记录大小，记录入口，访问标记记录外面一圈水面，该位置可能同时属于多个岛屿，若置为1，则这些岛屿连接。
        //用map维护每个位置代表的置1后的更新面积。
        //遍历其它岛屿时，将外圈水面累加(非0时减1)，哨兵记录最大值。
        for (int i = 0; i < grid.length; ++i)
            for (int j = 0; j < grid[0].length; ++j)
                if (grid[i][j] == 1)
                    dfs(grid, Arrays.asList(new Integer[]{i, j}), i, j);
        return getMax();
    }
    
    public void dfs(int[][] grid, List<Integer> island, int i, int j) {
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] == 2) return;
        if (grid[i][j] == 0) {
            List<Integer> cur = Arrays.asList(new Integer[]{i, j});
            if (neighbors.get(cur) == null) neighbors.put(cur, new HashSet<List<Integer>>());
            neighbors.get(cur).add(island);
            return;
        }
        grid[i][j] = 2;
        map.put(island, map.getOrDefault(island, 0) + 1);
        dfs(grid, island, i + 1, j);
        dfs(grid, island, i - 1, j);
        dfs(grid, island, i, j + 1);
        dfs(grid, island, i, j - 1);
    }
    
    public int getMax() {
        int max = 1;
        for (Set<List<Integer>> set : neighbors.values()) {
            int tmp = 1;
            for (List<Integer> island : set)
                tmp += map.get(island);
            max = Math.max(max, tmp);
        }
        max = Math.max(max, map.getOrDefault(Arrays.asList(new Integer[]{0, 0}), 1));
        return max;
    }
}