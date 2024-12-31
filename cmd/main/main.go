package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"music_catalog/internal/config"
	membershipHandler "music_catalog/internal/handler/membership"
	trackHandler "music_catalog/internal/handler/track"
	membershipModel "music_catalog/internal/model/membership"
	membershipRepository "music_catalog/internal/repository/membership"
	spotifyRepository "music_catalog/internal/repository/spotify"
	membershipService "music_catalog/internal/service/membership"
	trackService "music_catalog/internal/service/track"
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
	NewTrack(router, cfg)

	_ = router.Run(cfg.Service.Port)
}

func NewTrack(router *gin.Engine, cfg *config.Config) {
	client := pkg.NewClient(&http.Client{})
	repository := spotifyRepository.NewRepository(cfg, client)
	service := trackService.NewService(cfg, repository)
	handler := trackHandler.NewHandler(router, service)
	handler.RegisterRoutes()
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
