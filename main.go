package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/router"
	"bwastartup/user"
	"bwastartup/campaign"
	"log"

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
	campaignRepo := campaign.InstanceRepository(db)

	// init service
	userService := user.InstanceService(userRepo)
	authService := auth.InstanceService()
	campaignService := campaign.InstanceService(campaignRepo)

	// init handler
	userHandler := handler.InstanceUserHandler(userService, authService)
	campaignHandler := handler.InstanceCampaignHandler(campaignService)
	authMiddleware := middleware.InstanceAuthMiddleware(userService, authService)

	// router
	router := router.InstanceRouter(userHandler, campaignHandler, authMiddleware)
	router.Router()

}
