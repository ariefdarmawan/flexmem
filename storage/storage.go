package storage

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/eaciit/toolkit"
)

type Storage struct {
	lock *sync.RWMutex
	list map[string]interface{}
	keys []string

	Name string
}

func NewStorage(name string) *Storage {
	s := new(Storage)
	s.lock = new(sync.RWMutex)
	s.Name = name
	s.list = make(map[string]interface{})
	return s
}

func (s *Storage) Length() int {
	return len(s.list)
}

func (s *Storage) Save(key string, data interface{}) error {
	var err error

	func() {
		handleRecover(&err)
		s.lock.Lock()
		_, ok := s.list[key]
		s.list[key] = data
		if !ok {
			s.keys = append(s.keys, key)
		}
		s.lock.Unlock()
	}()

	return err
}

func (s *Storage) Get(key string, dest interface{}) (bool, error) {
	var err error
	found := false

	func() {
		handleRecover(&err)

		s.lock.RLock()
		defer s.lock.RUnlock()

		d, ok := s.list[key]
		if !ok {
			err = fmt.Errorf("data %s with key %s not found", s.Name, key)
			return
		}

		//fmt.Printf("data %v\n", d)
		found = true
		//dest = d
		v := reflect.ValueOf(dest)
		if v.Kind() != reflect.Ptr {
			err = fmt.Errorf("destination should be a pointer")
			return
		}
		vd := reflect.Indirect(reflect.ValueOf(d))
		v.Elem().Set(vd)
	}()

	return found, err
}

func (s *Storage) Delete(keys ...string) error {
	var err error

	func() {
		handleRecover(&err)

		var keyDeleted bool
		s.lock.Lock()
		for _, key := range keys {
			delete(s.list, key)
		}
		newkeys := make([]string, len(s.keys))
		newkeyCount := 0
		for i, k := range s.keys {
			keyDeleted = false

			for _, dk := range keys {
				if k == dk {
					if len(keys) > 1 {
						keys = keys[1:]
					} else {
						keys = []string{}
					}
					continue
				}
			}

			if !keyDeleted {
				newkeys[i] = k
				newkeyCount++
			}
		}
		s.keys = newkeys
		s.lock.Unlock()
	}()

	return err
}

func (s *Storage) Keys() []string {
	return s.keys
}

func (s *Storage) Slice(dest interface{}) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("target should be ptr to a slice")
	}
	if v.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("target should be ptr to a slice")
	}

	objs := make([]interface{}, s.Length())
	s.lock.RLock()
	for i, k := range s.keys {
		//fmt.Printf("data %d key %s\n", i, k)
		if data, ok := s.list[k]; ok {
			objs[i] = data
		}
	}
	s.lock.RUnlock()
	err := toolkit.Serde(objs, dest, "")
	return err
}

func (s *Storage) SliceByKey(fn func(string) bool, dest interface{}) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("target should be ptr to a slice")
	}
	if v.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("target should be ptr to a slice")
	}

	var err error
	objs := []interface{}{}
	s.lock.RLock()
	for _, k := range s.keys {
		copied := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic detected %v", r)
				}
			}()
			copied = fn(k)
		}()
		if err != nil {
			return err
		}
		if copied {
			if data, ok := s.list[k]; ok {
				objs = append(objs, data)
			}
		}
	}
	s.lock.RUnlock()
	err = toolkit.Serde(objs, dest, "")
	return err
}

func (s *Storage) SliceByValue(fn func(interface{}) bool, dest interface{}) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("target should be ptr to a slice")
	}
	if v.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("target should be ptr to a slice")
	}

	var err error
	objs := []interface{}{}
	s.lock.RLock()
	for _, k := range s.keys {
		if data, ok := s.list[k]; ok {
			copied := false
			func() {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("panic %v", r)
					}
				}()
				copied = fn(data)
			}()
			if err != nil {
				return err
			}
			if copied {
				objs = append(objs, data)
			}
		}
	}
	s.lock.RUnlock()
	err = toolkit.Serde(objs, dest, "")
	return err
}

func handleRecover(err *error) {
	if rec := recover(); rec != nil {
		*err = fmt.Errorf("panic error detected. %v", rec)
	}
}
