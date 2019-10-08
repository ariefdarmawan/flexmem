package kvdb

import (
	"fmt"
	"time"

	"github.com/eaciit/toolkit"
)

type KvDB struct {
	start time.Time
}

func NewKvDB() *KvDB {
	k := new(KvDB)
	k.start = time.Now()
	return k
}

func (k *KvDB) Status() string {
	return fmt.Sprintf("Sebar In-Memory Server v 1.0\nIt has been run for %v from %v",
		time.Since(k.start), k.start)
}

func (k *KvDB) Ping(name string) string {
	return fmt.Sprintf("Hi %s, welcome to KvDB server", name)
}

var baseChar = "QWERTYUIOOPASDFGHJKLZXCVBBNMmnbvcxzlkjhgfdsapoiuytrewq6754890321@$!&"

type SaveResult struct {
	Key   string
	Error error
}

func (k *KvDB) Save(name string, key string, data interface{}) *SaveResult {
	r := new(SaveResult)

	newData := false
	if key == "" {
		newData = true
		key = toolkit.GenerateRandomString(baseChar, 16)
	}

	if newData {
		r.Key = key
	}
	return r
}
