class ExamRoom {

    PriorityQueue<Integer[]> queue;
    int size;
    
    public ExamRoom(int N) {
        size = N;
        queue = new PriorityQueue<Integer[]>((a, b) -> {
            int minDisA = a[0] == -1 || a[1] == size ? a[1] - a[0] - 1 : (a[1] - a[0]) / 2;
            int minDisB = b[0] == -1 || b[1] == size ? b[1] - b[0] - 1 : (b[1] - b[0]) / 2;
            if (minDisA == minDisB) return a[0] - b[0];
            else return minDisB - minDisA;
        });
        queue.offer(new Integer[]{-1, size});
    }
    
    public int seat() {
        if (queue.isEmpty()) return -1;
        Integer [] pos = queue.poll();
        int index = -1;
        if (pos[0] == -1) {
            index = 0;
            queue.offer(new Integer[]{index, pos[1]});
        } else if (pos[1] == size) {
            index = pos[1] - 1;
            queue.offer(new Integer[]{pos[0], index});
        } else {
            index = pos[0] + (pos[1] - pos[0]) / 2;
            queue.offer(new Integer[]{pos[0], index});
            queue.offer(new Integer[]{index, pos[1]});
        }
        return index;
    }
    
    public void leave(int p) {
        Integer[] left = null, right = null;
        for (Integer[] pos : queue) {
            if (pos[0] == p) right = pos;
            else if (pos[1] == p) left = pos;
        }
        Integer[] pos = new Integer[2];
        pos[0] = left == null ? -1 : left[0];
        pos[1] = right == null ? size : right[1];
        queue.offer(pos);
        queue.remove(left);
        queue.remove(right);
    }
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * ExamRoom obj = new ExamRoom(N);
 * int param_1 = obj.seat();
 * obj.leave(p);
 */