package proxypool

import (
	"github.com/jinzhu/gorm"
	"time"
)

func ReadProxy() (records []*Proxy, err error) {
	err = GetMySQLHandler().Where(
		"update_time>=last_fail_time OR total_crawl<5 OR total_fail/total_crawl<0.5").
		Find(&records).Error
	return
}

func InsertProxyStr(proxies []string) error {
	db := GetMySQLHandler()
	for _, proxy := range proxies {
		if err := db.Set("gorm:insert_option",
			"ON DUPLICATE KEY UPDATE update_time=now()").
			Create(&Proxy{IpPort: proxy,
				UpdateTime:   time.Now(),
				LastFailTime: time.Unix(0, 0)}).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertProxy(proxies []*Proxy) error {
	db := GetMySQLHandler()
	for _, proxy := range proxies {
		if err := db.Set("gorm:insert_option",
			"ON DUPLICATE KEY UPDATE update_time=now()").
			Create(proxy).Error; err != nil {
			return err
		}
	}
	return nil
}

func InvalidProxy(addr string) error {
	return GetMySQLHandler().Model(Proxy{}).Where("ip_port=?", addr).Updates(
		map[string]interface{}{
			"last_fail_time": gorm.Expr("now()"),
			"total_fail":     gorm.Expr("total_fail+?", 1),
		}).Error
}

func AcquireProxy(addr string) error {
	return GetMySQLHandler().Model(Proxy{}).Where("ip_port=?", addr).Updates(
		map[string]interface{}{
			"total_crawl": gorm.Expr("total_crawl+?", 1),
		}).Error
}
