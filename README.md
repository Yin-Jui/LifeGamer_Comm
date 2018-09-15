# LifeGamer_Comm
Communicated module of LifeGamer.

## Setup 

* Install dependencies
```bash
# grpc
go get -u google.golang.org/grpc

# protobuf plugin
go get -u github.com/golang/protobuf/protoc-gen-go
```

* Setup compiler plugin, protoc-gen-go
```bash
export PATH=$PATH:$GOPATH/bin
```

## Compile protobuf

```bash
protoc -I proto/ proto/<spec>.proto --go_out=plugins=grpc:proto
```