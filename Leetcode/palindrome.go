func isPalindrome(x int) bool {
	num := x
	reverse := 0
	if num < 0 {
		return false
	}
	for num != 00 {
		reverse = 10*reverse + num%10
		num /= 10
	}
	return (reverse == x)
}