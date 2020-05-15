package main

import "github.com/gin-gonic/gin"

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
