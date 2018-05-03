package proxypool

import (
	"testing"
)

func TestMysql(t *testing.T) {
	db := GetMySQLHandler()
	t.Log(db)

	ret, err := ReadProxy()
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
