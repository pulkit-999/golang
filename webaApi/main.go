package main

import (
	"context"
	"webapi/webaApi/config"
	"webapi/webaApi/controller"
	"webapi/webaApi/repository"
	"webapi/webaApi/routes"
	"webapi/webaApi/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	router := gin.New()
	db := config.Connect()

	// Redis connection
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:5000",
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic("Redis connection failed")
	}

	userRepo := repository.NewEmployeeRepository(db)
	userService := service.NewEmployeeService(userRepo, redisClient)
	userCtrl := controller.NewEmployeeController(userService)

	routes.EmployeeRoutes(router, userCtrl)

	router.Run(":8000")
}
