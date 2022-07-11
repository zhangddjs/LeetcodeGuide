class SmallestInfiniteSet {

    TreeSet<Integer> set;
    int curSmallest;

    public SmallestInfiniteSet() {
        set = new TreeSet<>();
        curSmallest = 1;
    }
    
    public int popSmallest() {
        if (!set.isEmpty()) {
            int smallest = set.iterator().next();
            set.remove(smallest);
            return smallest;
        }
        return curSmallest++;
    }
    
    public void addBack(int num) {
        if (num < curSmallest) set.add(num);
    }
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * SmallestInfiniteSet obj = new SmallestInfiniteSet();
 * int param_1 = obj.popSmallest();
 * obj.addBack(num);
 */