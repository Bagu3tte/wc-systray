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

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Sixmon WC")
	close()
	mQuitOrig := systray.AddMenuItem("Quitter", "Quitte l'application")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	r := req.New()
	for range time.Tick(2 * time.Second) {
		getWCStatus(r)
	}
}

func onExit() {
	fmt.Println("exit")
}

func getWCStatus(req *req.Req) {
	r, err := req.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var status Status
	err = r.ToJSON(&status)
	if err != nil {
		log.Fatal(err)
	}

	if status.Vaccant {
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
