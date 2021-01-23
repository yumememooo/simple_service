package main

import (
	"fmt"
	"os"
	"os/signal"
	"simple_service/internal/config"
	"simple_service/internal/log"
	"syscall"
	"time"
)

func main() {
	start := time.Now()
	if err := config.LoadConfig(); err != nil {
		fmt.Println("read config fail:", err)
		os.Exit(1) //退出程序
	}
	log.InitLogger()
	fmt.Println("start program:", config.Configuration.Service.StartupMsg)
	fmt.Println("connect program:", config.Configuration.Services["Elastic"].Host)
	log.ZapLogger.Warn("inital ok")
	errs := make(chan error, 1)
	listenForSignal(errs)

	//Since返回t 到现在經過的时间
	fmt.Println("Service started in:", time.Since(start))

	c := <-errs
	fmt.Println("terminating:", c)
}

func listenForSignal(errChan chan error) {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		errChan <- fmt.Errorf("%s", <-c)
	}()
}
