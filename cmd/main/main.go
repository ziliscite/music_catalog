package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"music_catalog/internal/config"
	membershipHandler "music_catalog/internal/handler/membership"
	membershipModel "music_catalog/internal/models/membership"
	membershipRepository "music_catalog/internal/repository/membership"
	membershipService "music_catalog/internal/service/membership"
	"music_catalog/pkg"
	"net/http"
)

func main() {
	cfg := NewConfig()

	db, err := pkg.ConnectDB(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// migrations with gorm -- will create tables
	err = db.AutoMigrate(&membershipModel.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	router := NewRouter()

	NewMembership(router, db, cfg)

	_ = router.Run(cfg.Service.Port)
}

func NewMembership(router *gin.Engine, db *gorm.DB, cfg *config.Config) {
	repository := membershipRepository.NewRepository(db)
	service := membershipService.NewService(cfg, repository)
	handler := membershipHandler.NewHandler(router, service)
	handler.RegisterRoutes()
}

func NewConfig() *config.Config {
	err := config.Init(
		config.WithConfigFolder([]string{
			"./",
			"./config/",
		}),
		config.WithConfigFile("config"),
		config.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to initialize configs: %v", err)
	}
	return config.Get()
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.ErrorLogger())
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return router
}
