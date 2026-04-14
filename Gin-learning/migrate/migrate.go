package main

import (
	"gin-learning/initializers"
	"gin-learning/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectionToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
