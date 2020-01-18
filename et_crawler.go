package proxypool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"crawler.club/et"
)

type Crawler struct {
	p *et.Parser
}

func NewCrawler(conf string) *Crawler {
	b, err := ioutil.ReadFile(conf)
	if err != nil {
		panic(err)
	}
	p := new(et.Parser)
	if err = json.Unmarshal(b, p); err != nil {
		panic(err)
	}
	return &Crawler{p: p}
}

func (c *Crawler) Crawl() []string {
	_, items, err := c.p.Do()
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
