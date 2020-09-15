class Solution {
    public boolean lemonadeChange(int[] bills) {
        if(bills.length != 0 && bills[0] != 5) return false;
        Map<Integer, Integer> map = new HashMap<>();
        for (int bill : bills) {
            map.put(bill, map.getOrDefault(bill, 0) + 1);
            int change = bill - 5;
            while (bill - 5 != 0) {
                while (change != 0 && (change >= bill || map.getOrDefault(change, 0) == 0)) change -= 5;
                if (change == 0) return false;
                bill -= change;
                map.put(change, map.get(change) - 1);
            }
        }
        return true;
    }
}

class Solution2 {
    public boolean lemonadeChange(int[] bills) {
        if(bills.length != 0 && bills[0] != 5) return false;
        int [] changes = new int[2];
        for (int bill : bills) {
            if(bill != 20) changes[bill / 5 - 1]++;
            switch(bill){
            case 10:
                if (changes[0]-- == 0) return false;
                break;
            case 20:
                if (changes[1] != 0) {
                    if(changes[0]-- == 0) return false;
                    changes[1]--;
                } else if (changes[0] < 3) return false;
                else changes[0] -= 3;
                break;
            }
        }
        return true;
    }
}