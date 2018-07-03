package pkg

import (
	"fmt"
	"os"
	"time"
	"unsafe"
)

// The largest step that can be taken when remapping the mmap.
const maxMmapStep = 1 << 30 // 1GB

// The data file format version.
const version = 2

// Represents a marker value to indicate that a file is a Bolt DB.
const magic uint32 = 0xED0CDAED

type DB struct {
	//数据库路径
	path string
	//文件对象
	file     *os.File
	dataref  *[maxMapSize]byte
	datasz   int
	filesz   int // current on disk file size
	pageSize int
	opened   bool
	freelist *freelist
	readOnly bool

	ops struct {
		writeAt func(b []byte, off int64) (n int, err error)
	}
}

func (db *DB) Path() string {
	return db.path
}

func (db *DB) close() error {
	if !db.opened {
		return nil
	}

	db.opened = false
	db.freelist = nil
	db.ops.writeAt = nil

	if db.file != nil {
		if err := db.file.Close(); err != nil {
			return fmt.Errorf("db file close: %s", err)
		}
		db.file = nil
	}
	db.path = ""
	return nil
}

func (db *DB) init() error {
	db.pageSize = os.Getpagesize()
	buf := make([]byte, db.pageSize*4)
	for i := 0; i < 2; i++ {
		p := db.pageInBuffer(buf[:], pgid(i))
		p.id = pgid(i)
		p.flags = metaPageFlag

		m := p.meta()
		m.magic = magic
		m.version = version
		m.pageSize = uint32(db.pageSize)
		m.freelist = 2
		m.root = bucket{root: 3}
		m.pgid = 4
		m.checksum = m.sum64()
	}

	p := db.pageInBuffer(buf[:], pgid(2))
	p.id = pgid(2)
	p.flags = freelistPageFlag
	p.count = 0

	p = db.pageInBuffer(buf[:], pgid(3))
	p.id = pgid(3)
	p.flags = leafPageFlag
	p.count = 0

	//写入buffer数据到db 文件
	if _, err := db.ops.writeAt(buf, 0); err != nil {
		return err
	}
	if err := fdatasync(db); err != nil {
		return err
	}
	return nil
}

func (db *DB) Close() error {
	return db.close()
}

func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	var db = &DB{opened: true}
	// Set default options if no options are provided.
	if options == nil {
		options = DefaultOptions
	}

	flag := os.O_RDWR
	if options.ReadOnly {
		flag = os.O_RDONLY
		db.readOnly = true
	}

	db.path = path
	var err error
	if db.file, err = os.OpenFile(db.path, flag|os.O_CREATE, mode); err != nil {
		_ = db.close()
		return nil, err
	}

	db.ops.writeAt = db.file.WriteAt

	//创建文件
	if info, err := db.file.Stat(); err != nil {
		return nil, err
	} else if info.Size() == 0 {
		if err := db.init(); err != nil {
			return nil, err
		}
	}
	// else {
	// 	//读取文件
	// 	//0x1000 十六进制 等于十进制的4096
	// 	var buf [0x1000]byte
	// 	if _, err := db.file.ReadAt(buf[:], 0); err == nil {
	// 		m := db.pageInBuffer(buf[:], 0).meta()
	// 	}
	// }
	return db, nil
}

func (db *DB) pageInBuffer(b []byte, id pgid) *page {
	return (*page)(unsafe.Pointer(&b[id*pgid(db.pageSize)]))
}

// Options represents the options that can be set when opening a database.
type Options struct {
	// Timeout is the amount of time to wait to obtain a file lock.
	// When set to zero it will wait indefinitely. This option is only
	// available on Darwin and Linux.
	Timeout time.Duration

	// Sets the DB.NoGrowSync flag before memory mapping the file.
	NoGrowSync bool

	// Open database in read-only mode. Uses flock(..., LOCK_SH |LOCK_NB) to
	// grab a shared lock (UNIX).
	ReadOnly bool

	// Sets the DB.MmapFlags flag before memory mapping the file.
	MmapFlags int

	// InitialMmapSize is the initial mmap size of the database
	// in bytes. Read transactions won't block write transaction
	// if the InitialMmapSize is large enough to hold database mmap
	// size. (See DB.Begin for more information)
	//
	// If <=0, the initial map size is 0.
	// If initialMmapSize is smaller than the previous database size,
	// it takes no effect.
	InitialMmapSize int
}

// DefaultOptions represent the options used if nil options are passed into Open().
// No timeout is used which will cause Bolt to wait indefinitely for a lock.
var DefaultOptions = &Options{
	Timeout:    0,
	NoGrowSync: false,
}
