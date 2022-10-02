class Solution {
    public int minimizeXor(int num1, int num2) {
        int cnt1 = bitSetCnt(num1), cnt2 = bitSetCnt(num2);
        if (cnt1 == cnt2) return num1;
        else if (cnt1 > cnt2) return clearTailBits(num1, cnt1 - cnt2);
        else return FillTailBits(num1, cnt2 - cnt1);
    }

    private int bitSetCnt(int num) {
        int cnt = 0;
        for (int i = 0; i < 32; i++) cnt += (((num >> i) & 1) == 1) ? 1 : 0;
        return cnt;
    }

    private int clearTailBits(int num, int cnt) {
        for (int i = 0; i < 32 && cnt > 0; i++) {
            if (((num >> i) & 1) == 1) {
                cnt--;
                num &= ~(1 << i);
            }
        }
        return num;
    }

    private int FillTailBits(int num, int cnt) {
        for (int i = 0; i < 32 && cnt > 0; i++) {
            if (((num >> i) & 1) != 1) {
                cnt--;
                num |= 1 << i;
            }
        }
        return num;
    }
}