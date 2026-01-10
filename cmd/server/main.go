package main

import (
	"github.com/anisimov-anthony/astraygo/internal/database"
	"github.com/anisimov-anthony/astraygo/internal/handlers"
	"github.com/anisimov-anthony/astraygo/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Config
	gin.SetMode(gin.DebugMode)

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
	router.Use(gin.Logger())

	router.GET("/ids", func(c *gin.Context) {
		handler.GetAllIDs(c)
	})

	router.GET("/objects/:id", func(c *gin.Context) {
		handler.GetObjectByID(c)
	})

	router.POST("/objects", func(c *gin.Context) {
		handler.PostObject(c)
	})

	router.Run(":8080")
}
