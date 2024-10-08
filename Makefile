gen:
	protoc --proto_path=proto proto/*.proto --go_out=paths=source_relative:pb --go-grpc_out=paths=source_relative:pb

clean:
	rm -rf pb/*.go

run:
	go run main.go

test:
	go test --cover --race ./...