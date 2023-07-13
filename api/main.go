package main

import "os"

func main() {
	rediOptions := RediStoreOptions{
		Addr:     "",
		Password: "",
		DB:       0,
	}

	rediStore := NewRedisStore(rediOptions)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	h := NewApiHandler(rediStore, port)
	h.Run()
}
