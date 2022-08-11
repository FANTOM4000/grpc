# grpc
# brew install grpc
# proto --version
# go install google.golang.org/protobuf/cmd/protoc-gen-go
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
# export PATH="$PATH:$(go env GOPATH)/bin"
# protoc --go_out=./ *.proto --go-grpc_out=./ *.proto --go-grpc_opt=require_unimplemented_servers=false
# option:buf generate