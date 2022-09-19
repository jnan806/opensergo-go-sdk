package io_opensergo

import (
	v1 "github.com/jnan806/opensergo-protocol-grpc-go/service_contract/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type OpenSergoClient struct {
}

func NewOpenSergoClient(host string, port uint8) {
	clientConn, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer clientConn.Close()

	metadataService := v1.NewMetadataServiceClient(clientConn)

	result, err := metadataService.ReportMetadata(context.Background(), &v1.ReportMetadataRequest{})
	if err != nil {
		log.Fatalf("invoke grpc error")
	}
	log.Printf("invoke result is : %v ", result)

}
