package proxypool

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"crawler.club/dl"
	"github.com/golang/glog"
	"github.com/robertkrimen/otto"
)

var (
	DiggerUrl        = "http://www.site-digger.com/html/articles/20110516/proxieslist.html"
	DiggerPattern    = `<td><script>document.write\(decrypt\("(.*)"\)\);</script></td>`
	ReDigger         = regexp.MustCompile(DiggerPattern)
	DiggerKeyPattern = `var baidu_union_id = "(.+)";`
	ReDiggerKey      = regexp.MustCompile(DiggerKeyPattern)
	vm               = otto.New()
)

func init() {
	b, err := ioutil.ReadFile("files/aes.js")
	if err != nil {
		panic(err)
	}
	aes := string(b)
	/*
		aes, err := box.String("aes.js")
		if err != nil {
			panic(err)
		}
	*/
	if _, err = vm.Run(aes); err != nil {
		panic(err)
	}

	NameFuncs["digger"] = crawlDigger
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
	key := matches[0][1]
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
