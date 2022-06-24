class Solution {
    public String[] reorderLogFiles(String[] logs) {
        List<String> digit = new ArrayList<String>();
        List<String> let = new ArrayList<String>();
        dividLogFiles(logs, digit, let);
        sortLetLogs(let);
        let.addAll(digit);
        return let.toArray(new String[0]);
    }

    private void dividLogFiles(String[] logs, List<String> digit, List<String> let) {
        for (String log : logs) {
            if (isDigLog(log)) digit.add(log);
            else let.add(log);
        }
    }

    private boolean isDigLog(String log) {
        return Character.isDigit(log.charAt(log.indexOf(" ") + 1));
    }

    private void sortLetLogs(List<String> letLogs) {
        Collections.sort(letLogs, new Comparator<String>() {
            public int compare(String a, String b) {
                int c = a.substring(a.indexOf(" ")).compareTo(b.substring(b.indexOf(" ")));
                return c == 0 ? a.substring(0, a.indexOf(" ")).compareTo(b.substring(0, b.indexOf(" "))) : c;
            }
        });
    }
}