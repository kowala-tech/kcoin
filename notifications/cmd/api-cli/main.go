package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := flag.String("addr", "localhost:3000", "grpc server address")
	op := flag.String("o", "", "operation to run (register/unregister)")
	wallet := flag.String("w", "", "ethereum wallet to register or unregister")
	email := flag.String("e", "", "e-mail address to register")

	flag.Parse()

	if *op != "register" && *op != "unregister" {
		fmt.Println("Invalid operation. Must be either `register` or `unregister`")
		os.Exit(1)
	}
	if *wallet == "" {
		fmt.Println("Invalid wallet")
		os.Exit(1)
	}

	if *op == "register" && *email == "" {
		fmt.Println("Invalid email")
		os.Exit(1)
	}

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error connecting to the GRPC server: %v", err)
		os.Exit(1)
	}

	client := protocolbuffer.NewEmailMappingClient(conn)

	ctx, done := context.WithTimeout(context.Background(), 10*time.Second)
	defer done()
	switch *op {
	case "register":
		_, err := client.Register(ctx, &protocolbuffer.RegisterRequest{
			Email:  *email,
			Wallet: *wallet,
		})
		if err != nil {
			fmt.Printf("Error registering wallet-email mapping: %v", err)
			os.Exit(1)
		}
	case "unregister":
		_, err := client.Unregister(ctx, &protocolbuffer.UnregisterRequest{
			Wallet: *wallet,
		})
		if err != nil {
			fmt.Printf("Error unregistering wallet mapping: %v", err)
			os.Exit(1)
		}
	}
}
