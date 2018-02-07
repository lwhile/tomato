package store

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/lwhile/tomato"
)

const (
	boltDBFile = "/var/lib/tomato/tomato.db"
)

type boltDBCtrl struct {
	driver     *bolt.DB
	bucket     *bolt.Bucket
	bucketName string
}

// NewBoltDBCtrl return a boltdb controller
func NewBoltDBCtrl() Store {
	return &boltDBCtrl{
		bucketName: "tomato",
	}
}

// InitDBEnv init the db environment
func (b *boltDBCtrl) InitDBEnv() error {
	var err error
	// prepare the db driver
	if b.driver, err = bolt.Open(boltDBFile, 0600, &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return err
	}

	// prepare bucket
	fn := func(tx *bolt.Tx) error {
		b.bucket, err = tx.CreateBucketIfNotExists([]byte(b.bucketName))
		return err
	}
	return b.driver.Update(fn)
}

func (b *boltDBCtrl) Save(t *tomato.Tomato) error {
	return nil
}

func (b *boltDBCtrl) Read() error {
	return nil
}
