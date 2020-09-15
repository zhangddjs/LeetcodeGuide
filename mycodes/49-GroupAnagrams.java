class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for (String str : strs) {
            char[] arr = str.toCharArray();
            Arrays.sort(arr);
            String tmp = new String(arr);
            if (map.get(tmp) == null) map.put(tmp, new ArrayList<String>());
            map.get(tmp).add(str);
        }
        return map.values().stream().collect(Collectors.toList());      //return new ArrayList(map.values());
    }
}