// the num of x and y should be even, then can be make equal
// two situations here:
// 1. s1 `....x...y...` s2 `....y...x...`  -> two swap option to make equal
// 2. s1 `....x...x...` s2 `....y...y...`  -> one swap option to make equal
// then we can get the min opration num by counting the two dif situations num
func minimumSwap(s1 string, s2 string) int {
    xdifnum := 0
    ydifnum := 0
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            if s1[i] == 'x' {
                xdifnum++;
            } else {
                ydifnum++;
            }
        }
    }
    totaldif := xdifnum + ydifnum
    if totaldif % 2 != 0 {      //each situation need a pair of dif to fix, so here couldn't make equal
        return -1
    } else if xdifnum % 2 != 0 {    //then ydifnum will also be odd. there should be one situation 1
        return totaldif / 2 + 1
    }
    return totaldif / 2     // all can be fixed by situation 2
}