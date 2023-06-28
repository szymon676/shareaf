package main

type RediStoreOptions struct {
	Addr     string
	Password string
	DB       int
}

type Paste struct {
	Name any `json:"name"`
	Data any `json:"data"`
}
