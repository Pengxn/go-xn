package proto

//go:generate protoc --version
//go:generate protoc-gen-go --version
//go:generate protoc-gen-go-grpc --version
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto

// Use `go generate ./...` to generate protobuf files and gRPC code.
// The generated files will be placed in the same directory as the source files.
// Make sure you have installed the protoc, protoc-gen-go and protoc-gen-go-grpc commands.
// If not, you can install them by running:
// 	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
// 	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
