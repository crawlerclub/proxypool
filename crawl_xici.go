package proxypool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"crawler.club/et"
)

var onceXici sync.Once
var xici *et.Parser

func Xici() *et.Parser {
	onceXici.Do(func() {
		b, err := ioutil.ReadFile("files/xici.json")
		if err != nil {
			panic(err)
		}
		xici = new(et.Parser)
		if err = json.Unmarshal(b, xici); err != nil {
			panic(err)
		}
	})
	return xici
}

func crawlXici() []string {
	_, items, err := Xici().Do()
	if err != nil || len(items) == 0 {
		return nil
	}
	urls := []string{}
	for _, item := range items {
		vs := item["item"].([]interface{})
		for _, v := range vs {
			m := v.(map[string]interface{})
			if len(m) == 0 {
				continue
			}
			urls = append(urls, fmt.Sprintf("%s:%s", m["ip"], m["port"]))
		}
	}
	return urls
}
