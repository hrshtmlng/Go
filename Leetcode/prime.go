func numPrimeArrangements(n int) int {
	const mod int64 = 1e9 + 7
	count := 0

	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}

	result := (permutation(count, mod) * permutation(n-count, mod)) % mod
	return int(result)
}

func permutation(count int, mod int64) int64 {
	var ans int64 = 1
	for i := 2; i <= count; i++ {
		ans = (ans * int64(i)) % mod
	}
	return ans
}