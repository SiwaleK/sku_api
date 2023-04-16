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

//GetAllProduct

func GetAllProduct(c *gin.Context) {
	var sku []models.Sku
	result := initializers.DB.Find(&sku)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, sku)
}

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

	var skus []models.Sku
	query := fmt.Sprintf("%s = ?", fieldName)
	result := initializers.DB.Where(query, fieldValue).Find(&skus)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": fieldValue})
			return
		} else {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

	}
	c.JSON(200, skus)
}

func DeleteSku(c *gin.Context) {
	var sku models.Sku
	id := c.Param("sk_uid")
	if err := initializers.DB.Where("sk_uid = ?", id).First(&sku).Error; err != nil {
		c.JSON(404, gin.H{"error": "Sku not found"})
		return
	}
	if err := initializers.DB.Delete(&sku).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete Sku"})
		return
	}
	c.JSON(200, gin.H{"message": "Sku deleted successfully"})
}
