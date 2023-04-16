package main

import (
	"goelster/controllers"
	"goelster/initializers"
	_ "goelster/models"

	"github.com/gin-gonic/gin"
	_ "github.com/golang/mock/gomock"
)

func init() { // if this work correctly this would run env file and work on port 3000
	initializers.LoadEnvVariables()
	initializers.Connect()
}
func main() {
	r := gin.Default()

	// Sku
	r.GET("/", controllers.PostsCreate)
	r.GET("/Sku", controllers.GetAllProduct)
	//r.GET("/Sku",controllers.getAllProductHandler)
	r.GET("/Sku/:field_name/:field_value", controllers.GetByFeild)
	r.PUT("/Sku/:sk_uid", controllers.UpdateSku)
	r.DELETE("/Sku/:sk_uid", controllers.DeleteSku)

	//SkuBranch
	r.GET("/SkuBranch", controllers.SkuBranchGetAllProduct)
	r.GET("/SkuBranch/:field_name/:field_value", controllers.SkuBranchGetByFeild)

	r.PUT("/SkuBranch/:sk_uid", controllers.UpdateSkuBranch)
	r.PUT("/SkuBranch/:sk_uid/:branch_id", controllers.UpdateSkuBranch)

	r.DELETE("/SkuBranch/:sk_uid", controllers.DeleteSkuBranch)

	r.Run()
}
