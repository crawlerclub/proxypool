package main

import (
	"flag"
	"github.com/crawlerclub/proxypool"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	wg.Add(1)
	go proxypool.Run(exitCh, &wg)

	server := &proxypool.ProxyServer{}
	wg.Add(1)
	go server.Run(exitCh, &wg)

	go server.Web()

	wg.Wait()
	glog.Info("exit!")
}
