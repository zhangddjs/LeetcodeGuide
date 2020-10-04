/*
// Definition for an Interval.
class Interval {
    public int start;
    public int end;

    public Interval() {}

    public Interval(int _start, int _end) {
        start = _start;
        end = _end;
    }
};
*/

//Time 99%
//Space 85%

class Solution {
    public List<Interval> employeeFreeTime(List<List<Interval>> schedule) {
        //求交集，资源分配
        //两个员工比较，用贪心法，比较结果和下一个员工继续比，直到结束。
        List<Interval> tmp = schedule.get(0);
        List<Interval> res = new ArrayList<>();
        for (int i = 1; i < schedule.size(); ++i) {
            List<Interval> cur = schedule.get(i);
            int j = 0, index = 0;
            for (j = 0; j < cur.size(); ++j) {
                Interval interval = cur.get(j);
                index = merge(interval, tmp, index);
            }
        }
        for (int i = 0; i < tmp.size() - 1; ++i)
            if(tmp.get(i).end != tmp.get(i + 1).start)
                res.add(new Interval(tmp.get(i).end, tmp.get(i + 1).start));
        return res;
    }
    
    public int merge(Interval interval, List<Interval> tmp, int index) {
        while (index < tmp.size() && tmp.get(index).end < interval.start) index++;
        if (index >= tmp.size()) {
            tmp.add(interval);
            return index + 1;
        }
        if (tmp.get(index).start > interval.end) {
            tmp.add(index, interval);
            return index + 1;
        }
        while (index + 1 < tmp.size() && tmp.get(index + 1).start <= interval.end) {
            interval.end = Math.max(interval.end, tmp.get(index + 1).end);
            tmp.remove(index + 1);
        }
        tmp.get(index).end = Math.max(interval.end, tmp.get(index).end);
        tmp.get(index).start = Math.min(interval.start, tmp.get(index).start);
        return index;
    }
}