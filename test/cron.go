package test

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"time"
)

func Cron() {
	c := cron.New()
	c.AddFunc("* * * * * *", RunEverySecond)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func RunEverySecond() {
	fmt.Printf("%v\n", time.Now())
}
