package store

import (
	"github.com/lwhile/tomato"
)

// Store interface
type Store interface {
	Save(*tomato.Tomato) error
	Read() error
}
