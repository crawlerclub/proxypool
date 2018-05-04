package proxypool

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/liuzl/dl"
	"regexp"
)

func init() {
	NameFuncs["cz88"] = crawlCZ88
}

var (
	Cz88Url     = "http://www.cz88.net/proxy/index.shtml"
	Cz88Pattern = `<div class="ip">.*?([\d\.]*?)</div><div class="port">([\d]*?)</div>`
	ReCz88      = regexp.MustCompile(Cz88Pattern)

	DiggerUrl        = "http://www.site-digger.com/html/articles/20110516/proxieslist.html"
	DiggerPattern    = `<td><script>document.write\(decrypt\("(.*)"\)\);</script></td>`
	ReDigger         = regexp.MustCompile(DiggerPattern)
	DiggerKeyPattern = `var baidu_union_id = "(.+)";`
	ReDiggerKey      = regexp.MustCompile(DiggerKeyPattern)
)

func crawlCZ88() []string {
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

func crawlDigger() []string {
	glog.Infof("get proxies from: %s", DiggerUrl)
	resp := dl.DownloadUrl(DiggerUrl)
	if resp.Error != nil {
		glog.Error(resp.Error)
		return nil
	}
	matches := ReDiggerKey.FindAllStringSubmatch(resp.Text, -1)
	if len(matches) != 1 {
		glog.Error("couldn't find baidu_union_id in page text")
		return nil
	}
	//key := matches[0][1]

	matches = ReDigger.FindAllStringSubmatch(resp.Text, -1)
	var ret []string
	for _, m := range matches {
		glog.Info(m[1])
	}
	return ret
}
