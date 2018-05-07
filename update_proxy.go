package proxypool

import (
	"github.com/golang/glog"
	"github.com/liuzl/dl"
	"github.com/liuzl/goutil"
	"strings"
	"sync"
	"time"
)

var NameFuncs = make(map[string]func() []string)

func Run(exitCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-exitCh:
			return
		default:
			CrawlProxy(exitCh, wg)
			glog.Info("done!")
			goutil.Sleep(60*time.Minute, exitCh)
		}
	}
}

func CrawlProxy(exitCh chan bool, wg *sync.WaitGroup) {
	glog.Info("start crawling proxies")
	proxy_set := make(map[string]bool)
	proxy_num := 0
	for name, f := range NameFuncs {
		select {
		case <-exitCh:
			return
		default:
			glog.Infof("run %s", name)
			for _, p := range f() {
				proxy_set[p] = true
				proxy_num++
			}
		}
	}
	glog.Infof("total: %d, deduped: %d", proxy_num, len(proxy_set))
	glog.Info("begin to validate")
	var proxyChan = make(chan string)
	for i := 0; i < 60; i++ {
		go doValidate(exitCh, proxyChan, wg)
	}
	for p, _ := range proxy_set {
		proxyChan <- p
	}
	close(proxyChan)
}

func doValidate(exitCh chan bool, proxyChan chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-exitCh:
			return
		case p, ok := <-proxyChan:
			if !ok {
				return
			}
			if Validate(p) {
				if err := InsertProxyStr(p); err != nil {
					glog.Errorf("InsertProxyStr(%s) error: %+v", p, err)
				}
			}
		}
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
