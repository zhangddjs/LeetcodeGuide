class RLEIterator {

    int[] A;
    int j = 0, remain = 0;
    
    public RLEIterator(int[] A) {
        this.A = A;
        remain = A.length == 0 ? 0 : A[j];
    }
    
    public int next(int n) {
        if (A.length == 0) return -1;
        remain -= n;
        while (remain < 0 && j + 2 < A.length) {
            j += 2;
            remain += A[j];
        }
        if (remain < 0) return -1;
        return A[j + 1];
    }
}