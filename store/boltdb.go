package store

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/lwhile/tomato"
)

type boltDBCtrl struct {
	driver     *bolt.DB
	bucketName []byte
}

// NewBoltDBCtrl return a boltdb controller
func NewBoltDBCtrl() (Store, error) {
	dbCtrl := boltDBCtrl{
		bucketName: []byte("tomato"),
	}
	err := dbCtrl.InitDBEnv()
	return &dbCtrl, err
}

// InitDBEnv init the db environment
func (b *boltDBCtrl) InitDBEnv() error {
	prepareDBFile()

	var err error
	// prepare the db driver
	if b.driver, err = bolt.Open(boltDBPath(), 0700, &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return err
	}

	// prepare bucket
	fn := func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(b.bucketName))
		return err
	}
	return b.driver.Update(fn)
}

func homeDir() string {
	// Not support Windows
	return os.Getenv("HOME")
}

func boltDBDir() string {
	return path.Join(homeDir(), "tomato")
}

func boltDBPath() string {
	return path.Join(boltDBDir(), "tomato.db")
}

func prepareDBFile() {
	p := boltDBPath()

	// data file and dir not exist

	if _, err := os.Stat(boltDBPath()); os.IsNotExist(err) {

		// if dir not exist, create it.
		if _, err := os.Stat(boltDBDir()); os.IsNotExist(err) {
			err = os.MkdirAll(p, 0700)
			if err != nil {
				log.Fatal(err)
			}
		}

		fp, err := os.Create(boltDBPath())
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()
	}
}

func (b *boltDBCtrl) Save(t *tomato.Tomato) error {
	fn := func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.bucketName)
		id, err := b.setupID(bucket, t)
		if err != nil {
			return err
		}
		data, err := json.Marshal(t)
		if err != nil {
			return err
		}
		return bucket.Put(id, data)
	}

	return b.driver.Update(fn)
}

func (b *boltDBCtrl) Read() error {
	return nil
}

func (b *boltDBCtrl) setupID(bk *bolt.Bucket, t *tomato.Tomato) ([]byte, error) {
	var err error
	t.ID, err = bk.NextSequence()
	if err != nil {
		return nil, err
	}
	return []byte(strconv.FormatUint(t.ID, 10)), nil
}
