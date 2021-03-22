## GRPC Docs

~~I have no idea what to write here, since gRPC is already documented through the protos file~~

The protos are under protos/, and there you can see the whole API description.

The protos mirror the APIs defined on `pkg/repository/plant.go`, and thus all data types are based on the function headers there.

The generated RentitServerInterface is implemented on `pkg/transport/grpc/plant.go`, and an example client is under `test_clients/grpcClient.go`
