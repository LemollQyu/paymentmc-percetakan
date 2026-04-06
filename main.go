package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"paymentmc/cmd/app/handler"
	"paymentmc/cmd/app/repository"
	"paymentmc/cmd/app/resource"
	"paymentmc/cmd/app/service"
	"paymentmc/cmd/app/storage"
	"paymentmc/cmd/app/usecase"
	"paymentmc/config"
	"paymentmc/middleware"
	"paymentmc/proto/paymentpb"
	"syscall"
	"time"

	grpcServerPkg "paymentmc/infrastructure/grpc"

	notificationGrpc "paymentmc/grpc/notification"
	orderGrpc "paymentmc/grpc/order"
	userGrpc "paymentmc/grpc/user"

	"paymentmc/infrastructure/log"
	"paymentmc/routes"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Service Payment")

	cfg := config.LoadConfig()
	fmt.Println("Config semua disni")
	fmt.Println("APP CONFIG:", cfg.App)
	fmt.Println("DATABASE CONFIG:", cfg.Database)
	fmt.Println("PATH UPLOADS:", cfg.Storage)
	fmt.Println("GRPC Config: ", cfg.GRPC)
	fmt.Println("KAFKA Config: ", cfg.Kafka)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := resource.InitDB(&cfg)

	log.SetupLogger()

	paymentRepostory := repository.NewPaymentRepository(db)

	orderGRPC := orderGrpc.NewOrderClient(cfg.GRPC.OrderURL)
	notificationGRPC := notificationGrpc.NewNotificationClient()
	userGRPC := userGrpc.NewUserClient()

	paymentService := service.NewPaymentService(*paymentRepostory, orderGRPC, notificationGRPC, userGRPC)
	paymentStorage := storage.NewStorage(cfg.Storage.UploadBaseDir, cfg.App.Url)
	paymentUsecase := usecase.NewPaymentUsecase(
		*paymentService,
		*paymentStorage,
	)

	//  Jalankan Worker set status order code yang expired
	paymentUsecase.StartBrokerExpiredPaymentAndOrder(ctx)

	paymentHandler := handler.NewPaymentHandler(*paymentUsecase)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CORS([]string{"http://localhost:3000", "http://localhost:3001"}))
	router.Use(gin.Logger())
	router.Static("/static", "./uploads")
	routes.SetupRoutes(router, *paymentHandler, cfg.Secret.JWTSecret)

	srv := &http.Server{
		Addr:    ":" + cfg.App.Port,
		Handler: router,
	}

	go func() {
		log.Logger.Infof("Server jalan di port : %s", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatalf("Gagal menjalankan server: %s\n", err)
		}
	}()

	// ---- gRPC SERVER ----
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Logger.Fatalf("Failed to listen gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()
	paymentGRPCServer := &grpcServerPkg.GRPCServer{
		PaymentUsecase: *paymentUsecase,
	}

	paymentpb.RegisterPaymentServiceServer(
		grpcServer,
		paymentGRPCServer,
	)

	for service, info := range grpcServer.GetServiceInfo() {
		log.Logger.Println("gRPC Service:", service)
		for _, method := range info.Methods {
			log.Logger.Println("  └─ Method:", method.Name)
		}
	}

	log.Logger.Println("gRPC server running on port :50054")

	if err := grpcServer.Serve(lis); err != nil {
		log.Logger.Fatalf("Failed to serve gRPC: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Logger.Info("Menerima sinyal berhenti, mematikan server...")

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Logger.Fatal("Server dipaksa berhenti: ", err)
	}

	log.Logger.Info("Server keluar dengan aman (Graceful Shutdown Berhasil).")
}
