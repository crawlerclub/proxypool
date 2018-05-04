package proxypool

import (
	"testing"
)

func TestMysql(t *testing.T) {
	db := GetMySQLHandler()
	t.Log(db)
}
