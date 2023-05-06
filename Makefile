proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./grpc_stuff/schema.proto

clean: 
	go clean 
	rm ${BINARY_NAME}_linux_amd64

dep: 
	go mod tidy && go mod vendor && go fmt

build: 
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}_linux_amd64 server.go
