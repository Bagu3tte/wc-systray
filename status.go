package main

import "time"

type Status struct {
	LastStatusChange int64     `json:"LastStatusChange"`
	Vaccant          bool      `json:"Vaccant"`
	ReceivedAt       time.Time `json:"-"`
}
