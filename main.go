package main

import (
	"fmt"

	"github.com/lflxp/tsdb/pkg"
)

const metaPageFlag = 0x04

type C struct {
	A  string
	B  string
	CC string
	D  int32
	E  int64
	F  byte
	H  int
}

func main() {
	// pkg.Test()
	// fmt.Println(os.Getpagesize())
	// fmt.Println(metaPageFlag)
	// var w *C = new(C)
	// w.A = "heelo"
	// // fmt.Printf("%d", unsafe.Sizeof(*w))

	// fmt.Println(unsafe.Alignof(w.A), unsafe.Alignof(w.D), unsafe.Alignof(w.E), unsafe.Alignof(w.F), unsafe.Alignof(w.H))

	// a := []byte("123")
	// fmt.Println(a[:])
	_, err := pkg.Open("test.db", 0777, nil)
	if err != nil {
		fmt.Println(err)
	}
}
