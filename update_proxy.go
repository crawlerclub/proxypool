package proxypool

import (
	"github.com/golang/glog"
)

var NameFuncs = make(map[string]func() []string)

func CrawlProxy() {
	glog.Info("start crawling")
	funcs := []string{"cz88"}
	for _, f := range funcs {
		if NameFuncs[f] == nil {
			continue
		}
		ret := NameFuncs[f]()
		glog.Info(ret)
	}
}
