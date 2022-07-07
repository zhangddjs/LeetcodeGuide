// TLE
class Solution {
    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        List<List<String>> res = new ArrayList<>(), total = new ArrayList<>();
        if (!endWord.equals(wordList.get(wordList.size() - 1))) return res;
        compute(beginWord, endWord, wordList, total, new ArrayList<>(Arrays.asList(beginWord)));
        int minLen = Integer.MAX_VALUE;
        for (List<String> l : total) minLen = Math.min(l.size(), minLen);
        for (List<String> l : total) if (l.size() == minLen) res.add(l);
        return res;
    }

    private void compute(String beginWord, String endWord, List<String> wordList, List<List<String>> total, List<String> cur) {
        if (beginWord.equals(endWord)) {
            List<String> temp = new ArrayList<>();
            temp.addAll(cur);
            total.add(temp);
            return;
        }
        for (int i = 0; i < wordList.size(); i++) {
            String word = wordList.get(i);
            if (!isNextLadder(beginWord, word)) continue;
            Collections.swap(wordList, i, wordList.size() - 1);
            wordList.remove(wordList.size() - 1);
            cur.add(word);
            compute(word, endWord, wordList, total, cur);
            cur.remove(cur.size() - 1);
            wordList.add(word);
            Collections.swap(wordList, i, wordList.size() - 1);
        }
    }

    private boolean isNextLadder(String a, String b) {
        int diff = 0;
        for (int i = 0; i < a.length(); i++) if (a.charAt(i) != b.charAt(i)) diff++;
        return diff == 1;
    }
}

// TLE
class Solution {
    Map<String, List<List<String>>> routes;
    Set<String> beginLadders;
    int minLen;
    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        List<List<String>> res = new ArrayList<>();
        boolean flag = false;
        //if (!endWord.equals(wordList.get(wordList.size() - 1))) return res;
        for (int i = 0; i < wordList.size(); i++) {
            if (wordList.get(i).equals(endWord)) {
                wordList.remove(i);
                flag = true;
            }
        }
        if (!flag) return res;
        if (isNextLadder(beginWord, endWord)) {
            res.add(Arrays.asList(beginWord, endWord));
            return res;
        }
        routes = new HashMap<>();
        beginLadders = new HashSet<>();
        minLen = Integer.MAX_VALUE;
        routes.put(endWord, new ArrayList<>(Arrays.asList(new ArrayList<>(Arrays.asList(endWord)))));
        compute(beginWord, endWord, wordList, 2);
        for (String w : beginLadders) {
            for (List<String> l : routes.get(w)) {
                if (l.size() != minLen) continue;
                l.add(beginWord);
                Collections.reverse(l);
                res.add(l);
            }
        }
        
        return res;
    }

    private void compute(String beginWord, String endWord, List<String> wordList, int len) {
        if (wordList.size() == 0 || len > minLen) return;
        List<String> nextLayer = new ArrayList<>();
        List<String> ladder = new ArrayList<>();
        for (String word : wordList) {
            if (isNextLadder(word, endWord)) {
                ladder.add(word);
                if(!routes.containsKey(word)) routes.put(word, new ArrayList<>());
                for (List<String> route : routes.get(endWord)) {
                    List<String> l = new ArrayList<>(route);
                    l.add(word);
                    routes.get(word).add(l);
                }
                if (isNextLadder(word, beginWord)) {
                    minLen = Math.min(minLen, len);
                    beginLadders.add(word);
                }
            } else {
                nextLayer.add(word);
            }
        }
        for (String w : ladder) {
            compute(beginWord, w, nextLayer, len + 1);
        }
    }

    private boolean isNextLadder(String a, String b) {
        int diff = 0;
        for (int i = 0; i < a.length(); i++) if (a.charAt(i) != b.charAt(i)) diff++;
        return diff == 1;
    }
}

