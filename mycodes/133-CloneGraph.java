/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> neighbors;

    public Node() {
        val = 0;
        neighbors = new ArrayList<Node>();
    }

    public Node(int _val) {
        val = _val;
        neighbors = new ArrayList<Node>();
    }

    public Node(int _val, ArrayList<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
}
*/

class Solution1 {
    public Node cloneGraph(Node node) {
        if (node == null) return null;
        Node res = new Node(node.val);
        Set<Node> visited = new HashSet<>();
        Map<Node, Node> map = new HashMap<>();
        Node cur = res;
        map.put(node, cur);
        dfs(cur, visited, map, node.neighbors);
        return res;
    }

    public void dfs(Node cur, Set<Node> visited, Map<Node, Node> map, List<Node> neighbors) {
        if (visited.contains(cur)) return;
        visited.add(cur);
        for (Node neighbor : neighbors) {
            if (map.get(neighbor) == null) map.put(neighbor, new Node(neighbor.val));
            cur.neighbors.add(map.get(neighbor));
            dfs(map.get(neighbor), visited, map, neighbor.neighbors);
        }
    }
}

class Solution2 {
    public Node cloneGraph(Node node) {
        if (node == null) return null;
        Node res = new Node(node.val);
        Set<Node> visited = new HashSet<>();
        Map<Node, Node> map = new HashMap<>();
        Queue<Node> queue = new LinkedList<>();
        queue.offer(node);
        map.put(node, res);
        visited.add(node);
        while (!queue.isEmpty()) {
            node = queue.poll();
            Node cur = map.get(node);
            for (Node neighbor : node.neighbors) {
                if (!visited.contains(neighbor)) {
                    queue.offer(neighbor);
                    visited.add(neighbor);
                }
                if (map.get(neighbor) == null) map.put(neighbor, new Node(neighbor.val));
                cur.neighbors.add(map.get(neighbor));
            }
        }

        return res;
    }
}