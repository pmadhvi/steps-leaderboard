package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Hi Welcome to Steps Leaderboard Application")

	http.HandleFunc("/stepcounter", stepCounterHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("The Steps Leaderboard application crashed")
	}
}
