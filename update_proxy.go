package proxypool

import (
	"github.com/golang/glog"
	"github.com/liuzl/dl"
	"strings"
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

func Validate(addr string) bool {
	req := &dl.HttpRequest{
		Url:      "https://www.baidu.com/",
		UseProxy: true,
		Proxy:    "http://" + addr,
	}
	resp := dl.Download(req)
	if resp.Error != nil {
		glog.Error(resp.Error)
		return false
	}
	return strings.Contains(resp.Text, "<title>百度一下，你就知道</title>")
}
