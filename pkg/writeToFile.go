package pkg

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type Xp struct {
	A string
	B byte
	C int32
	D int64
}

func readInBuffer(data []byte, id uint64) *Xp {
	return (*Xp)(unsafe.Pointer(&data[id*uint64(os.Getpagesize())]))
}

func prepareBufData() []byte {
	pageSize := os.Getpagesize()
	buf := make([]byte, pageSize*12)
	for i := 0; i < 12; i++ {
		tmp := readInBuffer(buf[:], uint64(i))
		tmp.A = fmt.Sprintf("%d%s", i, "hello")
		tmp.C = int32(998 + i)
	}
	return buf
}

func CreateFile(path string, data []byte) error {
	var db *os.File
	var err error
	if db, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777); err != nil {
		_ = db.Close()
		return err
	}
	defer db.Close()

	if info, err := db.Stat(); err != nil {
		return err
	} else if info.Size() == 0 {
		if _, err := db.WriteAt(data, 0); err != nil {
			return err
		}

		if err := syscall.Fdatasync(int(db.Fd())); err != nil {
			return err
		}
	}
	return nil
}

func mains() {
	data := prepareBufData()
	err := CreateFile("what.db", data)
	if err != nil {
		fmt.Println(err)
	}
}
