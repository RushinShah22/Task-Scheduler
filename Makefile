
aLl: build run


build: ./main.go
	go build -o ./bin/main ./main.go

run: build
	./bin/main