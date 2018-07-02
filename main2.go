package main

import (
	"fmt"
)

func main() {
	fmt.Println(^uint(0))
	fmt.Println(^uint(0) >> 1)
	fmt.Println(int(^uint(0) >> 1))
	fmt.Println(^int(^uint(0) >> 1))
	x := "ok"
	fmt.Println(len(&x))
	// fd,_ := os.Open("/tmp")
	// defer fd.Close()
	// len,err := fd.Seek(0,2)
}