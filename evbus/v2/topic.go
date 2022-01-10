package v2

import "sync"

type Topic struct {
	rw        sync.RWMutex
	receivers []Receiver
}
