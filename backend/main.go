package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lits-06/go-auth/database"
	"github.com/lits-06/go-auth/routes"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Origin", "Content-Type", "Accept", "Connection", "Accept-Encoding", "Accept-Language", "Content-Length"},
		AllowCredentials: true,
	}))
	
	routes.Setup(r)

	r.Run(":8000") 
}
