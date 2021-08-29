// Ologn, not optimal。
func hasAlternatingBits(n int) bool {
    before := -1
    for i := n; i != 0; i >>= 1 {
        cur := i & 1
        if cur == before {
            return false
        }
        before = cur
    }
    return true
}