package proxypool

import (
	"testing"
	"time"
)

func TestReadProxy(t *testing.T) {
	ret, err := ReadProxy()
	if err != nil {
		t.Error(err)
	}
	for _, p := range ret {
		t.Logf("%+v", p.IpPort)
	}
}

func TestInsertProxy(t *testing.T) {
	proxies := []*Proxy{
		&Proxy{IpPort: "127.0.0.1:8080",
			UpdateTime:   time.Now(),
			LastFailTime: time.Unix(0, 0)}}
	if err := InsertProxy(proxies); err != nil {
		t.Error(err)
	}
}

func TestInvalidProxy(t *testing.T) {
	addr := "127.0.0.1:8080"
	if err := InvalidProxy(addr); err != nil {
		t.Error(err)
	}
}

func TestAcquireProxy(t *testing.T) {
	addr := "127.0.0.1:8080"
	if err := AcquireProxy(addr); err != nil {
		t.Error(err)
	}
}
