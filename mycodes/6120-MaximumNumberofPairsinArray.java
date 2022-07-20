class Solution {
    public int[] numberOfPairs(int[] nums) {
        int[] cnt = new int[101];
        int[] answer = new int[2];
        for (int num : nums) {
            cnt[num]++;
            if (cnt[num] == 2) {
                cnt[num] = 0;
                answer[0]++;
                answer[1]--;
            } else {
                answer[1]++;
            }
        }
        return answer;
    }
}