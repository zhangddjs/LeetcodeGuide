class Solution {
    public int[] decode(int[] encoded) {
        int xor = 0;
        int[] perm = new int[encoded.length + 1];
        for (int i = 1; i <= encoded.length + 1; i++) xor ^= i;
        for (int i = 1; i < encoded.length; i += 2) xor ^= encoded[i];
        perm[0] = xor;
        for (int i = 1; i < perm.length; i++) perm[i] = perm[i - 1] ^ encoded[i - 1];
        return perm;
    }
}