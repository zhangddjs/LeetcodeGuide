// TLE
class Solution {
    class Server {
        int idx;
        int weight;
        Server(int idx, int weight) {
            this.idx = idx;
            this.weight = weight;
        }
    }
    public int[] assignTasks(int[] servers, int[] tasks) {
        int[] res = new int[tasks.length];
        PriorityQueue<Server> q = new PriorityQueue<>((a, b) -> a.weight - b.weight != 0 ? a.weight - b.weight : a.idx - b.idx);
        for (int i = 0; i < servers.length; i++) q.offer(new Server(i, servers[i]));
        TreeMap<Integer, List<Server>> map = new TreeMap<>();
        int passedSeconds = 0;
        for (int i = 0; i < tasks.length; i++) {
            passedSeconds = Math.max(passedSeconds, i);
            TreeMap<Integer, List<Server>> tmpMap = new TreeMap<>();
            if (i >= passedSeconds) {
                for (Server s : map.getOrDefault(1, new ArrayList<>())) q.offer(s); 
                map.remove(1);
                for (int key : map.keySet()) tmpMap.put(key - 1, map.get(key));
                map = tmpMap;
                passedSeconds++;
            }
            if (q.isEmpty()) {
                tmpMap = new TreeMap<>();
                int k = map.keySet().iterator().next();
                for (Server s : map.get(k)) q.offer(s);
                map.remove(k);
                for (int key : map.keySet()) tmpMap.put(key - k, map.get(key));
                map = tmpMap;
                passedSeconds += k;
            }
            if (!map.containsKey(tasks[i])) map.put(tasks[i], new ArrayList<>());
            Server s = q.poll();
            map.get(tasks[i]).add(s);
            res[i] = s.idx;
        }
        return res;
    }
}