package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	r *gin.Engine
)

func main() {
	r = gin.Default()

	api := r.Group("/api")
	api.GET("/people", Read)
	api.POST("/people", Create)
	api.PUT("/people", Update)
	api.DELETE("/people/:id", Delete)

	r.Run()
}
