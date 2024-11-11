package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"music_catalog/internal/config"
	"music_catalog/pkg"
	"net/http"
)

func main() {
	cfg := NewConfig()

	_, err := pkg.ConnectDB(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	router := NewRouter()
	_ = router.Run(cfg.Service.Port)
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
