package main

import (
	"fmt"
	"unsafe"
)

/**
8bit(位)=1Byte(字节)
1024Byte(字节)=1KB
1024KB=1MB
1024MB=1GB
换算率约等于1000（1024），从大到小顺序为T、GB、MB（兆Zhao）、KB、B再小就是位了。
 */
func main() {
	var x = 100
	fmt.Println(unsafe.Sizeof(x)) // 8

	var y int64 = 1
	fmt.Println(unsafe.Sizeof(y)) // 8
	var y1 int32 = 1
	fmt.Println(unsafe.Sizeof(y1)) // 4

	var z uint64 = 1
	fmt.Println(unsafe.Sizeof(z)) // 8
	var z1 uint32 = 1
	fmt.Println(unsafe.Sizeof(z1)) // 4
}
