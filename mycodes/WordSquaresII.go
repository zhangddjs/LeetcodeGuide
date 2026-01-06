func wordSquares(words []string) [][]string {
    res := make([][]string, 0)
    for i := 0; i < len(words)-3; i++ {
        for j := i+1; j < len(words)-2; j++ {
            for k := j+1; k < len(words)-1; k++ {
                for h := k+1; h < len(words); h++ {
                    p := make([][]string, 0)
                    square := []string{words[i],words[j],words[k],words[h]}
                    permu(square, 0, &p)
                    for i := range p {
                        if valid(p[i]) {
                            res = append(res, p[i])
                        }
                    }
                }
            }
        }
    }
    sort.Slice(res, func(i, j int) bool {
        for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
            if res[i][k] != res[j][k] {
                return res[i][k] < res[j][k]
            }
        }
        return len(res[i]) < len(res[j])
    })
    return res
}

func permu(words []string, i int, res *[][]string) {
    if i == 4 {
        copied := make([]string, 4)
        for i := range 4 {
            copied[i] = words[i]
        }
        (*res) = append((*res), copied)
    }
    for j := i; j < len(words); j++ {
        words[j], words[i] = words[i], words[j]
        permu(words, i+1, res)
        words[j], words[i] = words[i], words[j]
    }
}

func valid(words []string) bool {
    if len(words) != 4 {
        return false
    }
    return words[0][0] == words[1][0] &&
            words[0][3] == words[2][0] &&
            words[3][0] == words[1][3] &&
            words[3][3] == words[2][3]
}©leetcode
