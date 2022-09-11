// Brute Force Permutation, TLE
class Solution {
    int min = Integer.MAX_VALUE;
    public int twoCitySchedCost(int[][] costs) {
        permutation(costs, 0);
        return min;
    }

    private void permutation(int[][] costs, int start) {
        if (start == costs.length - 1) min = Math.min(min, computeSum(costs));
        for (int i = start; i < costs.length; i++) {
            swap(costs, i, start);
            permutation(costs, start + 1);
            swap(costs, i, start);
        }
    }

    private void swap(int[][] costs, int i, int j) {
        int[] tmp = costs[i];
        costs[i] = costs[j];
        costs[j] = tmp;
    }

    private int computeSum(int[][] costs) {
        int sum = 0;
        for (int i = 0; i < costs.length / 2; i++) sum += costs[i][0];
        for (int i = costs.length / 2; i < costs.length; i++) sum += costs[i][1];
        return sum;
    }
}

class Solution {
    public int twoCitySchedCost(int[][] costs) {
        Arrays.sort(costs, (a, b) -> (costSavingChooseCityA(a) - costSavingChooseCityA(b)));
        int n = costs.length, sum = 0, i, j;
        for (i = 0, j = n - 1; i < n / 2 && j >= n / 2;) {
            if (Math.abs(costSavingChooseCityA(costs[i])) < Math.abs(costSavingChooseCityA(costs[j]))) sum += costs[j--][0];
            else sum += costs[i++][1];
        }
        for (i = i; i < n / 2; i++) sum += costs[i][1];
        for (j = j; j >= n / 2; j--) sum += costs[j][0];
        return sum;
    }

    private int costSavingChooseCityA(int[] cost) {
        return cost[1] - cost[0];
    }
}
