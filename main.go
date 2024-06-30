package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	sigInt := make(chan os.Signal, 1)
	signal.Notify(sigInt, os.Interrupt)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var last_time = time.Now().Format("15:04")
	fmt.Println(last_time)

	for {
		select {
		case <-sigInt:
			return
		case t := <-ticker.C:
			var time_now = t.Format("15:04")

			if time_now != last_time {
				fmt.Println(time_now)
				last_time = time_now
			}
		}
	}
}
