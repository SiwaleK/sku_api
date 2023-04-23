package controllers

import (
	"bytes"
	"encoding/json"
	"goelster/initializers"
	"goelster/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getSKUsBranch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getSKUs called"})
}
func TestGetSKUsBranchHTTP(t *testing.T) {
	req, err := http.NewRequest("GET", "/SkuBranch", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	r := gin.Default()

	r.GET("/SkuBranch", getSKUsBranch)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

}

func TestUpdateSKUsBranchHTTP(t *testing.T) {
	req, err := http.NewRequest("PUT", "/SkuBranch/:sk_uid", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	r := gin.Default()

	r.PUT("/SkuBranch/:sk_uid", getSKUsBranch)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

}

func TestGetByFeildSKUsBranch(t *testing.T) {
	// Create a mock gin context object
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set the necessary parameters in the context object
	c.Params = gin.Params{
		gin.Param{
			Key:   "field_name",
			Value: "merchant_id",
		},
		gin.Param{
			Key:   "field_value",
			Value: "000021",
		},
		gin.Param{
			Key:   "field_name",
			Value: "price",
		},
		gin.Param{
			Key:   "field_value",
			Value: "55.0000",
		},
	}

	// Create a mock gorm.DB object and initialize the data
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"SKUID", "MerchantID", "Price"}).
		AddRow("017b1", "000021", "55.0000")
	mock.ExpectQuery("^SELECT (.+) FROM skus WHERE SKUID = (.+)$").WithArgs("017b1").WillReturnRows(rows)

	// Set the mock DB object as the global DB instance
	initializers.DB, err = gorm.Open(postgres.Open("postgresql://root:secret@localhost:5435/sku_5435?sslmode=disable"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error opening mock database connection: %v", err)
	}

	// Call the function with the mock context object
	SkuBranchGetByFeild(c)

	// Verify the response
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	var skus []models.Sku
	err = json.Unmarshal(w.Body.Bytes(), &skus)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}
	if len(skus) != 1 || skus[0].SKUID != "017b1" {
		t.Errorf("expected sku with ProductName '017b1' but got: %+v", skus)
	}
}

func TestUpdateSkuBranch(t *testing.T) {
	// Create a mock gin context object
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set the necessary parameters in the context object
	c.Params = gin.Params{
		gin.Param{
			Key:   "sk_uid",
			Value: "02156ce9-2378-4570-b941-90509f2d69a4",
		},
		gin.Param{
			Key:   "branch_id",
			Value: "00002600000",
		},
	}

	// Define the update payload based on request parameters
	payload := gin.H{
		"Price": "100.0000",
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("error marshalling payload: %v", err)
	}

	// Create a mock gorm.DB object and set it as the global DB instance
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()
	initializers.DB, err = gorm.Open(postgres.Open("postgresql://root:secret@localhost:5435/sku_5435?sslmode=disable"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error opening mock database connection: %v", err)
	}

	// Mock the database operation
	mock.ExpectExec("^UPDATE sku_branch_prices SET (.+) WHERE sk_uid = \\? AND branch_id = \\?$").
		WithArgs("02156ce9-2378-4570-b941-90509f2d69a4", "00002600000", "100.0000").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function with the mock context object
	c.Request = httptest.NewRequest(http.MethodPut, "/sku_branch_prices/02156ce9-2378-4570-b941-90509f2d69a4/00002600000", bytes.NewBuffer(jsonPayload))
	UpdateSkuBranch(c)

	// // Verify the response
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	var sku models.SkuBranchPrice
	err = json.Unmarshal(w.Body.Bytes(), &sku)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}
}
