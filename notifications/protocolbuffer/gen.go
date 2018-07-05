package protocolbuffer

//go:generate protoc api.proto --go_out=plugins=grpc:.
//go:generate mockery -name=TransactionServiceClient
