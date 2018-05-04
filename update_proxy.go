package proxypool

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/liuzl/dl"
	"reflect"
	"regexp"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func CrawlCZ88() []string {
	url := "http://www.cz88.net/proxy/index.shtml"
	regex := `<div class="ip">.*?([\d\.]*?)</div><div class="port">([\d]*?)</div>`
	re := regexp.MustCompile(regex)
	glog.Infof("get proxies from: %s", url)
	resp := dl.DownloadUrl(url)
	if resp.Error != nil {
		glog.Error(resp.Error)
		return nil
	}
	matches := re.FindAllStringSubmatch(resp.Text, -1)
	var ret []string
	for _, m := range matches {
		ret = append(ret, fmt.Sprintf("%s:%s", m[1], m[2]))
	}
	return ret
}

func CrawlProxy() {
	glog.Info("start crawling")
	funcs := []func() []string{CrawlCZ88}
	for _, f := range funcs {
		ret := f()
		glog.Info(ret)
	}
}
