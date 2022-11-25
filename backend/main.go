package main

import (
	"backend/constant"
	"backend/infra"
	"backend/infra/repository/article"
	"backend/server"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	domainrepo "backend/domain/repository"

	"github.com/spf13/viper"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 環境変数読み込み
	loadEnv()
	// DB初期化処理
	gormHandler := infra.NewGormHandler()

	// GRPC接続
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	repos := domainrepo.Repositories{
		ArticleRepository: article.NewArticleRepository(gormHandler),
	}

	grpcServer := server.NewGRPCServer(repos)
	go func() {
		<-sigCh
		log.Printf("shutdown signal is called")
		grpcServer.GracefulStop()
	}()

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server start!")
	// サーバーリフレクションを有効にしています。
	// 有効にすることでシリアライズせずとも後述する`grpc_cli`で動作確認ができるようになります。
	reflection.Register(grpcServer)
	serverErr := grpcServer.Serve(listener)

	if serverErr != nil {
		log.Fatal(serverErr)
	}
}

func loadEnv() {
	viper.AutomaticEnv()
	viper.SetDefault(constant.DBHostEnv, "db")
	viper.SetDefault(constant.DBUserEnv, "root")
	viper.SetDefault(constant.DBNameEnv, "dev")
}
