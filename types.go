package proxypool

import (
	"time"
)

type Proxy struct {
	Id           uint64    `json:"id" gorm:"id"`
	IpPort       string    `json:"ip_port" gorm:"ip_port"`
	UpdateTime   time.Time `json:"update_time" gorm:"update_time"`
	LastFailTime time.Time `json:"last_fail_time" gorm:"last_fail_time"`
	TotalCrawl   uint64    `json:"total_crawl" gorm:"total_crawl"`
	TotalFail    uint64    `json:"total_fail" gorm:"total_fail"`
}
