//贪心法，优先体重小的，行不通。
//反例
//people = [3,5,3,4,2,5,7,1,6,3,6,8,3,1,5,9,1,4,8,5,1,9], limit=9

class Solution0 {
    public int numRescueBoats(int[] people, int limit) {
        //求最值
        //资源分配
        //贪心？似乎行不通
        Arrays.sort(people);
        int cur = 0, res = 1;
        for (int i = 0; i < people.length; ++i) {
            if (cur + people[i] <= limit) cur += people[i];
            else {
                cur = people[i];
                res++;
            }
        }
        return res;
    }
}

// 大小根堆 --> 双指针
// 但一只船只能载两个人
class Solution0 {
    public int numRescueBoats(int[] people, int limit) {
        Arrays.sort(people);
        int cur = 0, res = 0, i = 0, j = people.length - 1;
        while (i <= j) {
            while (i < j && cur + people[j] <= limit) cur += people[j--];
            while (i <= j && cur + people[i] <= limit) cur += people[i++];
            cur = 0;
            res++;
        }
        return res;
    }
}

// 对双指针法进行改进。
// Time 32%
// Space 70%
class Solution {
    public int numRescueBoats(int[] people, int limit) {
        Arrays.sort(people);
        int cur = 0, res = 0, i = 0, j = people.length - 1;
        while (i <= j) {
            if (cur + people[j] <= limit) cur += people[j--];
            if (cur + people[i] <= limit) cur += people[i++];
            cur = 0;
            res++;
        }
        return res;
    }
}

//贪心最优性证明 https://leetcode.com/problems/boats-to-save-people/discuss/156740/C++JavaPython-Two-Pointers/162296