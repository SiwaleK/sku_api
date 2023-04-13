package models

import (
	_ "encoding/json"
	"time"

	"gorm.io/gorm"
)

type Sku struct {
	gorm.Model
	SKUID           string    `json:"SKUID" gorm:primary_key`
	BarcodePOS      string    `json:"BarcodePOS" gorm:default:null`
	ProductName     string    `json:"ProductName" gorm:default:null`
	BrandID         int32     `json:"BrandID" gorm:default:null`
	ProductGroupID  int32     `json:"ProductGroupID" gorm:default:null`
	ProductCatID    int32     `json:"ProductCatID" gorm:default:null`
	ProductSubCatID int32     `json:"ProductSubCatID" gorm:default:null`
	ProductSizeID   int32     `json:"ProductSizeID" gorm:default:null`
	ProductUnit     int32     `json:"ProductUnit" gorm:default:null`
	PackSize        string    `json:"PackSize" gorm:default:null`
	Unit            int32     `json:"Unit" gorm:default:null`
	BanForPracharat int32     `json:"BanForPracharat" gorm:default:null`
	IsVat           int16     `json:"IsVat" gorm:default:null`
	CreateBy        string    `json:"CreateBy" gorm:default:null`
	CreateDate      time.Time `json:"CreateDate"`
	IsActive        int16     `json:"IsActive"`
	MerchantID      string    `json:"MerchantID" gorm:default:null`
	MapSKU          string    `json:"MapSKU" gorm:default:null`
	IsFixPrice      int16     `json:"IsFixPrice"`
}
