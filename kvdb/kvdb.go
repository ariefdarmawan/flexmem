package kvdb

import (
	"fmt"
	"sync"
	"time"

	"github.com/ariefdarmawan/flexmem/storage"
	"github.com/eaciit/toolkit"
)

type KvDB struct {
	start time.Time

	lock     *sync.RWMutex
	storages map[string]*storage.Storage
}

func NewKvDB() *KvDB {
	k := new(KvDB)
	k.start = time.Now()
	k.lock = new(sync.RWMutex)
	k.storages = map[string]*storage.Storage{}
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

	var (
		found bool
		list  *storage.Storage
	)
	k.lock.RLock()
	list, found = k.storages[name]
	k.lock.RUnlock()

	if !found {
		k.lock.Lock()
		list = storage.NewStorage(name)
		k.storages[name] = list
		k.lock.Unlock()
	}

	if err := list.Save(key, data); err != nil {
		r.Error = err
	}

	return r
}
