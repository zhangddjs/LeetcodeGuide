//9% Time
//5% Space
//Using a stack
class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) return false;
        int temp = x, div = 10;
        Stack<Integer> stack = new Stack<>();
        while (temp != 0) {
            stack.push(temp % div);
            temp /= div;
        }
        while (!stack.isEmpty() && x != 0) {
            if (stack.pop() != x % div) return false;
            x /= div;
        }
        return stack.isEmpty() && x == 0;
    }
}

//reverse the integer
//100% Time
//69% Space
class Solution2 {
    public boolean isPalindrome(int x) {
        if (x < 0) return false;
        long rev = reverse(x);
        return rev - x == 0;
    }

    private long reverse(int x) {
        int div = 10;
        long res = x % div;
        x /= div;
        while (x != 0) {
            res = res * div + x % div;
            x /= div;
        }
        return res;
    }
}
