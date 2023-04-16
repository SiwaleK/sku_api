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

func UpdateSkuBranch(c *gin.Context) {
	var sku models.SkuBranchPrice
	id1 := c.Param("sk_uid")
	id2 := c.Param("branch_id")

	var result *gorm.DB

	if id1 != "" && id2 != "" {
		result = initializers.DB.Where("sk_uid = ? AND branch_id = ?", id1, id2).First(&sku)
	} else {
		c.JSON(400, gin.H{"error": "Invalid request. Both sk_uid and branch_id must be provided"})
		return
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "Sku not found"})
		return
	} else if result.Error != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to query Sku: %v", result.Error)})
		return
	}

	if err := c.BindJSON(&sku); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := initializers.DB.Where("sk_uid = ? AND branch_id = ?", id1, id2).Updates(&sku).Error; err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to update Sku: %v", err)})
		return
	}

	c.JSON(200, sku)
}

func SkuBranchGetByFeild(c *gin.Context) {
	fieldName := c.Param("field_name")
	fieldValue := c.Param("field_value")

	var sku []models.SkuBranchPrice
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

func DeleteSkuBranch(c *gin.Context) {
	var sku models.SkuBranchPrice
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
