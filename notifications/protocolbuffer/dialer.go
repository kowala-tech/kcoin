package protocolbuffer

import "google.golang.org/grpc"

func DialNewTransactionServiceClient(endpoint string) TransactionServiceClient {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		panic("Error connection to notifications endpoint: " + err.Error())
	}

	return NewTransactionServiceClient(conn)
}
