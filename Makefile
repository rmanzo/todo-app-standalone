run:
	go run ./cmd/cli

add:
	go run ./cmd/cli -add "This is a test task"

list:
	go run ./cmd/cli -list

cleft:
	go run ./cmd/cli -cleft

buildcli:
	go build -o bin/cli ./cmd/cli

buildserver:
	go build -o bin/server ./cmd/server