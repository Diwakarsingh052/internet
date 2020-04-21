package notify

import (
	"gopkg.in/toast.v1"
	"log"
	"os"
)

func Notification() {
	mydir, _ := os.Getwd()
	notification := toast.Notification{
		AppID:   "Internet Lost",
		Title:   "Internet Lost",
		Message: "Wifi Signal Lost",
		Icon:    mydir+`\emergency.png`, // This file must exist (remove this line if it doesn't)
		Actions: []toast.Action{
			{"protocol", "Dismiss", ""},
		},
	}
	if err := notification.Push(); err != nil {
		log.Fatalln(err)
	}
}
