## I know this is crazy but I had to add a symlink because I kept seeing this error
## protoc-gen-go_grpc: program not found or is not executable
## this was even after I had installed protoc-gen-go-grpc through brew
## The solution is to add a symlink to the protoc-gen-go-grpc binary in the $(go env GOPATH)/bin directory
## ln -s $(go env GOPATH)/bin/protoc-gen-go-grpc $(go env GOPATH)/bin/protoc-gen-go_grpc
generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go_grpc_out=.

## Run the server
run:
	cd server && go run main.go

## Run the client
client:
	cd client && go run main.go
