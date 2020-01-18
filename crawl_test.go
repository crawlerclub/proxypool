package proxypool

import (
	"testing"
)

func TestCrawlDigger(t *testing.T) {
	ret := crawlDigger()
	t.Log(ret)
}

func TestCrawlXici(t *testing.T) {
	c := NewCrawler("files/xici.json")
	ret := c.Crawl()
	t.Log(ret)
}

func TestCrawlKuaidaili(t *testing.T) {
	c := NewCrawler("files/kuaidaili.json")
	ret := c.Crawl()
	t.Log(ret)
}
