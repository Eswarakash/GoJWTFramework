package migrate

import (
	"githiub.com/Eswarakash/GoJWTFramework/initializer"
	"githiub.com/Eswarakash/GoJWTFramework/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectDB()

}

func MigrateAuto() {
	initializer.DB.AutoMigrate(&models.User{})
}
