package main

import (
	"flag"
	"github.com/crawlerclub/proxypool"
)

func main() {
	flag.Parse()
	proxypool.CrawlProxy()
}
