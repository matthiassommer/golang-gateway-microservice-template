package utils

import "log"

// closable is an interface for all closable streams and resources.
type closable interface {
	Close() error
}

// Close a closable instance, check for errors and print them to console.
func Close(c closable) {
	if c == nil {
		return
	}
	if err := c.Close(); err != nil {
		log.Print(err.Error())
	}
}
