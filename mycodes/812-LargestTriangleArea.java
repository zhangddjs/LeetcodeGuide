class Solution {
    public double largestTriangleArea(int[][] points) {
        List<Double> list = new ArrayList<>();
        for(int i = 0; i < points.length - 2; i++)
            for(int j = i; j < points.length - 1; j++)
                for(int k = j; k < points.length; k++)
                    list.add(triangleArea(points[i], points[j], points[k]));
        Collections.sort(list);
        return list.get(list.size() - 1);
    }

    public double triangleArea(int[] point0, int[] point1, int[] point2){
        return Math.abs((point1[1] - point0[1]) * (point2[0] - point0[0]) -
                        (point2[1] - point0[1]) * (point1[0] - point0[0])) / 2.0;
    }
}

class Solution2 {
    public double largestTriangleArea(int[][] points) {
        double max = Double.MIN_VALUE;
        for(int i = 0; i < points.length - 2; i++)
            for(int j = i; j < points.length - 1; j++)
                for(int k = j; k < points.length; k++)
                    max = Math.max(triangleArea(points[i], points[j], points[k]), max);
        return max;
    }

    public double triangleArea(int[] point0, int[] point1, int[] point2){
        return Math.abs((point1[1] - point0[1]) * (point2[0] - point0[0]) -
                        (point2[1] - point0[1]) * (point1[0] - point0[0])) / 2.0;
    }
}