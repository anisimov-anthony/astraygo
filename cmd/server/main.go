package main

import (
	"net/http"

	"github.com/anisimov-anthony/astraygo/internal/database"
	"github.com/anisimov-anthony/astraygo/internal/handlers"
	"github.com/anisimov-anthony/astraygo/internal/logging"
	"github.com/anisimov-anthony/astraygo/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Config
	logger := logging.InitLogger()
	defer logger.Sync()

	// Postgres
	pgPool := database.InitPostgres()
	defer pgPool.Close()

	// Redis
	redisClient := database.InitRedis()
	defer redisClient.Close()

	// Service
	repository := service.NewPostgresRepo(pgPool)
	cache := service.NewRedisCache(redisClient)
	cache.WarmUp(repository)

	service := service.AstrayServiceInit(repository, cache)
	handler := handlers.InitAstrayHandler(service)

	// Routing
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logging.ZapLogger(logger))

	router.GET("/objects/ids", func(c *gin.Context) {
		handler.GetAllIDs(c)
	})

	router.GET("/objects/:id", func(c *gin.Context) {
		handler.GetObjectByID(c)
	})

	router.POST("/objects", func(c *gin.Context) {
		handler.PostObject(c)
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Run(":8080")
}
