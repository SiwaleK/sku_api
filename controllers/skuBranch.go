package controllers

import (
	"errors"
	"fmt"
	"goelster/initializers"
	"goelster/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllProduct

func SkuBranchGetAllProduct(c *gin.Context) {
	var skubranch []models.SkuBranchPrice
	result := initializers.DB.Find(&skubranch)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, skubranch)
}

// func SkuBranchID(c *gin.Context) {
// 	id := c.Param("sk_uid")
// 	var sku models.Sku
// 	result := initializers.DB.Where("sk_uid = ?", id).Find(&sku)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			c.JSON(404, gin.H{"error": id})
// 			return
// 		} else {
// 			c.JSON(500, gin.H{"error": result.Error.Error()})
// 			return
// 		}
// 	}
// 	c.JSON(200, sku)
// }

func UpdateSkuBranch(c *gin.Context) {
	var sku models.SkuBranchPrice
	id := c.Param("sk_uid")
	if err := initializers.DB.Where("sk_uid = ?", id).First(&sku).Error; err != nil {
		c.JSON(404, gin.H{"error": "Sku not found"})
		return
	}
	if err := c.BindJSON(&sku); err != nil {
		c.JSON(500, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := initializers.DB.Save(&sku).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update Sku"})
		return
	}
	c.JSON(200, sku)
}

func SkuBranchGetByFeild(c *gin.Context) {
	fieldName := c.Param("field_name")
	fieldValue := c.Param("field_value")

	var sku models.SkuBranchPrice
	query := fmt.Sprintf("%s = ?", fieldName)
	result := initializers.DB.Where(query, fieldValue).Find(&sku)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": fieldValue})
			return
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

	}
	c.JSON(200, sku)
}
