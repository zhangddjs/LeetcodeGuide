// 优化方向：先循环找出入度为0的，加入队列，然后每次遍历时确定入度为0的新的节点，继续加入队列。
//16% Time
//17% Space
class Solution {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        if (numCourses == 0 || prerequisites == null || prerequisites.length == 0 ||
            prerequisites[0].length == 0) return true;
        Map<Integer, Set<Integer>> mapOut = new HashMap<>();
        Map<Integer, Set<Integer>> mapIn = new HashMap<>();
        for (int[] prerequisite : prerequisites) {
            if (mapIn.get(prerequisite[0]) == null) mapIn.put(prerequisite[0], new HashSet<>());
            if (mapOut.get(prerequisite[1]) == null) mapOut.put(prerequisite[1], new HashSet<>());
            mapIn.get(prerequisite[0]).add(prerequisite[1]);
            mapOut.get(prerequisite[1]).add(prerequisite[0]);
        }
        while (!mapOut.isEmpty()) {
            boolean flag = false;
            Iterator<Map.Entry<Integer, Set<Integer>>> it = mapOut.entrySet().iterator();
            while (it.hasNext()) {
                Map.Entry<Integer, Set<Integer>> entry = it.next();
                if (!mapIn.containsKey(entry.getKey())) {
                    for (int to : entry.getValue()) {
                        mapIn.get(to).remove(entry.getKey());
                        if (mapIn.get(to).size() == 0) mapIn.remove(to);
                    }
                    it.remove();
                    flag = true;
                }
            }
            if (!flag) break;
        }
        return mapOut.isEmpty();
    }
}