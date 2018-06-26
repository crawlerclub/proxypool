package proxypool

import (
	"flag"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/liuzl/goutil"
	"github.com/liuzl/goutil/rest"
)

var (
	addr = flag.String("addr", ":8118", "bind address")
)

type ProxyServer struct {
	sync.RWMutex
	ProxyList []*Proxy
	N         int
}

func (s *ProxyServer) ReadProxy() {
	s.Lock()
	defer s.Unlock()
	var err error
	if s.ProxyList, err = ReadProxy(); err != nil {
		glog.Error(err)
	} else {
		s.N = len(s.ProxyList)
	}
}

func (s *ProxyServer) Get() string {
	s.RLock()
	defer s.RUnlock()
	if s.N > 0 {
		i := rand.Intn(s.N)
		p := s.ProxyList[i]
		if err := AcquireProxy(p.IpPort); err != nil {
			glog.Error(err)
		}
		return p.IpPort
	}
	return ""
}

func (s *ProxyServer) Run(exitCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-exitCh:
			return
		default:
			glog.Info("reloading from db...")
			s.ReadProxy()
			glog.Info("reload done!")
			goutil.Sleep(600*time.Second, exitCh)
		}
	}
}

func (s *ProxyServer) Web() {
	http.Handle("/get", rest.WithLog(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(s.Get()))
	}))

	replacer := strings.NewReplacer("http://", "", "https://", "")

	http.Handle("/bad", rest.WithLog(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		p := r.FormValue("p")
		glog.Infof("proxy %s invalid", p)
		if p != "" {
			p = replacer.Replace(p)
			if err := InvalidProxy(p); err != nil {
				glog.Error(err)
				w.Write([]byte("err"))
			} else {
				w.Write([]byte("ok"))
			}
		}
	}))

	glog.Info("server listen on", *addr)
	glog.Error(http.ListenAndServe(*addr, nil))
}
