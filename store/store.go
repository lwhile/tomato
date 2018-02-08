package store

import (
	"log"

	"github.com/lwhile/tomato"
)

// Store interface
type Store interface {
	Save(*tomato.Tomato) error
	Read(string) error
	ReadAll() ([]tomato.Tomato, error)
	Delete(string) error
}

// DefaultStore :
var DefaultStore Store

func init() {
	var err error
	DefaultStore, err = NewBoltDBCtrl()
	if err != nil {
		log.Fatal(err)
	}
}
