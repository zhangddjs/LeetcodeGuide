class Solution {
    public boolean carPooling(int[][] trips, int capacity) {
        TreeMap<Integer, Integer> up = new TreeMap<>();
        TreeMap<Integer, Integer> down = new TreeMap<>();
        for (int[] trip : trips) {
            up.put(trip[1], up.getOrDefault(trip[1], 0) + trip[0]);
            down.put(trip[2], down.getOrDefault(trip[2], 0) + trip[0]);
        }
        Iterator<Map.Entry<Integer, Integer>> uper = up.entrySet().iterator();
        Iterator<Map.Entry<Integer, Integer>> downer = down.entrySet().iterator();
        Map.Entry<Integer, Integer> nextDown = downer.next();
        while (uper.hasNext()) {
            Map.Entry<Integer, Integer> curUp = uper.next();
            capacity -= curUp.getValue();
            while (nextDown.getKey() <= curUp.getKey()) {   // means some passengers will get down before cur trip passengers get up
                capacity += nextDown.getValue();
                downer.remove();
                nextDown = downer.next();
            }
            uper.remove();
            if (capacity < 0) return false;
        }
        return true;
    }
}