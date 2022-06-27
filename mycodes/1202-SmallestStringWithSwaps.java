class Solution {
    public String smallestStringWithSwaps(String s, List<List<Integer>> pairs) {
        Map<Integer, Set<Integer>> graph = buildGraph(pairs);
        List<TreeSet<Integer>> subGraphs = getConnectedSubGraphs(graph);
        return sortString(s, subGraphs);
    }

    private Map<Integer, Set<Integer>> buildGraph(List<List<Integer>> pairs) {
        Map<Integer, Set<Integer>> graph = new HashMap<Integer, Set<Integer>>();
        for (List<Integer> pair : pairs) {
            if (pair.get(0) == pair.get(1)) continue;
            if (!graph.containsKey(pair.get(0))) {
                graph.put(pair.get(0), new HashSet<Integer>());
            }
            if (!graph.containsKey(pair.get(1))) {
                graph.put(pair.get(1), new HashSet<Integer>());
            }
            graph.get(pair.get(0)).add(pair.get(1));
            graph.get(pair.get(1)).add(pair.get(0));
        }
        return graph;
    }

    private List<TreeSet<Integer>> getConnectedSubGraphs(Map<Integer, Set<Integer>> graph) {
        List<TreeSet<Integer>> list = new ArrayList<TreeSet<Integer>>();
        while (!graph.isEmpty()) {
            TreeSet<Integer> set = new TreeSet<Integer>();
            Queue<Integer> queue = new LinkedList<Integer>();
            int key = graph.keySet().iterator().next();
            queue.offer(key);
            set.add(key);
            while (!queue.isEmpty()) {
                int k = queue.poll();
                for (int v : graph.get(k)) {
                    if (set.contains(v)) continue;
                    set.add(v);
                    queue.offer(v);
                }
                graph.remove(k);
            }
            list.add(set);
        }
        return list;
    }

    private String sortString(String s, List<TreeSet<Integer>> subGraphs) {
        char[] chs = s.toCharArray();
        for (TreeSet<Integer> vertices : subGraphs) {
            Queue<Character> q = new PriorityQueue<>();
            for (int v : vertices) {
                q.offer(chs[v]);
            }
            for (int v : vertices) {
                chs[v] = q.poll();
            }
        }
        return String.valueOf(chs);
    }
}