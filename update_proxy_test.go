package proxypool

import (
	"testing"
)

func TestValidate(t *testing.T) {
	addr := "120.92.88.202:10000"
	ret := Validate(addr)
	t.Log(ret)
}
