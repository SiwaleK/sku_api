package main

import (
	"goelster/controllers"
	"goelster/initializers"
	_ "goelster/models"

	"github.com/gin-gonic/gin"
)

func init() { // if this work correctly this would run env file and work on port 3000
	initializers.LoadEnvVariables()
	initializers.Connect()
}
func main() {
	r := gin.Default()
	// Sku
	r.GET("/", controllers.PostsCreate)
	r.GET("/Sku", controllers.GetAllProductAPI)
	//r.GET("/Sku/:sk_uid", controllers.GetByID)
	r.GET("/SkuGetFixPrice", controllers.GetFixPrice)
	r.GET("/Sku/:field_name/:field_value", controllers.GetByFeild)
	r.PUT("/Sku/:sk_uid", controllers.UpdateSku)

	//SkuBranch
	r.GET("/SkuBranch", controllers.SkuBranchGetAllProduct)
	r.GET("/SkuBranch/:field_name/:field_value", controllers.SkuBranchGetByFeild)
	r.PUT("/SkuBranch/:sk_uid", controllers.UpdateSkuBranch)

	r.Run()
}
