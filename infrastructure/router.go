package infrastructure

import (
	"fmt"
	"log"
	"net"
	"os"

	todov1 "github.com/nagisa599/golang-grpc-template/gen/go/v1/todo"
	"github.com/nagisa599/golang-grpc-template/internal/handler"
	"google.golang.org/grpc"
)

func Router() {
	// databaseHandler := NewDatabaseHandler()
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
	
	todoHandler := handler.NewTodoHandler()
	srv := grpc.NewServer()
	todov1.RegisterTodoServiceServer(srv, todoHandler)
		// ログを出力するmiddlewareを実行
	if os.Getenv("ENV") == "development" {
		log.Print("development")
		// srv.AroundOperations(middleware.LoggerMiddleware)
	}
	
	if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
	}
    
}