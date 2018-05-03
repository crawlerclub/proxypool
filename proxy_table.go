package proxypool

import (
	"fmt"
)

func ReadProxy() ([]*Proxy, error) {
	db := GetMySQLHandler()
	var records []*Proxy
	ret := db.Where(
		"update_time>last_fail_time OR total_crawl<5 OR total_fail/total_crawl<0.5").
		Find(&records)
	fmt.Printf("%+v\n", ret)
	return records, nil
}

func InsertProxy(proxyList []*Proxy) error {
	// TODO https://github.com/btfak/sqlext
	return nil
}
