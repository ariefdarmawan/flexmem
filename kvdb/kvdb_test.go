package kvdb_test

import (
	"strings"
	"testing"
	"time"

	//flexmem "github.com/ariefdarmawan/flexmem/fmserver"
	"github.com/ariefdarmawan/flexmem/kvdb"
	"github.com/eaciit/toolkit"

	"github.com/smartystreets/goconvey/convey"
)

var (
	dataCount = 100
)

func TestKvDB(t *testing.T) {
	convey.Convey("insert single data", t, func() {
		keys := []string{}
		table := "test"

		db := kvdb.NewKvDB()

		errs := []string{}
		for i := 0; i < dataCount; i++ {
			sr := db.Save(table, "", newData(1))
			if sr.Error != nil {
				errs = append(errs, sr.Error.Error())
			} else {
				keys = append(keys, sr.Key)
			}
		}
		errTxt := strings.Join(errs, "\n")
		convey.So(errTxt, convey.ShouldEqual, "")

		convey.Convey("get all data", func() {
			//db.Query()
		})
	})
}

type testData struct {
	ID   string
	Name string
	Int  int
	Dec  float64
	Date time.Time
}

func newData(i int) *testData {
	td := new(testData)
	td.ID = toolkit.Sprintf("%d_%s", i, toolkit.RandomString(16))
	td.Name = toolkit.Sprintf("Name %d", i)
	td.Int = toolkit.RandInt(1000)
	td.Dec = toolkit.RandFloat(10000, 2) + 1001.30
	td.Date = time.Date(200, 1, 1, 0, 0, 0, 0, time.Local).Add(
		time.Duration(toolkit.RandInt(365)*24) * time.Hour)
	return td
}
