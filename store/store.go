package store

// Store interface
type Store interface {
	Save() error
	Read() error
}
