package main

import (
	"context"
	"fmt"
	"log"

	"ecommerce-backend/internal/config"
	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/pkg/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	cfg := config.Load()
	client := mongodb.Connect(cfg.MongoURI)
	db := client.Database(cfg.DBName)

	var user domain.User
	err := db.Collection("users").FindOne(context.Background(), bson.M{"email": "rahmatrafiindrayani@gmail.com"}).Decode(&user)
	if err != nil {
		log.Fatalf("DECODE ERROR: %v", err)
	}
	fmt.Printf("User ID: %v\n", user.ID)
}
