package main

import (
	// "context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/levelstudio/payroll-4ta-crud/pkg/db"
	"github.com/levelstudio/payroll-4ta-crud/pkg/models"
	"github.com/levelstudio/payroll-4ta-crud/pkg/server"
	"github.com/levelstudio/payroll-4ta-crud/pkg/utils"
)

func main() {

	// ctx := context.Background()

	ServerDoneChan := make(chan os.Signal, 1)
	signal.Notify(ServerDoneChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		// err := server.ListenAndServe()
		db.DBConnection()
		
		db.DB.AutoMigrate(models.User{})

		if err := utils.EnsureDir("files"); err != nil {
			panic(err)
		}
		err := server.HttpServer(":8080")
		if err != nil {
			panic(err)
		}
	}()

	log.Println("server Started")

	<-ServerDoneChan

	// http.Serve.Shutdown(ctx)
	log.Println("server Closed")

}
