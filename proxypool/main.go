package main

import (
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/crawlerclub/proxypool"
	"github.com/golang/glog"
)

var (
	update = flag.Bool("update", false, "proxy update flag")
)

func stop(sigs chan os.Signal, exitCh chan bool) {
	<-sigs
	glog.Info("receive stop signal!")
	close(exitCh)
}

func main() {
	flag.Parse()

	exitCh := make(chan bool)
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go stop(sigs, exitCh)

	var wg sync.WaitGroup
	if *update {
		wg.Add(1)
		go proxypool.Run(exitCh, &wg)
	}

	server := &proxypool.ProxyServer{}
	wg.Add(1)
	go server.Run(exitCh, &wg)

	go server.Web()

	wg.Wait()
	glog.Info("exit!")
}
