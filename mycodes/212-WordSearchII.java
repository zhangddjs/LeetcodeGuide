class Solution {
    public List<String> findWords(char[][] board, String[] words) {
        List<String> res = new ArrayList<>();
        for (String word : words) {
            boolean flag = false;
            for (int i = 0; i < board.length && !flag; ++i) {
                for (int j = 0; j < board[0].length && !flag; ++j) {
                    if (board[i][j] == word.charAt(0) &&
                        findWord(board, word, i, j, new boolean[board.length * board[0].length])) {
                        res.add(word);
                        flag = true;
                    }
                }
            }
        }
        return res;
    }

    public boolean findWord(char[][] board, String word, int i, int j, boolean[] visited) {
        if (word == null || word.equals("")) return true;
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length ||
            visited[i * board[0].length + j] || board[i][j] != word.charAt(0))
            return false;
        visited[i * board[0].length + j] = true;
        return findWord(board, word.substring(1), i + 1, j, (boolean[]) visited.clone()) ||
               findWord(board, word.substring(1), i - 1, j, (boolean[]) visited.clone()) ||
               findWord(board, word.substring(1), i, j + 1, (boolean[]) visited.clone()) ||
               findWord(board, word.substring(1), i, j - 1, (boolean[]) visited.clone());
    }
}

class Solution2 {

    class TrieNode {
        Map<Character, TrieNode> map;
        boolean isEndOfWord;

        public TrieNode() {
            map = new HashMap<>();
            isEndOfWord = false;
        }

        public boolean isLeaf() {
            return map.isEmpty();
        }

        public TrieNode get(char c) {
            return map.get(c);
        }

        public void put(char c) {
            map.put(c, new TrieNode());
        }

        public void delete(char c) {
            map.remove(c);
        }
    }

    public List<String> findWords(char[][] board, String[] words) {
        List<String> res = new ArrayList<>();
        TrieNode trie = initTrie(words);
        int m = board.length, n = board[0].length;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                find(board, i, j, new boolean[m * n], trie, res, "");
            }
        }
        return res;
    }

    public TrieNode initTrie(String[] words) {
        TrieNode root = new TrieNode();
        for (String word : words) {
            TrieNode node = root;
            for (char c : word.toCharArray()) {
                if (node.get(c) == null) node.put(c);
                node = node.get(c);
            }
            node.isEndOfWord = true;
        }
        return root;
    }

    public boolean find(char[][] board, int i, int j, boolean[] visited, TrieNode trieNode, List<String> res, String prefix) {
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length ||
            trieNode.get(board[i][j]) == null || visited[i * board[0].length + j])
            return false;
        visited[i * board[0].length + j] = true;
        TrieNode cur = trieNode.get(board[i][j]);
        prefix = prefix + board[i][j];
        if (cur.isEndOfWord) {
            res.add(prefix);
            cur.isEndOfWord = false;
        }
        if (find(board, i + 1, j, (boolean[]) visited.clone(), cur, res, prefix)) cur.delete(board[i + 1][j]);
        if (find(board, i - 1, j, (boolean[]) visited.clone(), cur, res, prefix)) cur.delete(board[i - 1][j]);
        if (find(board, i, j + 1, (boolean[]) visited.clone(), cur, res, prefix)) cur.delete(board[i][j + 1]);
        if (find(board, i, j - 1, (boolean[]) visited.clone(), cur, res, prefix)) cur.delete(board[i][j - 1]);
        return cur.isLeaf();
    }
}