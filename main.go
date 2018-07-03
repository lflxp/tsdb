package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lflxp/tsdb/pkg"
)

// const metaPageFlag = 0x04

// type C struct {
// 	A  string
// 	B  string
// 	CC string
// 	D  int32
// 	E  int64
// 	F  byte
// 	H  int
// }

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
	// _, err := pkg.Open("test.db", 0777, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// data := pkg.PrepareBufData()
	db, err := pkg.CreateFile("this.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// err = pkg.SaveToFile(db, data)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// buf := bufio.NewReader(db)

	// b, err := buf.ReadString('\n')
	// if err == io.EOF {
	// 	fmt.Println(err)
	// }

	// fmt.Fprintf(os.Stdout, "%s", b)
	// buf := make([]byte, os.Getpagesize())

	// fd, err := ioutil.ReadAll(db)
	// for i := 0; i < 12; i++ {
	// 	tmp := (*pkg.Xp)(unsafe.Pointer(&fd[uint64(i)*uint64(os.Getpagesize())]))
	// 	fmt.Println(tmp)
	// 	fmt.Println(tmp.C)
	// 	fmt.Printf("%s %T\n", tmp.A, tmp.A)
	// }

	// for {
	// 	n, _ := db.Read(buf)
	// 	// fmt.Printf("%d\n", n)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	tmp := (*pkg.Xp)(unsafe.Pointer(&buf[uint64(os.Getpagesize())]))
	// 	fmt.Println(tmp.A)
	// }

	buf := make([]byte, os.Getpagesize()*12)
	// test, err := db.ReadAt(buf[:], int64(i))
	test, err := db.Read(buf)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(test)
	for i := 0; i < 12; i++ {

		m := pkg.ReadInBuffer(buf[:], uint64(i))
		fmt.Println(m.G, m.B, m.C, m.D, m.M, string(m.M))
		fmt.Println("waitting")
		time.Sleep(1 * time.Second)
		fmt.Println("ok")
	}
	err = db.Sync()
	if err != nil {
		fmt.Println("err", err)
	}
}
