/**
use a buffer `buf`,  `buf[i][j]` means the max connection nums of subarray `nums1[i:]` with `nums2[j:]`, build the buffer from bottom to top, and return `buf[0][0]`, means max connection of `nums1` and `nums2`.

for e.g
num1 = [1,2,2,4]
num2 = [1,2,4,2]

Step1: init `buf`

||1|2|4|2| num2 |
|---| ----------- | ----------- | ---| ---| ---|
|1||||1|sub num1:[1,2,2,4] --- sub num2:[2]|
|2||||1|[2,2,4] --- [2] also 1 cuz `2` already connected|
|2||||1|[2,4] --- [2] max connection is 1, connect `2`s|
|4|1|1|1|0|means [4] and [2]'s max connection is 0|
|num1|[4] --- [1,2,4,2]|[4] --- [2,4,2]|[4] --- [4,2]|[4] --- [2]|---|

Step2: compute colume by colume
||1|2|4|2| num2 |
|---| ----------- | ----------- | ---| ---| ---|
|1|||1|⬅️1|[1,2,2,4] --- [4,2] 1 != 4 then Max([2,2,4]---[4,2], [1,2,2,4]---[2]|
|2|||⬆️1|⬅️1|[2,2,4] --- [4,2] 2 != 4 then Max([2,4]---[4,2], [2,2,4]---[2])|
|2|||⬆️1|⬅️1|[2,4] --- [4,2] cuz 2 != 4 then Max([4]---[4,2], [2,4]---[2])|
|4|1|1|⬆️1|0|2 != 4 means cannot connect leftmost item of sub num1 and sub num2|
|num1|---|---|---|---|try to connect leftmost item of sub num1 and sub num2|

||1|2|4|2| num2 |
|---| ----------- | ----------- | ---| ---| ---|
|1||2|⬅️1|1|[1,2,2,4] --- [2,4,2] 1 != 2 then Max([2,2,4]---[2,4,2], [1,2,2,4]---[4,2])|
|2||⬆️2|⬅️1|1|[2,2,4] --- [2,4,2] 2 == 2 but not `ok` then Max([2,4]---[2,4,2], [2,2,4]---[4,2])|
|2||⬆️2➡️|↖️1|1|[2,4] --- [2,4,2] cuz 2 == 2 and `ok` then apply val([4]---[4,2]) + 1|
|4|1|1➡️|↖️1|0|val([4]---[2,4,2]) == val([4]---[4,2]) means `ok` for [2,4]---[2,4,2]|
|num1|---|---|---|---|`ok` means the leftmost item (`2`) of sub num2 haven't been connected by any sub num1 item|

||1|2|4|2| num2 |
|---| ----------- | ----------- | ---| ---| ---|
|1|3|2|1|1|1 == 1 and `ok` then val(↘️)+1|
|2|2➡️|↖️⬅️2|1|1|2 != 1 then Max|
|2|⬆️2|⬅️2|1|1|2 != 1 then Max|
|4|⬆️1|1|1|0|---|
|num1|---|---|---|---|---|
 */
class Solution {
    public int maxUncrossedLines(int[] nums1, int[] nums2) {
        int M = nums1.length, N = nums2.length;
        int[][] buf = new int[M][N];
        buf[M - 1][N - 1] = nums1[M - 1] == nums2[N - 1] ? 1 : 0;
        for (int i = M - 2; i >= 0; i--) {
            buf[i][N - 1] = buf[i + 1][N - 1] == 1 ? 1 : nums1[i] == nums2[N - 1] ? 1 : 0;
        }
        for (int i = N - 2; i >= 0; i--) {
            buf[M - 1][i] = buf[M - 1][i + 1] == 1 ? 1 : nums2[i] == nums1[M - 1] ? 1 : 0;
        }
        for (int i = M - 2; i >= 0; i--) {
            for (int j = N - 2; j >= 0; j--) {
                buf[i][j] = Math.max(buf[i][j + 1], buf[i + 1][j]);
                if (nums1[i] == nums2[j] && buf[i + 1][j + 1] == buf[i + 1][j]) {
                    buf[i][j] = buf[i + 1][j + 1] + 1;
                }
            }
        }
        return buf[0][0];
    }
}