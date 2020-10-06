//Time 5%
//Space 23%
//BFS
class Solution {
    public int numBusesToDestination(int[][] routes, int S, int T) {
        if (S == T) return 0;
        Map<Integer, Set<Set<Integer>>> map = new HashMap<>();
        for (int[] route : routes) {
            Set<Integer> set = Arrays.stream(route).boxed().collect(Collectors.toSet());
            for (int ele : route) {
                if (map.get(ele) == null) map.put(ele, new HashSet<Set<Integer>>());
                map.get(ele).add(set);
            }
        }
        if (!map.containsKey(S) || !map.containsKey(T)) return -1;
        Set<Integer> visited = new HashSet<>();
        return getCount(visited, map, S, T);
    }
    
    public int getCount(Set<Integer> visited, Map<Integer, Set<Set<Integer>>> map, int S, int T) {
        Set<Set<Integer>> visitedSet = new HashSet<>();
        int cnt = 1;
        Queue<Integer> queue = new LinkedList<>();
        for (Set<Integer> set : map.get(S)) {
            for (int ele : set) {
                queue.offer(ele);
            }
        }
        int size = queue.size();
        while (!visited.contains(T) && !queue.isEmpty()) {
            int ele = queue.poll();
            visited.add(ele);
            Set<Set<Integer>> sets = map.get(ele);
            for (Set<Integer> set : sets) {
                if (visitedSet.contains(set)) continue;
                for (int neighbor : set) {
                    if (visited.contains(neighbor)) continue;
                    queue.offer(neighbor);
                }
                visitedSet.add(set);
            }
            size--;
            if (size == 0 && !visited.contains(T)) {
                size = queue.size();
                cnt++;
            }
        }
        return visited.contains(T) ? cnt : -1;
    }
}