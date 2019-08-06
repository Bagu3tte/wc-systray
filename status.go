package main

type Status struct {
	LastStatusChange int64 `json:"LastStatusChange"`
	Vaccant          bool  `json:"Vaccant"`
}
