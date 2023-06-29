package main

func main() {
	rediOptions := RediStoreOptions{
		Addr:     "containers-us-west-206.railway.app:6675",
		Password: "INRgMXHW1Xx0O5xtb9kq",
		DB:       0,
	}

	rediStore := NewRedisStore(rediOptions)

	h := NewApiHandler(rediStore, ":3000")
	h.Run()
}
