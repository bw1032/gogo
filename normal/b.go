package main

import (
	"fmt"
	"unicode/utf8"
	"golang.org/x/exp/utf8string"
)

//string和byte区别 (&go test 基准测试 benchmem): https://www.jianshu.com/p/e45f2a69f0aa


func B() {

	//string的索引操作返回的是byte（或uint8），
	//如想获取字符可使用for range，也可使用unicode/utf8包和golang.org/x/exp/utf8string包的At()方法
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	var c = "a你好"
	fmt.Println(c[0]) //97
	fmt.Println(c[1]) //228
	fmt.Println(string(c[1])) //ä
	fmt.Println(utf8string.NewString(c).At(0)) //97
	fmt.Println(utf8string.NewString(c).At(1)) //20320
	fmt.Println(string(utf8string.NewString(c).At(1))) //你
	fmt.Println(utf8string.NewString(c).Slice(1,2))  //你

	//utf8.RuneCountInString(a) 字符长度  len(a) 字节长度
	var a = "a"
	fmt.Println(utf8.RuneCountInString(a)) //1
	fmt.Println(len(a)) //1
	var b = "你好"
	fmt.Println(utf8.RuneCountInString(b)) //2
	fmt.Println(len(b)) //6
}