// TLE
class Solution {
    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        List<List<String>> res = new ArrayList<>();
        if (!remove(wordList, endWord)) return res;
        Map<String, List<List<String>>> map = new HashMap<>();
        Set<String> needDelete = new HashSet<>();
        Queue<String> q = new LinkedList<>();
        map.put(endWord, new ArrayList<>(Arrays.asList(new ArrayList<>(Arrays.asList(endWord)))));
        remove(wordList, beginWord);
        wordList.add(beginWord);
        q.offer(endWord);
        boolean flag = false;
        int size = q.size();
        while (!q.isEmpty()) {
            String s = q.poll();
            for (String word : wordList) {
                if (isLadder(word, s)) {
                    if (!needDelete.contains(word)) q.offer(word);
                    needDelete.add(word);
                    if(!map.containsKey(word)) map.put(word, new ArrayList<>());
                    for (List<String> route : map.get(s)) {
                        List<String> l = new ArrayList<>(route);
                        l.add(word);
                        map.get(word).add(l);
                    }
                    if (word.equals(beginWord)) flag = true;
                }
            }
            size--;
            if (size == 0) {
                if (flag) break;
                int len = wordList.size();
                for (int i = 0; i < len; i++) if (!needDelete.contains(wordList.get(i))) wordList.add(wordList.get(i));
                wordList = wordList.subList(len, wordList.size());
                needDelete = new HashSet<>();
                size = q.size();
            }
        }

        if (!map.containsKey(beginWord)) return res;
        for (List<String> beginLadders : map.get(beginWord)) {
            Collections.reverse(beginLadders);
            res.add(beginLadders);
        }
        return res;
    }

    private boolean remove(List<String> wordList, String endWord) {
        for (int i = 0; i < wordList.size(); i++) {
            if (wordList.get(i).equals(endWord)) {
                Collections.swap(wordList, i, wordList.size() - 1);
                wordList.remove(wordList.size() - 1);
                return true;
            }
        }
        return false;
    }

    private boolean isLadder(String a, String b) {
        int diff = 0;
        for (int i = 0; i < a.length(); i++) if (a.charAt(i) != b.charAt(i)) diff++;
        return diff == 1;
    }
}

class Solution {
    Map<String, List<List<String>>> buf;
    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        if (!remove(wordList, endWord)) return new ArrayList<>();
        remove(wordList, beginWord);
        wordList.add(endWord);
        buf = new HashMap<>();
        Map<String, List<String>> graph = buildLadderGraph(beginWord, endWord, wordList);
        traverseLadderGraph(graph, beginWord, endWord);
        return buf.getOrDefault(beginWord, new ArrayList<>());
    }

    private List<List<String>> traverseLadderGraph(Map<String, List<String>> graph, String beginWord, String endWord) {
        if (beginWord.equals(endWord)) {
            buf.put(beginWord, new ArrayList<>());
            buf.get(beginWord).add(new ArrayList<>(Arrays.asList(endWord)));
            return buf.get(beginWord);
        }
        if (buf.containsKey(beginWord)) return buf.get(beginWord);
        if (!graph.containsKey(beginWord)) return new ArrayList<>();
        List<List<String>> next = new ArrayList<>();
        for (String s : graph.get(beginWord)) {
            next.addAll(traverseLadderGraph(graph, s, endWord));
        }
        buf.put(beginWord, new ArrayList<>());
        for (List<String> route : next) {
            List<String> l = new ArrayList<>(route);
            l.add(0, beginWord);
            buf.get(beginWord).add(l);
        }
        return buf.get(beginWord);
    }

    private Map<String, List<String>> buildLadderGraph(String beginWord, String endWord, List<String> wordList) {
        Map<String, List<String>> map = new HashMap<>();
        Set<String> needDelete = new HashSet<>();
        Queue<String> q = new LinkedList<>();
        q.offer(beginWord);
        boolean flag = false;
        int size = q.size();
        while (!q.isEmpty()) {
            String s = q.poll();
            map.put(s, new ArrayList<>());
            for (String word : wordList) {
                if (isLadder(word, s)) {
                    if (!needDelete.contains(word)) q.offer(word);
                    needDelete.add(word);
                    map.get(s).add(word);
                    if (word.equals(endWord)) flag = true;
                }
            }
            if (--size == 0) {
                if (flag) break;
                int len = wordList.size();
                for (int i = 0; i < len; i++) if (!needDelete.contains(wordList.get(i))) wordList.add(wordList.get(i));
                wordList = wordList.subList(len, wordList.size());
                needDelete = new HashSet<>();
                size = q.size();
            }
        }
        return map;
    }

    private boolean remove(List<String> wordList, String endWord) {
        for (int i = 0; i < wordList.size(); i++) {
            if (wordList.get(i).equals(endWord)) {
                Collections.swap(wordList, i, wordList.size() - 1);
                wordList.remove(wordList.size() - 1);
                return true;
            }
        }
        return false;
    }

    private boolean isLadder(String a, String b) {
        int diff = 0;
        for (int i = 0; i < a.length(); i++) if (a.charAt(i) != b.charAt(i)) diff++;
        return diff == 1;
    }
}