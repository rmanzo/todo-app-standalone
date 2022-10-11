run:
	go run ./cmd/reprogramable

add:
	go run ./cmd/reprogramable -add 180

list:
	go run ./cmd/reprogramable -list

cleft:
	go run ./cmd/reprogramable -cleft

build:
	go build -o bin/main ./cmd/reprogramable