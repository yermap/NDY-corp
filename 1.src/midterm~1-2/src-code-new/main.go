package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/controllers"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/initializers"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/product", controllers.ListProducts)
	r.POST("/product", middleware.RequireAuth, controllers.CreateProduct)
	r.POST("/comment/:id", middleware.RequireAuth, controllers.CreateComment)
	r.POST("/order", middleware.RequireAuth, controllers.CreateOrder)
	r.PATCH("/order/:id", middleware.RequireAuth, controllers.ChangeStatus)
	r.DELETE("/order/:id", middleware.RequireAuth, controllers.ChangeStatus)

	r.Run()
}
