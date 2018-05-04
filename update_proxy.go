package proxypool

import (
	"github.com/golang/glog"
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func CrawlCZ88() []string {
	glog.Info("get proxies from: http://www.cz88.net/proxy/index.shtml")
	return nil
}

func CrawlProxy() {
	glog.Info("start crawling")
	funcs := []func() []string{CrawlCZ88}
	for _, f := range funcs {
		f()
	}
}
