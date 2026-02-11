func trailingZeroes(n int) int {
	sum := 0

	for n/5 > 0 {
		sum += n / 5
		n /= 5
	}
	return sum
}