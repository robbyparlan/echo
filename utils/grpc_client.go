package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	pb "sip/config/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GRPCClient *pb.PaymentServiceClient
	GRPCConn   *grpc.ClientConn
)

// InitGRPCClient menginisialisasi koneksi gRPC ke Order Service
func InitGRPCClient(address string) error {
	var err error
	GRPCConn, err = grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to gRPC server: %v", err)
		return err
	}

	client := pb.NewPaymentServiceClient(GRPCConn)
	GRPCClient = &client
	return nil
}

// InitializeGRPC menginisialisasi gRPC dan mengatur penutupan koneksi
func InitializeGRPC(address string) {
	if err := InitGRPCClient(address); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	log.Println("GRPC running on :", address)
	// Menangani penutupan koneksi saat aplikasi berhenti
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Closing gRPC connection...")
		if err := GRPCConn.Close(); err != nil {
			log.Printf("Error closing gRPC connection: %v", err)
		}
		os.Exit(0)
	}()
}

func init() {
	InitializeGRPC(GRPC_HOST + ":" + GRPC_PORT)
}
