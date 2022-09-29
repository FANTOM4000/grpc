# grpc
### Option https://docs.buf.build/installation
```
$ brew install grpc
$ proto --version
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ export PATH="$PATH:$(go env GOPATH)/bin"
$ protoc --go_out=./ *.proto --go-grpc_out=./ *.proto --go-grpc_opt=require_unimplemented_servers=false
$ option:buf generate
```

### If use `Buf`
##### Example buf.gen.yaml
```
version: v1
managed: 
  enabled: true
  go_package_prefix:
    default: github.com/<repo>/path_to_root_store_proto
plugins:
  - name: go
    out: .
    opt: paths=source_relative
  - name: go-grpc
    out: .
    opt: paths=source_relative,require_unimplemented_servers=false
```
##### Example buf.yaml
```
version: v1
name: buf.build/fantom400/pingpong
deps:
  - buf.build/googleapis/googleapis
  - buf.build/grpc-ecosystem/grpc-gateway
```
