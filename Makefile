build:
	protoc --go_out=types doc.proto

run:
	go run main.go
