package flexmem_test

import (
	"fmt"
	"testing"

	"github.com/ariefdarmawan/flexmem"
	"github.com/smartystreets/goconvey/convey"
)

var (
	host = "localhost:7777"
)

func TestServer(t *testing.T) {
	convey.Convey("prepare server", t, func() {
		fms := new(flexmem.Server)
		err := fms.Start(host)
		convey.So(err, convey.ShouldBeNil)
		defer fms.Stop()

		convey.Convey("client", func() {
			fmc, err := flexmem.NewClient(host)
			convey.So(err, convey.ShouldBeNil)

			convey.Convey("call", func() {
				r := fmc.Call("kvdb.status", nil)
				convey.So(r.Err(), convey.ShouldBeNil)

				convey.Convey("validate", func() {
					data := string(r.Data)
					convey.So(data, convey.ShouldContainSubstring, "It has been run")
					fmt.Println("Validate result:", data)
				})
			})
		})
	})
}
