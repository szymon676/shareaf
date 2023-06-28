package main

func main() {
	rediOptions := RediStoreOptions{
		Addr:     "addr",
		Password: "password",
		DB:       0,
	}

	rediStore := NewRedisStore(rediOptions)

	h := NewApiHandler(rediStore, ":3000")
	h.Run()
}
