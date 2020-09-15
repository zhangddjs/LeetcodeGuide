class Solution {
    public boolean isValid(String s) {
        Set<Character> set = new HashSet<>(Arrays.asList('{', '[', '('));
        Map<Character, Character> map = new HashMap<>();
        Stack<Character> stack = new Stack<>();
        map.put('}', '{');
        map.put(']', '[');
        map.put(')', '(');
        for(int i = 0; i < s.length(); ++i) {
            if (set.contains(s.charAt(i))) stack.push(s.charAt(i));
            else if (stack.isEmpty() || stack.peek() != map.get(s.charAt(i))) return false;
            else stack.pop();
        }
        return stack.isEmpty();
    }
}