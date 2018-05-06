package proxypool

import (
	"testing"
)

func TestCrawlDigger(t *testing.T) {
	ret := crawlDigger()
	t.Log(ret)
}
