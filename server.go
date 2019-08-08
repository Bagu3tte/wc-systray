package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func setAlive(ip string) {
	for _, friend := range Friends {
		if friend.Ip == ip {
			friend.Alive = true
			return
		}
	}
	Friends = append(Friends, Pong{Ip: ip, Alive: true})
}

func receive(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&CurrentStatus)
	if err != nil {
		fmt.Println("Can't deserislize", r.Body)
	}
	setAlive(strings.Split(r.RemoteAddr, ":")[0])
	CurrentStatus.ReceivedAt = time.Now()
	CheckStatus()
}

func Server() {
	http.HandleFunc("/status", receive)
	go func() {
		log.Fatal(http.ListenAndServe(":"+Port, nil))
	}()
}
