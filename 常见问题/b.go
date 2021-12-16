package main

import (
	"fmt"
	"unicode/utf8"
)

func B() {

	var a = "a"
	fmt.Println(utf8.RuneCountInString(a)) //1
	fmt.Println(len(a)) //1
	var b = "你好"
	fmt.Println(utf8.RuneCountInString(b)) //2
	fmt.Println(len(b)) //6
}
