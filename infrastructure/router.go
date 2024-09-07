package infrastructure

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	todov1 "github.com/nagisa599/golang-grpc-template/gen/go/v1/todo"
	"github.com/nagisa599/golang-grpc-template/internal/domain/repository"
	"github.com/nagisa599/golang-grpc-template/internal/handler"
	"github.com/nagisa599/golang-grpc-template/internal/usecase"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func Router() {
	databaseHandler := NewDatabaseHandler()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set.")
	}
	// 直接ポート文字列を使用
	address := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp",  address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	todoHandler := handler.NewTodoHandler(usecase.NewTodoUsecase(repository.NewTodoRepository(databaseHandler)))
	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
				return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	var zapLogger *zap.Logger
	if os.Getenv("ENV") == "development" {
		// Development environment: detailed logging
		zapLogger, _ = zap.NewDevelopment()
	} else {
		// Production environment: minimal logging
		zapLogger, _ = zap.NewProduction()
	}
	grpc_zap.ReplaceGrpcLogger(zapLogger)
	srv := grpc.NewServer(   
		grpc_middleware.WithUnaryServerChain(
        grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
        grpc_zap.UnaryServerInterceptor(zapLogger, opts...),        
	),)

	todov1.RegisterTodoServiceServer(srv, todoHandler)
		// ログを出力するmiddlewareを実行
	
	if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
	}
    
}