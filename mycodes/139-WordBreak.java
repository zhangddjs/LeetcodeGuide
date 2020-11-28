//90% Time
//34% Space
class Solution {
    
    class TrieNode{
        private Map<Character, TrieNode> children;
        private boolean isEnd;
        
        public TrieNode() {
            children = new HashMap<>();
            isEnd = false;
        }
        
        public TrieNode get(char c) {
            return children.get(c);
        }
        
        public void put(char c, TrieNode node) {
            children.put(c, node);
        }
        
        public void setEnd(boolean isEnd) {
            this.isEnd = isEnd;
        }
        
        public boolean isEnd() {
            return this.isEnd;
        }
    }
    
    class Trie{
        
        private TrieNode root;
        
        public Trie() {
            root = new TrieNode();
        }
        
        public void insert(String word) {
            TrieNode node = root;
            for (char c : word.toCharArray()) {
                if (node.get(c) == null) node.put(c, new TrieNode());
                node = node.get(c);
            }
            node.setEnd(true);
        }
    }
    
    int [] canBuiltByDict;
    
    public boolean wordBreak(String s, List<String> wordDict) {
        Trie trie = new Trie();
        for (String word : wordDict) {
            trie.insert(word);
        }
        canBuiltByDict = new int[s.length() + 1];
        canBuiltByDict[0] = 1;
        return helper(s, trie.root);
    }
    
    public boolean helper(String s, TrieNode root) {
        if (s.length() == 0) return true;
        if (canBuiltByDict[s.length()] != 0)
            return canBuiltByDict[s.length()] == 1 ? true : false;
        TrieNode node = root;
        boolean flag = false;
        for (int i = 0; i < s.length() && !flag; ++i) {
            char cur = s.charAt(i);
            node = node.get(cur);
            if (node == null) {
                break;
            }
            if (node.isEnd()) flag |= helper(s.substring(i + 1), root);
        }
        canBuiltByDict[s.length()] = flag ? 1 : -1;
        return flag;
    }
}