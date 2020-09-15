class Solution {
    public int evaluate(String expression) {
        String [] elements = expression.split(" ");
        Stack<Map<String, String>> stack = new Stack<>();
        Stack<String> keys = new Stack<>();
        Stack<String> optkeys = new Stack<>();
        Map<String, String> curMap = new HashMap<>();
        int res = 0;
        for (int i = 0; i < elements.length; ++i) {
            String element = elements[i];
            if (element.startsWith("(")) {
                if (element.equals("(let")) {
                    stack.push(curMap);
                    curMap = new HashMap<String, String>(curMap);
                }
                optkeys.push(element);
                keys.push(element);
            } else if (element.endsWith(")")) {
                StringBuilder valuebuf = new StringBuilder();
                int j;
                for (j = 0; element.charAt(j) != ')'; ++j) valuebuf.append(element.charAt(j));
                String value = curMap.getOrDefault(valuebuf.toString(), valuebuf.toString());
                for (j = j; j < element.length(); ++j) {
                    String opt = optkeys.pop();
                    String subres = keys.pop();
                    subres = curMap.getOrDefault(subres, subres);
                    if (opt.equals("(add")){
                        subres = Integer.valueOf(subres) + Integer.valueOf(value) + "";
                        keys.pop();
                    }
                    else if (opt.equals("(mult")){
                        subres = Integer.valueOf(subres) * Integer.valueOf(value) + "";
                        keys.pop();
                    }
                    else {
                        subres = value.toString();
                        curMap = stack.pop();
                    }
                    value = subres;
                }
                if(keys.isEmpty()) {
                    keys.push(value);
                    break;
                }
                String key = keys.peek();
                if (key.equals(optkeys.peek())) keys.push(value);
                else {
                    key = keys.pop();
                    curMap.put(key, curMap.getOrDefault(value, value));
                }
            } else {
                String key = keys.peek();
                if (key.equals(optkeys.peek())) keys.push(element);
                else {
                    key = keys.pop();
                    curMap.put(key, curMap.getOrDefault(element, element));
                }
            }
        }
        return Integer.valueOf(keys.pop());
    }
}