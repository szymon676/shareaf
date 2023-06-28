package main

func main() {
	rediOptions := RediStoreOptions{
		Addr:     "containers-us-west-148.railway.app:6377",
		Password: "y1bdYc5SBR1wv3LHN6ai",
		DB:       0,
	}

	rediStore := NewRedisStore(rediOptions)

	h := NewApiHandler(rediStore, ":3000")
	h.Run()
}
