package main

import "log"

func main() {
	rediOptions := RediStoreOptions{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	rediStore := NewRedisStore(rediOptions)
	err := rediStore.SavePaste("idkPaste", "some data in here...")
	if err != nil {
		log.Fatal(err)
	}

	paste, err := rediStore.RetrievePaste("idkPaste")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("paste:", paste)
}
