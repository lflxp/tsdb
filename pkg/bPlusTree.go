package pkg

// https://www.jianshu.com/p/ffeeb3d0efd6
import (
	"log"
	"os"

	"github.com/boltdb/bolt"
)

func Test() {
	db, err := bolt.Open("test.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(db.Path())

	if err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("hello world"))
		if err != nil {
			return err
		}

		if err := b.Put([]byte("foo"), []byte("bra")); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket([]byte("hello world")).Get([]byte("foo"))
		log.Println(string(value))
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
