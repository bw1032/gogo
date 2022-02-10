package main

import (
	"fmt"
)

func A()  {

	c := make(chan string)
	fmt.Println(c)

	//1 主线程在通道写入之前  先行读取
	//fmt.Println(<-c)
	//go func() {
	//	c <- "1"
	//}()
	//结果: fatal error: all goroutines are asleep - deadlock!


	//2 主线程对无缓冲(或者缓冲已满)的通道  先行写入
    //c <- "1"
    //go func() {
    //	fmt.Println(<-c)
	//}()
	//结果:fatal error: all goroutines are asleep - deadlock!

	//颠倒顺序 可成功
	//go func() {
	//	fmt.Println(<-c)
	//}()
	//c <- "1"

}
