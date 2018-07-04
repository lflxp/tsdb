package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/lflxp/tsdb/pkg"
)

type D struct {
	C []byte
}

type Xp3 struct {
	A int64
	B int32
	C D
}

type Mock struct {
	addr uintptr
	len  int
	cap  int
}

func main() {
	// tmp := &Xp3{
	// 	A: int64(99),
	// 	B: 100,
	// 	C: D{
	// 		C: []byte("hello world"),
	// 	},
	// }

	// Len := unsafe.Sizeof(*tmp)
	// Bytes := &Mock{
	// 	addr: uintptr(unsafe.Pointer(tmp)),
	// 	cap:  int(Len),
	// 	len:  int(Len),
	// }

	// data := *(*[]byte)(unsafe.Pointer(Bytes))
	// fmt.Println("[]byte is :", data)

	db, err := pkg.CreateFile("test.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	info, _ := db.Stat()
	//save and write
	// fmt.Println("before", info.Size())
	// err = pkg.SaveToFile(db, data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("after", info.Size())

	//mmap
	mem, err := syscall.Mmap(int(db.Fd()), 0, int(info.Size()), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mem)

	//byte to struct

	sss := *(**Xp3)(unsafe.Pointer(&mem))
	fmt.Println("byte to struct ", sss)
}
