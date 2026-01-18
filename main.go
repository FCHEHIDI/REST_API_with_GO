package main

import (
	"github.com/gin-gonic/gin"
	"restapi-go/db"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Run(":8080") // localhost port 8080
}
