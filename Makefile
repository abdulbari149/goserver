build: main.go user.go
	@mkdir -p build
	@go build -o build/goserver main.go user.go

run: build
	@./build/goserver

