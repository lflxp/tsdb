package main

/*
常用换算单位有:

8 bit = 1 Byte ；1024 B = 1 KB （KiloByte） ；1024 KB = 1 MB （MegaByte） ；1024 MB = 1 GB （GigaByte） ；1024 GB = 1 TB （TeraByte） 。

字节 （byte）：8个二进制位为一个字节(B)，最常用的单位。计算机存储单位一般用B，KB，MB，GB，TB，PB，EB，ZB，YB，BB来表示，它们之间的关系是：

1B（Byte字节）=8bit

1KB (Kilobyte 千字节)=1024B，

1MB (Mega byte 兆字节 简称“兆”)=1024KB，

1GB (Giga byte 吉字节 又称“千兆”)=1024MB，

1TB (Tera byte 万亿字节 太字节)=1024GB，其中1024=2^10 ( 2 的10次方)，

1PB（Peta byte 千万亿字节 拍字节）=1024TB，

1EB（Exa byte 百亿亿字节 艾字节）=1024PB，

1ZB (Zetta byte 十万亿亿字节 泽字节)= 1024 EB,

1YB (Yotta byte 一亿亿亿字节 尧字节)= 1024 ZB,

1BB (Bronto byte 一千亿亿亿字节)= 1024 YB

1NB(Nona byte )= 1024BB

1DB(Dogga byte)= 1024NB
*/
import (
	"fmt"
	"unsafe"

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

func maind() {
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

	// // data := pkg.PrepareBufData()
	// db, err := pkg.CreateFile("this.db")
	// // db, err := pkg.CreateFile("/home/lxp/Downloads/license.pdf")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()

	// // err = pkg.SaveToFile(db, data)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }

	// info, _ := db.Stat()
	// fmt.Println(int(db.Fd()), info.Size(), info.Name(), os.Getpagesize())

	// mem, err := syscall.Mmap(int(db.Fd()), 0, int(info.Size()), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // fmt.Println(mem[0:10], &mem[1], &mem[2], len(mem))
	// // copy(mem[20:30], []byte("lIxuEpIng"))
	// // fmt.Println(mem)
	// _, _, e1 := syscall.Syscall(syscall.SYS_MADVISE, uintptr(unsafe.Pointer(&mem[0])), uintptr(len(mem)), syscall.MADV_RANDOM)
	// if e1 != 1 {
	// 	fmt.Println(e1)
	// }
	// // tmp := (*[0x1000]byte)(unsafe.Pointer(&mem[0]))
	// //获取数据 转换成struct
	// tmp := (*pkg.Xp)(unsafe.Pointer(&mem[uint64(1)*uint64(os.Getpagesize())]))
	// fmt.Println(tmp, []byte("hello"), mem[:100])
	// // fmt.Println([]byte("hello"))
	// // copy(mem[25:29], []byte("hello"))
	// // fmt.Println((*pkg.Xp)(unsafe.Pointer(&tmp)))
	// // fmt.Println("mem ", mem[:os.Getpagesize()])

	// //Close

	// err = syscall.Munmap(mem)
	// fmt.Println(err)
	// mem = nil

	ddd := pkg.Xp{}
	fmt.Println(unsafe.Sizeof(ddd.M))
	fmt.Println(unsafe.Sizeof(ddd.B))
	fmt.Println(unsafe.Sizeof(ddd.C))
	fmt.Println(unsafe.Sizeof(ddd.D))
	fmt.Println(unsafe.Sizeof(ddd.G))
	fmt.Println(unsafe.Sizeof(ddd.AG))
	fmt.Println(unsafe.Sizeof(ddd))

	// type Xp2 struct {
	// 	B  []byte
	// 	D  uint64
	// 	AD uint32
	// }

	// sss := Xp2{}
	// fmt.Println(unsafe.Sizeof(sss))

	// fmt.Println(uint64(100))
	// fmt.Println(mem[0:8])

	// copy(mem[0:5], []byte("hello"))
	// fmt.Println(mem[0:10])

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

	// buf := make([]byte, os.Getpagesize()*12)
	// // test, err := db.ReadAt(buf[:], int64(i))
	// test, err := db.Read(buf)
	// if err != nil {
	// 	fmt.Println("error", err)
	// }
	// fmt.Println(test)
	// for i := 0; i < 12; i++ {

	// 	m := pkg.ReadInBuffer(buf[:], uint64(i))
	// 	fmt.Println(m.G, m.B, m.C, m.D, m.M, string(m.M))
	// 	fmt.Println("waitting")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("ok")
	// }
	// err = db.Sync()
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
}
