package storage_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/eaciit/toolkit"

	"github.com/ariefdarmawan/flexmem/storage"

	"github.com/smartystreets/goconvey/convey"
)

type storageRecord struct {
	Name   string
	Join   time.Time
	Level  int
	Salary float64
}

func newRecord(name string) *storageRecord {
	r := new(storageRecord)
	r.Name = name
	r.Join = time.Now()
	r.Level = toolkit.RandInt(9) + 1
	r.Salary = toolkit.RandFloat(9999, 2) + 1
	return r
}

func TestStorage(t *testing.T) {
	convey.Convey("save data", t, func() {
		dummies := make([]storageRecord, 1000)

		str := storage.NewStorage("test")
		for i := 1; i <= 1000; i++ {
			r := newRecord(fmt.Sprintf("Emp %d", i))
			str.Save(fmt.Sprintf("%d", i), r)
			dummies[i-1] = *r
		}
		convey.So(str.Length(), convey.ShouldEqual, 1000)

		convey.Convey("get one data", func() {
			d := new(storageRecord)
			found, err := str.Get("80", d)
			convey.So(found, convey.ShouldEqual, true)
			convey.So(err, convey.ShouldBeNil)
			convey.Convey("validate get", func() {
				dummy := dummies[79]
				convey.So(*d, convey.ShouldResemble, dummy)
			})
		})

		convey.Convey("delete data", func() {
			d := new(storageRecord)
			str.Delete("100", "101", "102")
			convey.So(str.Length(), convey.ShouldEqual, 997)
			convey.Convey("validate delete", func() {
				found, _ := str.Get("101", d)
				convey.So(found, convey.ShouldEqual, false)
			})
		})

		convey.Convey("slice", func() {
			ds := []storageRecord{}
			err := str.Slice(&ds)
			convey.So(err, convey.ShouldBeNil)
			convey.So(len(ds), convey.ShouldEqual, str.Length())

			convey.Convey("slice key", func() {
				dks := []storageRecord{}
				err := str.SliceByKey(func(k string) bool {
					i, _ := strconv.Atoi(k)
					return i >= 200 && i <= 206
				}, &dks)
				convey.So(err, convey.ShouldBeNil)
				convey.So(len(dks), convey.ShouldEqual, 7)
			})

			convey.Convey("slice value", func() {
				dks := []storageRecord{}
				err := str.SliceByValue(func(d interface{}) bool {
					rs := d.(*storageRecord)
					i, _ := strconv.Atoi(strings.Split(rs.Name, " ")[1])
					return i >= 200 && i <= 206
				}, &dks)
				convey.So(err, convey.ShouldBeNil)
				convey.So(len(dks), convey.ShouldEqual, 7)
			})
		})

	})
}
