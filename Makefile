depedency:
	go mod download
	
run:
	go run cmd/main.go

build:
	go build -o build/main.out cmd/main.go
	./build/main.out

clean:
	go clean
	rm ./build/main.out