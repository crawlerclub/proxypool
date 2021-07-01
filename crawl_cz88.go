package proxypool

import (
	"fmt"
	"regexp"

	"crawler.club/dl"
	"github.com/golang/glog"
)

func init() {
	NameFuncs["cz88"] = crawlCz88
}

var (
	Cz88Url     = "http://www.cz88.net/proxy/index.shtml"
	Cz88Pattern = `<div class="ip">.*?([\d\.]*?)</div><div class="port">([\d]*?)</div>`
	ReCz88      = regexp.MustCompile(Cz88Pattern)
)

func crawlCz88() []string {
	glog.Infof("get proxies from: %s", Cz88Url)
	resp := dl.DownloadUrl(Cz88Url)
	if resp.Error != nil {
		glog.Error(resp.Error)
		return nil
	}
	matches := ReCz88.FindAllStringSubmatch(resp.Text, -1)
	var ret []string
	for _, m := range matches {
		ret = append(ret, fmt.Sprintf("%s:%s", m[1], m[2]))
	}
	return ret
}
