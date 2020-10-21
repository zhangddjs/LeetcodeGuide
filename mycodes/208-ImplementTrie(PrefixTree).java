class Trie {

    public Map<Character, Trie> children;
    public boolean endOfWord;

    /** Initialize your data structure here. */
    public Trie() {
        children = new HashMap<>();
        endOfWord = false;
    }

    /** Inserts a word into the trie. */
    public void insert(String word) {
        Trie cur = this;
        for (char c : word.toCharArray()) {
            cur.children.put(c, cur.children.getOrDefault(c, new Trie()));
            cur = cur.children.get(c);
        }
        cur.endOfWord = true;
    }

    /** Returns if the word is in the trie. */
    public boolean search(String word) {
        Trie cur = this;
        for (char c : word.toCharArray()) {
            cur = cur.children.get(c);
            if (cur == null) return false;
        }
        return cur.endOfWord;
    }

    /** Returns if there is any word in the trie that starts with the given prefix. */
    public boolean startsWith(String prefix) {
        Trie cur = this;
        for (char c : prefix.toCharArray()) {
            cur = cur.children.get(c);
            if (cur == null) return false;
        }
        return true;
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */