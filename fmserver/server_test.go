package fmserver_test

import (
	"testing"

	"github.com/ariefdarmawan/flexmem/fmclient"
	"github.com/ariefdarmawan/flexmem/fmserver"
	"github.com/eaciit/toolkit"

	"github.com/smartystreets/goconvey/convey"
)

var (
	host = "localhost:7777"
)

func TestServer(t *testing.T) {
	convey.Convey("prepare server", t, func() {
		fms := new(fmserver.Server).
			SetLogger(toolkit.NewLogEngine(false, false, "", "", ""))
		err := fms.Start(host)
		convey.So(err, convey.ShouldBeNil)
		defer fms.Stop()

		convey.Convey("client", func() {
			fmc, err := fmclient.NewClient(host)
			convey.So(err, convey.ShouldBeNil)

			convey.Convey("call status", func() {
				r1 := fmc.Call("KvDB.status")
				convey.So(r1.Err(), convey.ShouldBeNil)

				convey.Convey("validate", func() {
					data := string(r1.Data)
					convey.So(data, convey.ShouldContainSubstring, "It has been run")
					convey.Println("\nValidate result:", data)
				})
			})

			convey.Convey("call ping", func() {
				r1 := fmc.Call("KvDB.ping", "Arief Darmawan")
				convey.So(r1.Err(), convey.ShouldBeNil)

				convey.Convey("validate", func() {
					data := string(r1.Data)
					convey.So(data, convey.ShouldContainSubstring, "welcome to KvDB server")
					convey.Println("\nValidate result:", data)
				})
			})
		})
	})
}
