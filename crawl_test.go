package proxypool

import (
	"testing"
)

func TestCrawlDigger(t *testing.T) {
	ret := crawlDigger()
	t.Log(ret)
}

func TestCrawlXici(t *testing.T) {
	ret := crawlXici()
	t.Log(ret)
}
