package main

import (
	"github.com/diwakarsingh052/internet/notify" //local directory
	"log"
	"net"
	"time"
)

func main() {
	var (
		hostName = "google.com"
		portNum  = "80"
		seconds  = 2
		counter  = 0
		timeOut  = time.Duration(seconds) * time.Second
	)
	log.Print("Application started minimize the application and do your work")
	for {
		_, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)

		if err != nil {
			if counter != 2 {
				counter++
				time.Sleep(1 * time.Second)
				continue
			}

			notify.Notification()
			return
		}

		counter = 0
		time.Sleep(2 * time.Second)
	}

}
