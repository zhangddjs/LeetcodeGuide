def manacher(s: str) -> list:
    """Manacher 算法，O(n)。

    变换串: t = ^ # s0 # s1 # ... # sn-1 # $
    p[i]  : 以 t[i] 为中心的回文半径，数值上等于该回文在原串中的长度

    公式1: 原串开头下标 start = (i - p[i]) // 2, 对应原串回文串 s[start:start+Pi]
    公式2: 右边界 R = C + P[C]
    """
    t = '^#' + '#'.join(s) + '#$'
    n = len(t)
    p = [0] * n
    c = r = 0                                 # C: 最右回文的中心, R: 其右边界
    for i in range(1, n - 1):                 # 跳过两端哨兵
        if i < r:
            p[i] = min(r - i, p[2 * c - i])   # 取对称点的镜像值, 但不超过 R
        while t[i - p[i] - 1] == t[i + p[i] + 1]:   # 哨兵保证不越界, 无需边界判断
            p[i] += 1
        if i + p[i] > r:
            c, r = i, i + p[i]                # 公式2: R = C + P[C]
    return p

def manacher_ints(nums: list) -> list:
    """Manacher 算法的整数序列版，O(n)。

    变换序列: t = ^ None x0 None x1 ... None xn-1 None $
    分隔符用 None（None == 任何int 恒为 False，对任意整数数据都安全）
    哨兵用两个独立的 object()（互不相等，也不等于任何数据）

    公式1: 原序列开头下标 start = (i - p[i]) // 2
    公式2: 右边界 R = C + P[C]
    """
    HEAD, TAIL = object(), object()
    t = [HEAD, None]
    for x in nums:
        t.append(x)
        t.append(None)
    t.append(TAIL)
    n = len(t)
    p = [0] * n
    c = r = 0                                 # C: 最右回文的中心, R: 其右边界
    for i in range(1, n - 1):                 # 跳过两端哨兵
        if i < r:
            p[i] = min(r - i, p[2 * c - i])   # 取对称点的镜像值, 但不超过 R
        while t[i - p[i] - 1] == t[i + p[i] + 1]:   # 扩展回文半径，哨兵保证不越界, 无需边界判断
            p[i] += 1
        if i + p[i] > r:
            c, r = i, i + p[i]                # 公式2: R = C + P[C]
    return p
