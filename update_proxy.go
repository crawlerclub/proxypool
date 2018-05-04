package proxypool

import (
	"github.com/golang/glog"
)

var NameFuncs = make(map[string]func() []string)

func CrawlProxy() {
	glog.Info("start crawling")
	for name, f := range NameFuncs {
		glog.Infof("run %s", name)
		ret := f()
		glog.Info(ret)
	}
}
