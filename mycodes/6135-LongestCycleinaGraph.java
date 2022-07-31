class Solution {
    public int longestCycle(int[] edges) {
        Map<Integer, Integer> map = new HashMap<>();
        Set<Integer> cur;
        boolean[] visited = new boolean[edges.length];
        int max = -1;
        for (int i = 0; i < edges.length; i++) {
            if (visited[i]) continue;
            cur = new HashSet<>();
            cur.add(i);
            map.put(i, 0);
            visited[i] = true;
            int j = edges[i], idx = 1;
            while (!cur.contains(j) && j != -1 && !visited[j]) {
                visited[j] = true;
                cur.add(j);
                map.put(j, idx++);
                j = edges[j];
            }
            if (cur.contains(j)) {
                max = Math.max(max, idx - map.get(j));
            }
        }
        return max;
    }
}

