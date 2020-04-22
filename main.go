package main

import (
	"github.com/diwakarsingh052/internet/notify" //local directory
	"log"
	"net"
	"time"
)



func main() {
	hostName := "google.com"
	portNum := "80"
	seconds := 3
	timeOut := time.Duration(seconds) * time.Second
	log.Print("Application started minimize the application and do your work")
	for  {
		_, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)

		if err != nil {
			notify.Notification()
			return
		}
		time.Sleep(2 * time.Second)
	}

}
