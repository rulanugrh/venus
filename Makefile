depedency:
	go mod download
	
run:
	go run cmd/main.go

build:
	go build -o bin/main cmd/main.go
	./bin/main

clean:
	go clean
	rm ./build/main.out