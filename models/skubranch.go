package models

import (
	_ "encoding/json"
	"time"
)

type SkuBranchPrice struct {
	SKUID      string    `json:"SKUID" gorm:primary_key`
	MerchantID string    `json:"MerchantID" gorm:primary_key`
	BranchID   string    `json:"BranchID" gorm:primary_key`
	Price      string    `json:"Price" gorm:default:null`
	StartDate  time.Time `json:"StartDate"`
	EndDate    time.Time `json:"EndDate" gorm:default:null`
	IsActiv    int16     `json:"IsActiv"`
}
