package main

import (
	configurations "assignment/config"
	"assignment/db"
	"assignment/handler"
	"assignment/service/health"
	"assignment/service/rider"
	"assignment/service/restaurant"

	"assignment/service/user"
	"context"
	"log"
)

func main() {
	config := configurations.Init("config")
	port := ":" + config.GetString("service.port")

	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	//mongodb
	client, errClient := db.InitMongoClient(config)
	if errClient != nil {
		log.Println("errClient:", errClient)
	}
	defer client.Disconnect(context.Background())

	healthService := health.NewService()
	userService := user.NewUserService(config, client)
	RiderService := rider.NewRiderService(config, client)
	RestaurantService := restaurant.NewRestaurantService(config, client)

	// Start server
	e := handler.MakeHTTPHandler(config, healthService, userService, RiderService,RestaurantService)
	e.Logger.Fatal(e.Start(port))
}
