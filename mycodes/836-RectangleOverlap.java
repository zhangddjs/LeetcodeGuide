class Solution {
    class Point{
        int x;
        int y;
    }
    class Rect{
        Point bottomLeft = new Point();
        Point bottomRight = new Point();
        Point topLeft = new Point();
        Point topRight = new Point();
        Rect(int[] rec){
            bottomLeft.x = rec[0];
            bottomLeft.y = rec[1];
            bottomRight.x = rec[2];
            bottomRight.y = rec[1];
            topLeft.x = rec[0];
            topLeft.y = rec[3];
            topRight.x = rec[2];
            topRight.y = rec[3];
        }
    }
    public boolean isRectangleOverlap(int[] rec1, int[] rec2) {
        Rect rect1 = new Rect(rec1);
        Rect rect2 = new Rect(rec2);
        if((rect1.topLeft.x < rect2.bottomRight.x &&
            rect1.topLeft.y > rect2.bottomRight.y) &&
           rect1.topLeft.y <= rect2.topLeft.y &&
           (rect1.topRight.x > rect2.bottomLeft.x &&
            rect1.topRight.y > rect2.bottomLeft.y)) return true;
        if((rect2.topLeft.x < rect1.bottomRight.x &&
            rect2.topLeft.y > rect1.bottomRight.y) &&
           rect2.topLeft.y <= rect1.topLeft.y &&
           (rect2.topRight.x > rect1.bottomLeft.x &&
            rect2.topRight.y > rect1.bottomLeft.y)) return true;
        return false;
    }
}

class Solution2 {
    public boolean isRectangleOverlap(int[] rec1, int[] rec2) {
        if ((rec1[0] >= rec2[2]) ||
            (rec1[1] >= rec2[3]) ||
            (rec2[0] >= rec1[2]) ||
            (rec2[1] >= rec1[3])) return false;
        return true;
    }
}