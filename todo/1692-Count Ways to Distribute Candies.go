// Recursion -> TLE
// we can follow this fomula:
// waysToDistribute(n int, k int) == C(0, n-1) * waysToDistribute(n - 1, k - 1) + C(1, n-1) * waysToDistribute(n - 2, k - 1) 
// + ... 
// + C(n - k, n-1) * waysToDistribute(k - 1, k - 1) 

// f(n, k) = C(0, n - 1) * f(n - 1, k - 1) + C(1, n - 1) * f(n - 2, k - 1) + ... + C(n - k, n - 1) * f(k - 1, k - 1)
// 

func waysToDistribute(n int, k int) int {
	res := compute(n, k)
	return int(res.Mod(res, big.NewInt(1e9+7)).Int64())
}

func compute(n int, k int) *big.Int {
	if k <= 1 || n == k {
		return big.NewInt(1)
	}
	res := big.NewInt(0)
	for i := 0; i <= n - k; i++ {
		c := C(i, n - 1)
		res.Add(res, c.Mul(c, compute(n - 1 - i, k - 1)))
	}
	return res
}

func F(n int) *big.Int {
	res := big.NewInt(1)
	for i := 1; i <= n; i++ {
		res.Mul(res, big.NewInt(int64(i)))
	}
	return res
}

func C(k, n int) *big.Int {
	if k == 0 {
		return big.NewInt(1)
	}
	Fn := F(n)
	Fk := F(k)
	Fnk := F(n-k)

	res := Fn.Div(Fn, Fk).Div(Fn, Fnk)
	return res
}

// Iteration
var dp [][]*big.Int	//dp[i][j] means waysToDistribute(i, j)

func waysToDistribute(n int, k int) int {
	res := compute(n, k)
	return int(res.Mod(res, big.NewInt(1e9+7)).Int64())
}

func compute(n int, k int) *big.Int {
	if k <= 1 || n == k {
		return big.NewInt(1)
	}
	// init
	dp = make([][]*big.Int, k + 1)
	for i := 1; i <= k; i++ {
		dp[i] = make([]*big.Int, n + 1)
	}
	for i := 1; i <= n; i++ {
		dp[1][i] = big.NewInt(1)		//k == 1 ==> res == 1
	}
	
	// dp
	for i := 2; i <= k; i++ {
		for j := i; j <= n; j++ {
			tmp := big.NewInt(0)
			for h := 0; h <= j - i; h++ {
				c := C(h, j - 1)
				tmp.Add(tmp, c.Mul(c, dp[i - 1][j - 1 - h]))
			}
			dp[i][j] = tmp
		}
	}

	return dp[k][n]
}

func F(n int) *big.Int {
	res := big.NewInt(1)
	for i := 1; i <= n; i++ {
		res.Mul(res, big.NewInt(int64(i)))
	}
	return res
}

func C(k, n int) *big.Int {
	if k == 0 {
		return big.NewInt(1)
	}
	Fn := F(n)
	Fk := F(k)
	Fnk := F(n-k)

	res := Fn.Div(Fn, Fk).Div(Fn, Fnk)
	return res
}

// n, k Permutations


