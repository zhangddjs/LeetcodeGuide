//Time 34%
//Space 72%

class Solution {
    public int minSwapsCouples(int[] row) {
        //找出不匹配的对
        //进行交换，每次交换完成1-2对的匹配
        //不匹配的对完成匹配时，返回计数。
        int count = 0;
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < row.length; i += 2) {
            if ((row[i] % 2 == 0 && row[i] + 1 == row[i + 1]) ||
                (row[i] % 2 != 0 && row[i] - 1 == row[i + 1])) continue;
            map.put(row[i], row[i + 1]);
            map.put(row[i + 1], row[i]);
        }
        for (Iterator<Map.Entry<Integer, Integer>> it = map.entrySet().iterator(); it.hasNext();){
            Map.Entry<Integer, Integer> entry = it.next();
            int dest = entry.getKey();
            int tmp1 = entry.getValue();
            int swap = dest % 2 == 0 ? dest + 1 : dest - 1;
            int tmp2 = map.get(swap);
            if (tmp1 == swap)
                continue;
            map.put(tmp1, tmp2);
            map.put(tmp2, tmp1);
            map.put(dest, swap);
            map.put(swap, dest);
            count++;
        }
        return count;
    }
}