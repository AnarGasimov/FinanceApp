package main

import (
	"FinanceApp/config"
	hlr "FinanceApp/internal/account/handler"
	svc "FinanceApp/internal/account/services"
	db "FinanceApp/internal/db/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error occurred while loading config: %v\n", err)
	}
	router := gin.Default()

	sqlDB, err := sql.Open(conf.Database.Driver, conf.Database.URL)

	if err != nil {
		log.Fatalf("Occured an error with connection to DB %v\n", err)
	}

	dbStore := db.NewStore(sqlDB)

	accountService := svc.NewAccountServiceImp(dbStore)
	accountHandler := hlr.NewAccountHandler(accountService)
	accountHandler.RegisterAccountRoutes(&router.RouterGroup)

	if err := router.Run(conf.Server.Address); err != nil {
		log.Fatalf("Error occured while starting the server: %v\n", err)
	}
}
