package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ghchinoy/gotwilio/src/gotwilio"
)

func main() {

	SID := os.Getenv("TWILIO_SID")
	AUTH := os.Getenv("TWILIO_AUTH")
	t := gotwilio.NewClient(SID, AUTH)
	//log.Printf("%+v", t)
	to := "+1 209 210 4311"
	from := "+19702359226"
	time := time.Now().Format("03:04:05")
	message := fmt.Sprintf("Greetings! The current time is: %s KBF0C586LELXQ5Y", time)
	//log.Println("sending message:", message)
	err := t.SendSMS(from, to, message)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
