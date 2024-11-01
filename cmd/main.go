package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tegarpratama/checkout-service/internal/configs"

	userHandler "github.com/tegarpratama/checkout-service/internal/handler/users"
	userRepo "github.com/tegarpratama/checkout-service/internal/repository/users"
	userSvc "github.com/tegarpratama/checkout-service/internal/service/users"

	transactionHandler "github.com/tegarpratama/checkout-service/internal/handler/transactions"
	transactionRepo "github.com/tegarpratama/checkout-service/internal/repository/transactions"
	transactionSvc "github.com/tegarpratama/checkout-service/internal/service/transactions"

	"github.com/tegarpratama/checkout-service/pkg/internalsql"
	"github.com/tegarpratama/checkout-service/pkg/seeder"
)

func main() {
	r := gin.Default()

	var cfg *configs.Config

	err := configs.Init(configs.Option{
		ConfigFolder: "./internal/configs",
		ConfigFile:   "config",
		ConfigType:   "yaml",
	})

	if err != nil {
		log.Fatal("failed initiation config", err)
	}

	cfg = configs.Get()
	log.Println("configuration loaded")

	db, err := internalsql.ConnectDB(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("failed connect into database", err)
	}

	log.Println("database connected")

	seeder.SeedProducts(db)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")

	api.GET("/check-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it's works!",
		})
	})

	userRepo := userRepo.NewRepository(db)
	userService := userSvc.NewService(cfg, userRepo)
	userHandler := userHandler.NewHandler(api, userService)
	userHandler.RegisterRoute()

	transactionRepo := transactionRepo.NewRepository(db)
	transactionService := transactionSvc.NewService(cfg, transactionRepo)
	transactionHandler := transactionHandler.NewHandler(api, transactionService)
	transactionHandler.RegisterRoute()

	// server := fmt.Sprintf("127.0.0.1%s", cfg.Service.Port)
	// r.Run(server)
	r.Run(cfg.Service.Port)
}
