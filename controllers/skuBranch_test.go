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

// func SkuBranchTestGetAllProduct(t *testing.T) {
// 	// Create a mock gin context object
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	// Create a mock gorm.DB object and set it as the global DB instance
// 	mockDB, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("error creating mock database: %v", err)
// 	}
// 	defer mockDB.Close()
// 	initializers.DB, err = gorm.Open(postgres.Open("postgresql://root:secret@localhost:5435/sku_5435?sslmode=disable"), &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("error opening mock database connection: %v", err)
// 	}

// 	// Define the expected rows
// 	rows := sqlmock.NewRows([]string{"sku_id", "merchant_id", "branch_id", "price", "start_date", "end_date", "is_active"}).
// 		AddRow("017b1", "000021", "00002100000", 55.0000, "2023-02-20 09:30:52", nil, 1).
// 		AddRow("017b1", "000022", "00002200000", 100.0000, "2023-02-03 10:43:15", nil, 1).
// 		AddRow("017b1", "000024", "00002400000", 30.0000, "2023-01-30 03:51:07", nil, 1).
// 		AddRow("017b1", "1000000000", "1000000000000", 1024.0000, "2022-10-24 07:23:12", nil, 1).
// 		// ('017b1c0a-df2e-4f24-be7c-4ae9cadd24ef','1000000000','1000000000000',9999.0000,'2022-08-29 09:28:01',nil,1),
// 		// ('017b1c0a-df2e-4f24-be7c-4ae9cadd24ef','1000000000','1000000000001',299.0000,'2022-11-03 11:10:52',nil,1),
// 		// ('01e22ab6-84fe-4ebb-9d9e-b571ac9a6b31','000024','00002400000',159.0000,'2023-03-13 05:10:59',nil,1),
// 		// ('02156ce9-2378-4570-b941-90509f2d69a4','000026','00002600000',100.0000,'2023-03-08 07:28:25',nil,1),
// 		// ('039e2b17-b557-48ad-8ade-a51495bb4be6','1000000000','1000000000000',11.0000,'2023-01-20 08:04:32',nil,1),
// 		// ('0517bf6a-ffe1-4819-9c17-09155cdf8bb5','1000000000','1000000000000',0.0000,'2022-08-31 07:33:33',nil,1);

// 		// Mock the SELECT operation
// 		mock.ExpectQuery("^SELECT (.+) FROM skus$").WillReturnRows(rows)

// 	// Call the function with the mock context object
// 	SkuBranchGetAllProduct(c)

// 	// Verify the response
// 	if w.Code != http.StatusOK {
// 		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
// 	}
// 	var sku []models.Sku
// 	err = json.Unmarshal(w.Body.Bytes(), &sku)
// 	if err != nil {
// 		t.Errorf("error unmarshalling response: %v", err)
// 	}
// 	if len(sku) != 8 {
// 		t.Errorf("expected 2 SKUs but got %d", len(sku))
// 	}
// }
