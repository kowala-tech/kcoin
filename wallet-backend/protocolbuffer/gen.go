package protocolbuffer

//go:generate protoc -I ../../notifications/protocolbuffer api.proto --go_out=plugins=grpc:.
//go:generate mockery -name=TransactionServiceClient
