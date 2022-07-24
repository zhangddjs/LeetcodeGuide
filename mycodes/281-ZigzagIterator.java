public class ZigzagIterator {
    int round;
    int idx;
    List<Integer>[] v;

    public ZigzagIterator(List<Integer> v1, List<Integer> v2) {
        v = (ArrayList<Integer> []) new ArrayList[2];
        v[0] = v1;
        v[1] = v2;
        idx = v1.size() == 0 ? 1 : 0;
        round = 0;
    }

    public int next() {
        int res = v[idx].get(round);
        int nextIdx = idx + 1 & 1;
        if ((nextIdx & 1) == 0 || v[nextIdx].size() <= round) round++;
        if (v[nextIdx].size() > round) idx = nextIdx;
        return res;
    }

    public boolean hasNext() {
        return v[idx].size() > round;
    }
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * ZigzagIterator i = new ZigzagIterator(v1, v2);
 * while (i.hasNext()) v[f()] = i.next();
 */