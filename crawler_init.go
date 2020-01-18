package proxypool

func add(name string) {
	c := NewCrawler("files/" + name + ".json")
	NameFuncs[name] = c.Crawl
}

func init() {
	add("xici")
	add("kuaidaili")
	add("89ip")
}
