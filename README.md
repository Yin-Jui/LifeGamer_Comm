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

* support category:
    * event

```bash
# compile event's pb
protoc -I proto/<category> proto/<category>/<spec>.proto --go_out=plugins=grpc:proto-<category>
```

## Tool Set

* `src/tools`:
    * `protoc.exe`: executable protobuf compiler for windows.


## Server/Client 

* Server start first:
```bash
$ go run src/comm_engine.go
```

* Then activate client:
```bash
$ go run test/dummy_client.go
```

## Author 

* Kevin Cyu, kevinbird61@gmail.com
