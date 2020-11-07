package main

import (
	"fmt"
)

func testDefer()  {
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func main() {
	testDefer()
	deferCall()
}