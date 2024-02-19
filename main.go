package main

import (
	"githiub.com/Eswarakash/GoJWTFramework/controller"
	"githiub.com/Eswarakash/GoJWTFramework/initializer"
	"githiub.com/Eswarakash/GoJWTFramework/middleware"
	"githiub.com/Eswarakash/GoJWTFramework/migrate"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	migrate.MigrateAuto()
}

func main() {

	r := gin.Default()

	r.GET("/ping", controller.Ping)
	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)
	r.GET("/validate", middleware.RequireAuth, controller.Validate)

	r.Run()

}
