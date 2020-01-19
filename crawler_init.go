package proxypool

import (
	"io/ioutil"
	"strings"
)

func add(name string) {
	c := NewCrawler("files/" + name + ".json")
	NameFuncs[name] = c.Crawl
}

func init() {
	b, err := ioutil.ReadFile("files/active")
	if err != nil {
		panic(err)
	}
	for _, name := range strings.Fields(string(b)) {
		add(name)
	}
}
