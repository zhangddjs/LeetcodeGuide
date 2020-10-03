//Time 5%
//Space 98%

class Solution {
    public int[] advantageCount(int[] A, int[] B) {
        //全排列、枚举 n！
        //贪心, 取>B[i]的最小的A中的元素，BST
        List<Integer> list = new ArrayList<>();
        for (int a : A) list.add(a);
        Collections.sort(list);
        for (int i = 0; i < B.length; ++i) {
            Integer val = getAdvance(list, B[i]);
            list.remove(val);
            A[i] = val;
        }
        return A;
    }
    
    public Integer getAdvance(List<Integer> list, int v) {
        int low = 0, high = list.size() - 1;
        if (list.get(high) <= v) return list.get(low);
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (list.get(mid) <= v) low = mid + 1;
            else high = mid;
        }
        return list.get(high);
    }
}