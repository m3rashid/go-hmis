build:
	go build -o bin/hmis main.go

run:
	go run main.go

prod:
	go build -o bin/hmis main.go
	./bin/hmis
