package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"test/cmd/config"
	userhandler "test/cmd/handler/user"
	userrepository "test/cmd/repository/user"
	"test/cmd/routers"
	userservice "test/cmd/service/user"
	"test/pkg/database"
	"test/pkg/server"
)

func Run() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	cfg := config.LoadConfig()
	log.Println("config loaded")

	db := database.Connect(cfg.DBConfig)
	log.Println("success initialize database")

	userRepository := userrepository.NewUserRepository(db)

	userService := userservice.NewUserService(userRepository)

	userHandler := userhandler.NewUserHandler(userService)

	server := server.New(&cfg.ServerConfig)
	router := server.Router()

	routers.RegisterRouter(router, userHandler)

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	serveChan := server.Run()
	go server.Stop()

	select {
	case err := <-serveChan:
		if err != nil {
			panic(err)
		}
	case <-sigChan:
	}
}
