package main

import (
	"testing"
)

//go test -v 显示详细信息
//go test -v a.go 0_all_test.go -test.run TestA 执行具体方法
//go test -run TestA  执行具体方法  执行具体方法 加载所有.go 文件

func TestA(t *testing.T) {
	//方法或函数调用传参数, 数组用值传递, map、slice、channel、指针类型这些特殊类型是引用传递。
	A()
}

func TestB(t *testing.T) {
    //字符串相关
	B()
}
