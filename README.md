#### ASI User Management Service

#### Run protoc
```bash
protoc --go_out=generated/go/asi-user-management-service --go_opt=paths=source_relative --go-grpc_out=generated/go/asi-user-management-service --go-grpc_opt=paths=source_relative proto/user.proto
```

#### Run server
```bash
make run
```