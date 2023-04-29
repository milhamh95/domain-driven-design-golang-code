# Chapter 2

## Published Language

### Open API

- Open and enter folder `chapter2/oapi`
- Run in terminal

```
api-codegen --config=config.yml ./oapi.yaml
```

### gRPC

- Install underlying `protobuf` tools

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# update path

export PATH="$PATH:$(go env GOPATH)/bin"
```

- Install `buf`

```
brew install bufbuild/buf/buf
```

- Open and enter folder `chapter2/grpc`
- To lint proto buf, user

```
buf lint
```

- To generate gRPC client and server

```
buf generate
```