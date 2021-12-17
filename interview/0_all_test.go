package main

import (
	"testing"
)

// go test -run TestC
func TestA(t *testing.T) {
	//方法或函数调用传参数, 数组用值传递, map、slice、channel、指针类型这些特殊类型是引用传递。
	A()
}

func TestB(t *testing.T) {
    //字符串相关
	B()
}

func TestC(t *testing.T) {
	//翻转中文字符串
	C()
}
