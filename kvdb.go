package flexmem

import (
	"fmt"
	"time"
)

type kvdb struct {
	start time.Time
}

func newKvdb() *kvdb {
	k := new(kvdb)
	k.start = time.Now()
	return k
}

func (k *kvdb) Status() string {
	return fmt.Sprintf("Sebar In-Memory Server v 1.0\nIt has been run for %v from %v",
		time.Since(k.start), k.start)
}
