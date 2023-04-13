package controllers

import (
	"errors"
	"fmt"
	"goelster/initializers"
	"goelster/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostsCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"1": "kate",
	})
}

// GetAllProduct
func GetAllProduct(DB *gorm.DB) ([]models.Sku, error) {

	var sku []models.Sku
	result := initializers.DB.Find(&sku)
	if result.Error != nil {
		return nil, result.Error

	}
	return sku, nil

}
func GetAllProductAPI(c *gin.Context) {
	var sku []models.Sku
	result := initializers.DB.Find(&sku)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, sku)
}

// GetFixPrice
func GetFixPrice(c *gin.Context) {
	var sku []models.Sku
	result := initializers.DB.Where("is_fix_price = ?", 1).Find(&sku)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, sku)
}

// func GetByID(c *gin.Context) {
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

func UpdateSku(c *gin.Context) {
	var sku models.Sku
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

func GetByFeild(c *gin.Context) {
	fieldName := c.Param("field_name")
	fieldValue := c.Param("field_value")

	var sku models.Sku
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
