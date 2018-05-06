package proxypool

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/liuzl/dl"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"regexp"
)

func init() {
	NameFuncs["digger"] = crawlDigger
}

var (
	DiggerUrl        = "http://www.site-digger.com/html/articles/20110516/proxieslist.html"
	DiggerPattern    = `<td><script>document.write\(decrypt\("(.*)"\)\);</script></td>`
	ReDigger         = regexp.MustCompile(DiggerPattern)
	DiggerKeyPattern = `var baidu_union_id = "(.+)";`
	ReDiggerKey      = regexp.MustCompile(DiggerKeyPattern)
)

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
	key := matches[0][1]

	script, err := ioutil.ReadFile("aes.js")
	if err != nil {
		glog.Error(err)
		return nil
	}
	vm := otto.New()
	if _, err = vm.Run(script); err != nil {
		glog.Error(err)
		return nil
	}

	matches = ReDigger.FindAllStringSubmatch(resp.Text, -1)
	var ret []string
	for _, m := range matches {
		result, err := vm.Run(fmt.Sprintf(
			"var baidu_union_id='%s'; decrypt('%s');", key, m[1]))
		if err != nil {
			glog.Error(err)
			break
		}
		ret = append(ret, fmt.Sprintf("%v", result))
	}
	return ret
}
