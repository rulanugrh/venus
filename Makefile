depedency:
	go mod download
	
run:
	go run main.go

build:
	go build -o bin/main main.go
	./bin/main

clean:
	go clean
	rm ./build/main.out