package netx

import (
	"flag"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetIPByNICName(t *testing.T) {
	flag.Parsed()
	nic := flag.Arg(0)
	if nic == "" {
		nic = "lo"
	}

	Convey("获取网卡ip", t, func() {
		Convey("获取ipv4", func() {
			ip, err := GetIPv4ByNICName(nic)
			So(err, ShouldBeNil)
			println(ip)
		})
		Convey("获取ipv6", func() {
			ip, err := GetIPv6ByNICName(nic)
			So(err, ShouldBeNil)
			println(ip)
		})
		Convey("获取ip（ipv4优先）", func() {
			ip, err := GetIPByNICName(nic)
			So(err, ShouldBeNil)
			println(ip)
		})
	})
}
