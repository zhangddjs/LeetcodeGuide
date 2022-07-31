// WA
class Solution {
    public int closestMeetingNode(int[] edges, int node1, int node2) {
        boolean[] visited = new boolean[edges.length];
        int edge1 = edges[node1], edge2 = edges[node2];
        Map<Integer, Integer> map1 = new HashMap<>();
        Map<Integer, Integer> map2 = new HashMap<>();
        TreeSet<Integer> set = new TreeSet<>();
        int dis1 = 1, dis2 = 1;
        map1.put(node1, 0);
        map2.put(node2, 0);
        while (edge1 != -1 && !visited[edge1]) {
            visited[edge1] = true;
            map1.put(edge1, dis1++);
            edge1 = edges[edge1];
        }
        visited = new boolean[edges.length];
        while (edge2 != -1 && !visited[edge2]) {
            visited[edge2] = true;
            map1.put(edge2, dis2++);
            edge2 = edges[edge2];
        }
        int min = Integer.MAX_VALUE;
        for (int k : map1.keySet()) {
            if (map2.containsKey(k)) {
                int dis = map1.get(k) + map2.get(k);
                if (dis < min) {
                    set = new TreeSet<>();
                    min = dis;
                }
                if (dis == min) {
                    set.add(k);
                }
            }
        }
        return set.isEmpty() ? -1 : set.iterator().next();
    }
}