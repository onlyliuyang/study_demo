package main

import (
	"github.com/robfig/cron"
	"log"
)

func main() {
	log.Println("Starting...")

	crontab := cron.New()
	crontab.AddFunc("0 */1 * * * ?", func() {
		log.Println("Running1 ...")
	})

	crontab.AddFunc("0 */2 * * * ?", func() {
		log.Println("Running2 ...")
	})
	crontab.Start()
	select {}
	//t1 := time.NewTimer(time.Second * 10)
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 10)
	//	}
	//}
}
