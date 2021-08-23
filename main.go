package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connection to database
	dsn := "root:@tcp(127.0.0.1:3306)/learn_bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// init user repository
	userRepo := user.InstanceRepository(db)
	// init service
	userService := user.InstanceService(userRepo)
	// init handler
	userHandler := handler.InstanceUserHandler(userService)

	// router
	router := gin.Default()
	// api versioning
	apiV1 := router.Group("/api/v1")
	// route
	apiV1.POST("/users", userHandler.RegisterUser)

}
