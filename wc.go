package main

import (
	"fmt"
	"log"
	"time"

	"wc-systray/icon"

	"github.com/getlantern/systray"
	"github.com/imroc/req"
)

var url = "https://e1z1mdq2mb.execute-api.eu-west-1.amazonaws.com/default/currentstatus"
var Port = "8155"
var CurrentStatus Status
var Friends []Pong

func main() {
	Server()
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Sixmon WC")
	// set to close
	close()

	mRefresh := systray.AddMenuItem("Rafraîchir", "Rafraîchir l'application")
	mQuitOrig := systray.AddMenuItem("Quitter", "Quitte l'application")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go func() {
		<-mRefresh.ClickedCh
		r := req.New()
		getWCStatus(r)
	}()

	r := req.New()
	go func(r *req.Req) {
		for range time.Tick(15 * time.Second) {
			now := time.Now()
			if (CurrentStatus != Status{} && now.Sub(CurrentStatus.ReceivedAt).Seconds() > 10) {
				go getWCStatus(r)
			}
		}
	}(r)
	Friends = FindNewFriends()
}

func onExit() {
	fmt.Println("exit")
}

func getWCStatus(req *req.Req) {
	r, err := req.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	err = r.ToJSON(&CurrentStatus)
	if err != nil {
		log.Fatal(err)
	}
	CurrentStatus.ReceivedAt = time.Now()
	CheckStatus()

	shareToFriends(req)
}

func shareToFriends(r *req.Req) {
	for i, friend := range Friends {
		if friend.Alive {
			go func(i int, friend Pong) {
				_, err := r.Post("http://"+friend.Ip+":"+Port+"/status", req.BodyJSON(&CurrentStatus))
				if err != nil {
					friend.Alive = false
				}
			}(i, friend)
		}
	}
}

func CheckStatus() {
	if CurrentStatus.Vaccant {
		open()
	} else {
		close()
	}
}

func close() {
	systray.SetIcon(icon.Close)
	systray.SetTooltip("Occupé")
}

func open() {
	systray.SetIcon(icon.Open)
	systray.SetTooltip("Libre")
}
