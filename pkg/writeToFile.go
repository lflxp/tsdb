package pkg

import (
	"fmt"
	"os"
	"unsafe"
)

type Xp struct {
	B  uint32
	C  uint32
	D  uint64
	G  uint32
	AG uint32
	M  []byte
}

func ReadInBuffer(data []byte, id uint64) *Xp {
	return (*Xp)(unsafe.Pointer(&data[id*uint64(os.Getpagesize())]))
}

func PrepareBufData() []byte {
	pageSize := os.Getpagesize()
	buf := make([]byte, pageSize*2)
	for i := 0; i < 2; i++ {
		tmp := ReadInBuffer(buf[:], uint64(i))
		// tmp.A = fmt.Sprintf("%d%s", i, "hello")
		tmp.G = magic
		tmp.C = 0x00000046
		tmp.D = uint64(999 + i)
		tmp.M = []byte("hello")
	}
	return buf
}

func SaveToFile(db *os.File, data []byte) error {
	if info, err := db.Stat(); err != nil {
		return err
	} else if info.Size() == 0 {
		if _, err := db.WriteAt(data, 0); err != nil {
			return err
		}
		// db.Sync()
		// if err := syscall.Fdatasync(int(db.Fd())); err != nil {
		// 	return err
		// }
	}
	return nil
}

func CreateFile(path string) (*os.File, error) {
	var db *os.File
	var err error
	if db, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777); err != nil {
		_ = db.Close()
		return nil, err
	}
	// defer db.Close()

	return db, nil
}

func mains() {
	data := PrepareBufData()
	db, err := CreateFile("this.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	err = SaveToFile(db, data)
	if err != nil {
		fmt.Println(err)
	}
}
