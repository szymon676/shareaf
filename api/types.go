package main

import "time"

type RediStoreOptions struct {
	Addr     string
	Password string
	DB       int
}

type Paste struct {
	Name string        `json:"name"`
	Data any           `json:"data"`
	Time time.Duration `json:"time"`
}
