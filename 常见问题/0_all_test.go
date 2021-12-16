package main

import (
	"testing"
)

//go test -v a.go 0_all_test.go -test.run TestA
func TestA(t *testing.T) {
	//方法或函数调用时，传入参数都是值复制（跟赋值一致），数组用于函数传参时是值复制
	//除非是map、slice、channel、指针类型这些特殊类型是引用传递。
	A()
}
