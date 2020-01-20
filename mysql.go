package proxypool

import (
	"flag"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	mysqlUri = flag.String("mysql_uri",
		"root:@/proxy?charset=utf8&parseTime=True&loc=Local",
		"mysql uri")

	mysqlDB *gorm.DB
	mu      sync.Mutex
	once    sync.Once
)

func GetMySQLHandler() *gorm.DB {
	once.Do(func() {
		var err error
		mysqlDB, err = gorm.Open("mysql", *mysqlUri)
		if err != nil {
			panic(err)
		}
	})
	return mysqlDB
}
