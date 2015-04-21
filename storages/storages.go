package storages

import (
	"crypto/rand"
	"fmt"
)

// Storage .
type Storage struct {
	Driver   string
	Host     string
	Port     string
	Keyspace string
	Index    string
	Table    string
	ID       string
	Created  string
	Updated  string
}

// Driver interface
type Driver interface {
	Connect() error
	GetAll() ([]map[string]interface{}, error)
}

// UUID function by Russ Cox and changed by Michael Hofmann to version 4
func uuid() string {
	b := make([]byte, 16)
	rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10],
		b[10:])
}
