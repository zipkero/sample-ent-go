package util

import "log"

type Closer interface {
	Close() error
}

func Close(closer Closer) {
	err := closer.Close()
	if err != nil {
		log.Fatalf("failed closing connection: %v", err)
	}
}
