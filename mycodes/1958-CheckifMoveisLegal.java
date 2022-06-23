class Solution {
    int[][] dests = {
        {1, 0},
        {0, 1},
        {1, -1},
        {-1, -1}
    };
    public boolean checkMove(char[][] board, int rMove, int cMove, char color) {
        boolean flag = false;
        for (int[] dest : dests) {
            flag = check(board, rMove, cMove, color, dest, 1);
            if (flag) {
                break;
            }
            flag = check(board, rMove, cMove, color, dest, -1);
            if (flag) {
                break;
            }
        }
        return flag;
    }
    private boolean check(char[][] board, int rMove, int cMove, char color, int[] dest, int dir) {
        boolean isPreSameColor = false;
        boolean hasDifColor = false;
        rMove += dir * dest[0];
        cMove += dir * dest[1];
        while (0 <= rMove && 0 <= cMove && rMove < board.length && cMove < board[0].length) {
            if (board[rMove][cMove] == '.' || isPreSameColor) {
                break;
            }
            if (board[rMove][cMove] == color) {
                isPreSameColor = true;
            } else {
                hasDifColor = true;
            }
            rMove += dir * dest[0];
            cMove += dir * dest[1];
        }
        return hasDifColor && isPreSameColor;
    }
}