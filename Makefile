build: 
	@go build -o main.exe main.go
run: build
	@main.exe