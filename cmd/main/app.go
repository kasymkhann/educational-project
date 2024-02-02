package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"

	"REST/internal/configs"
	"REST/pkg/logging"
)

var (
	listener    net.Listener
	errListener error
)

func main() {

	//logger := logging.GetLogger()
	//logger.Info("create router")
	//router := httprouter.New()
	//
	//cfg := configs.GetConfig()
	//
	//mongoClient, err := mongodb.NewClients(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB)
	//if err != nil {
	//	panic(err)
	//}

	//user := users.CreateUsers{
	//
	//	Email:    "TestEmail",
	//	Username: "testUsername",
	//}
	//
	//if err != nil {
	//	panic(err)
	//}
	//logger.Info(user)
	//
	//logger.Info("register handler")
	//handler := users.NewHandler(logger)
	//handler.Register(router)
	//
	//start(router, cfg)
}

func start(router *httprouter.Router, cfg *configs.Config) {
	logger := logging.GetLogger()
	logger.Info("start server")

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("create unix socket ")
		listener, errListener = net.Listen("unix", socketPath)
		if errListener != nil {
			logger.Fatal(errListener)
		}
	} else {
		logger.Info("listener tcp ")
		listener, errListener = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		if errListener != nil {
			logger.Fatal(errListener)
		}
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	logger.Infof("server is listening port: %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	logger.Fatal(server.Serve(listener))

}
