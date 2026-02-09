package main

import "fmt"

func main() {
	input := "The quic brown fox jumped over lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("Original: %q\n", input)
	fmt.Printf("Reversed: %q\n", rev)
	fmt.Printf("Double Reversed: %q\n", doubleRev)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
