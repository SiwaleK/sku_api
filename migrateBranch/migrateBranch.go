package main

import (
	"goelster/initializers"
	"goelster/models"
)

func init() {
	initializers.Connect()
	initializers.LoadEnvVariables()
}

func main() {

	initializers.DB.AutoMigrate(&models.SkuBranchPrice{})

}
