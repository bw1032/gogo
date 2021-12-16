package main

import "fmt"

func A() {
	a := 1
	b := []int32{1, 2} //切片
	fmt.Println(fmt.Sprintf("a类型:%T,值:%+v,地址:%v", a, a, &a))
	fmt.Println(fmt.Sprintf("b类型:%T,值:%+v,地址:%p", b, b, &b))
	pri(a, b, b)
	fmt.Println(fmt.Sprintf("b类型:%T,值:%+v,地址:%p", b, b, &b))
}

func pri(a, b interface{}, c []int32) {
	fmt.Println("-----")
	fmt.Println(fmt.Sprintf("a类型:%T,值:%+v,地址:%v", a, a, &a))
	fmt.Println(fmt.Sprintf("b类型:%T,值:%+v,地址:%v", b, b, &b))
	fmt.Println(fmt.Sprintf("c类型:%T,值:%+v,地址:%p", c, c, &c))
	c[0] = 10
	fmt.Println(fmt.Sprintf("b类型:%T,值:%+v,地址:%v", b, b, &b))
	fmt.Println("-----")
}

/*
	结果输出

	a类型:int,值:1,地址:0xc00000e0c0
	b类型:[]int32,值:[1 2],地址:0xc000004078
	-----
	a类型:int,值:1,地址:0xc000030260
	b类型:[]int32,值:[1 2],地址:0xc000030270
	c类型:[]int32,值:[1 2],地址:0xc0000040d8
	b类型:[]int32,值:[10 2],地址:0xc000030270
	-----
	//b类型:[]int32,值:[10 2],地址:0xc000004078
*/
