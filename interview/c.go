package main

import"fmt"

func C() {
	src := "你abc好"
	dst := reverse([]rune(src))
	fmt.Printf("%v\n", string(dst))
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}