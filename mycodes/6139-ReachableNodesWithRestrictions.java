class Solution {
    public int reachableNodes(int n, int[][] edges, int[] restricted) {
        Map<Integer, Set<Integer>> tree = buildTree(edges);
        boolean[] visited = new boolean[n];
        for (int r : restricted) visited[r] = true;
        return dfs(tree, 0, visited);
    }
    
    private Map<Integer, Set<Integer>> buildTree(int[][] edges) {
        Map<Integer, Set<Integer>> tree = new HashMap<>();
        for (int[] edge : edges) {
            if (!tree.containsKey(edge[0])) tree.put(edge[0], new HashSet<Integer>());
            if (!tree.containsKey(edge[1])) tree.put(edge[1], new HashSet<Integer>());
            tree.get(edge[0]).add(edge[1]);
            tree.get(edge[1]).add(edge[0]);
        }
        return tree;
    }

    private int dfs(Map<Integer, Set<Integer>> tree, int root, boolean[] visited) {
        if (tree.isEmpty() || visited[root]) return 0;
        int cnt = 1;
        visited[root] = true;
        for (int neighbor : tree.get(root)) cnt += dfs(tree, neighbor, visited);
        return cnt;
    }
}