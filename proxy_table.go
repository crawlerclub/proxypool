package proxypool

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/liuzl/sqlext"
)

func ReadProxy() ([]*Proxy, error) {
	var records []*Proxy
	ret := GetMySQLHandler().Where(
		"update_time>last_fail_time OR total_crawl<5 OR total_fail/total_crawl<0.5").
		Find(&records)
	fmt.Printf("%+v\n", ret)
	return records, nil
}

func InsertProxy(proxies []*Proxy) error {
	if len(proxies) <= 0 {
		return nil
	}
	db := GetMySQLHandler()
	_, err := sqlext.BatchInsert(db.DB(), proxies)
	return err
}

func InvalidProxy(addr string) error {
	GetMySQLHandler().Model(Proxy{}).Where("IpPort=?", addr).Updates(
		map[string]interface{}{
			"last_fail_time": gorm.Expr("now()"),
			"total_fail":     gorm.Expr("total_fail+?", 1),
		})
	return nil
}

func AcquireProxy(addr string) error {
	GetMySQLHandler().Model(Proxy{}).Where("IpPort=?", addr).Updates(
		map[string]interface{}{
			"total_crawl": gorm.Expr("total_crawl+?", 1),
		})
	return nil

}
