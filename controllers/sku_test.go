package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goelster/initializers"
	"goelster/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getSKUs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getSKUs called"})
}

func TestGetSKUsHTTP(t *testing.T) {
	req, err := http.NewRequest("GET", "/Sku", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	r := gin.Default()

	r.GET("/Sku", getSKUsBranch)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestUpdateSKUsHTTP(t *testing.T) {
	req, err := http.NewRequest("PUT", "/Sku/:sk_uid", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	r := gin.Default()

	r.PUT("/Sku/:sk_uid", getSKUs)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

}

func TestGetByFeild(t *testing.T) {
	// Create a mock gin context object
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create a mock gorm.DB object and initialize the data
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()
	initializers.DB, err = gorm.Open(postgres.Open("postgresql://root:secret@localhost:5435/sku_5435?sslmode=disable"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error opening mock database connection: %v", err)
	}

	// Mock the SELECT operation
	rows := sqlmock.NewRows([]string{"sku_id", "barcode_pos", "product_name", "brand_id", "product_group_id", "product_cat_id", "product_sub_cat_id", "product_size_id", "product_unit", "pack_size", "unit", "ban_for_pracharat", "is_vat", "create_by", "create_date", "is_active", "merchant_id", "map_sku", "is_fix_price"}).
		AddRow("ef2b5a9b-a5a3-4db7-919e-4e35bc75c123", "OatInWZa", "โอ๊ตเทพซ่า", 1552, 1, 1, 0, 99, 1, "0", 0, 1, 1, "dev1", "2023-01-04 07:01:59", 1, nil, nil, 0).
		AddRow("f12e558d-9308-49d0-ab96-bac908806795", "234213652341", "เกก", 2328, 8, 62, 213, 8, 6, "5", 0, 1, 1, "dev1", "2022-09-15 11:55:35", 1, nil, nil, 0).
		AddRow("f18d2d44-c07d-4d76-95ce-5ede753794d9", "456892347658922", "test1", 5, 5, 30, 87, 7, 8, "5", 0, 1, 1, "dev1", "2022-09-15 11:19:50", 1, nil, nil, 0).
		AddRow("f7dd9cd8-789f-4530-9dbb-214942c7718e", "Cooller", "LIQUID COOLING COOLER MASTER MASTERLIQUID ML240L V.2 WHITE ARGB (MLW-D24M-A18PW-R)", 1338, 1, 1, 0, 99, 1, '0', 0, 1, 1, "songsit3", "2023-03-01 10:57:10", 1, "1000000019", nil, 0).
		AddRow("f96c476e-8143-4360-aa06-971d12ca7761", "bbb", "bbb", 1338, 1, 1, 0, 99, 1, "0", 0, 1, 1, "dev1", "2022-10-26 13:30:57", 1, "1000000000", nil, 0).
		AddRow("fa7c1b6e-18f7-4041-a0b7-f807611e9c00", "testPrice1", "testPost1_dev", 912, 1, 1, 0, 8, 1, "13.0", 0, 1, 0, "dev1", "2023-01-12 12:08:14", 1, "1000000000", nil, 0).
		AddRow("faceef73-4e12-4b3c-8c9a-4010bf83fae9", "prodUnitTest1", "prodUnitTest1", 85, 1, 1, 0, 1, 1, "4.0", 0, 1, 1, "dev1", "2023-01-12 12:07:39", 1, "1000000000", nil, 0).
		AddRow("fad4fe20-dc13-4712-ae35-3d9bfe26c14c", "testBlueFlag", "testBlueFlage", 1, 1, 1, 0, 1, 1, "12.0", 0, 1, 1, "dev1", "2023-01-12 12:06:56", 1, "1000000000", nil, 0).
		AddRow("fd16b16b-1f92-4494-a4b4-f0c42f6d6ea7", "355678", "ดดกปหกกด", 1338, 1, 1, 0, 99, 1, "0", 0, 1, 1, "tccinven3", "2023-02-03 15:26:53", 1, "000026", nil, 0).
		AddRow("fd5d5aeb-6722-4edb-9b85-f1bc5e81fde5", "123", "123", 1338, 1, 1, 0, 99, 1, "0", 0, 1, 1, "tccinven2", "2023-03-13 09:42:54", 1, "000025", nil, 0)

	mock.ExpectQuery("^SELECT (.+) FROM skus WHERE product_name = \\? ").
		WithArgs("โอ๊ตเทพซ่า").
		WillReturnRows(rows)

	// Call the function with the mock context object
	GetByFeild(c)

	// Verify the response
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	var skus []models.Sku
	err = json.Unmarshal(w.Body.Bytes(), &skus)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}
	if len(skus) != 1 || skus[0].ProductName != "โอ๊ตเทพซ่า" {
		t.Errorf("expected sku with ProductName 'โอ๊ตเทพซ่า' but got: %+v", skus)
	}
}

func TestUpdateSku(t *testing.T) {
	// Create a mock gin context object
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set the necessary parameters in the context object
	c.Params = gin.Params{
		gin.Param{
			Key:   "sk_uid",
			Value: "f12e558d-9308-49d0-ab96-bac908806795",
		},
	}
	// Define the update payload based on request parameters
	payload := gin.H{
		"BrandID": 1552,
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

	// Define the named parameters based on the payload keys
	namedParams := make([]string, 0, len(payload))
	for key := range payload {
		namedParams = append(namedParams, fmt.Sprintf("%s=:%s", key, key))
	}

	// Mock the update operation with the named parameters

	mock.ExpectExec("^UPDATE skus SET (.+) WHERE sk_uid = \\?$").
		WithArgs("f12e558d-9308-49d0-ab96-bac908806795").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function with the mock context object
	c.Request = httptest.NewRequest(http.MethodPut, "/Sku/f12e558d-9308-49d0-ab96-bac908806795", bytes.NewBuffer(jsonPayload))
	UpdateSku(c)

	// Verify the response
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	var sku models.Sku
	err = json.Unmarshal(w.Body.Bytes(), &sku)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}
	var respPayload gin.H
	err = json.Unmarshal(w.Body.Bytes(), &respPayload)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}

}

func TestGetAllProduct(t *testing.T) {
	// Create a mock gin context object
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

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

	// Define the expected rows
	rows := sqlmock.NewRows([]string{"sku_id", "barcode_pos", "product_name", "brand_id", "product_group_id", "product_cat_id", "product_sub_cat_id", "product_size_id", "product_unit", "pack_size", "unit", "ban_for_pracharat", "is_vat", "create_by", "create_date", "is_active", "merchant_id", "map_sku", "is_fix_price"}).
		AddRow("ef2b5a9b-a5a3-4db7-919e-4e35bc75c123", "OatInWZa", "โอ๊ตเทพซ่า", 1552, 1, 1, 0, 99, 1, "0", 0, 1, 1, "dev1", "2023-01-04 07:01:59", 1, nil, nil, 0).
		AddRow("f12e558d-9308-49d0-ab96-bac908806795", "234213652341", "เกก", 2328, 8, 62, 213, 8, 6, "5", 0, 1, 1, "dev1", "2022-09-15 11:55:35", 1, nil, nil, 0).
		AddRow("f18d2d44-c07d-4d76-95ce-5ede753794d9", "456892347658922", "test1", 5, 5, 30, 87, 7, 8, "5", 0, 1, 1, "dev1", "2022-09-15 11:19:50", 1, nil, nil, 0).
		AddRow("f7dd9cd8-789f-4530-9dbb-214942c7718e", "Cooller", "LIQUID COOLING COOLER MASTER MASTERLIQUID ML240L V.2 WHITE ARGB (MLW-D24M-A18PW-R)", 1338, 1, 1, 0, 99, 1, '0', 0, 1, 1, "songsit3", "2023-03-01 10:57:10", 1, "1000000019", nil, 0).
		AddRow("f96c476e-8143-4360-aa06-971d12ca7761", "bbb", "bbb", 1338, 1, 1, 0, 99, 1, "0", 0, 1, 1, "dev1", "2022-10-26 13:30:57", 1, "1000000000", nil, 0).
		AddRow("fa7c1b6e-18f7-4041-a0b7-f807611e9c00", "testPrice1", "testPost1_dev", 912, 1, 1, 0, 8, 1, "13.0", 0, 1, 0, "dev1", "2023-01-12 12:08:14", 1, "1000000000", nil, 0).
		AddRow("faceef73-4e12-4b3c-8c9a-4010bf83fae9", "prodUnitTest1", "prodUnitTest1", 85, 1, 1, 0, 1, 1, "4.0", 0, 1, 1, "dev1", "2023-01-12 12:07:39", 1, "1000000000", nil, 0).
		AddRow("fad4fe20-dc13-4712-ae35-3d9bfe26c14c", "testBlueFlag", "testBlueFlage", 1, 1, 1, 0, 1, 1, "12.0", 0, 1, 1, "dev1", "2023-01-12 12:06:56", 1, "1000000000", nil, 0).
		AddRow("fd16b16b-1f92-4494-a4b4-f0c42f6d6ea7", "355678", "ดดกปหกกด", 1338, 1, 1, 0, 99, 1, "0", 0, 1, 1, "tccinven3", "2023-02-03 15:26:53", 1, "000026", nil, 0).
		AddRow("fd5d5aeb-6722-4edb-9b85-f1bc5e81fde5", "123", "123", 1338, 1, 1, 0, 99, 1, "0", 0, 1, 1, "tccinven2", "2023-03-13 09:42:54", 1, "000025", nil, 0)

	// Mock the SELECT operation
	mock.ExpectQuery("^SELECT (.+) FROM skus$").WillReturnRows(rows)

	// Call the function with the mock context object
	GetAllProduct(c)

	// Verify the response
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	var sku []models.Sku
	err = json.Unmarshal(w.Body.Bytes(), &sku)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}
	if len(sku) != 8 {
		t.Errorf("expected 2 SKUs but got %d", len(sku))
	}
}

func TestGetFieldProduct2(t *testing.T) {
	// Create a mock gin context object
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

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

	// Define the expected rows
	// rows := sqlmock.NewRows([]string{"sku_id", "barcode_pos", "product_name", "brand_id", "product_group_id", "product_cat_id", "product_sub_cat_id", "product_size_id", "product_unit", "pack_size", "unit", "ban_for_pracharat", "is_vat", "create_by", "create_date", "is_active", "merchant_id", "map_sku", "is_fix_price"}).
	// 	AddRow("ef2b5a9b-a5a3-4db7-919e-4e35bc75c123", "OatInWZa", "โอ๊ตเทพซ่า", 1552, 1, 1, 0, 99, 1, "0", 0, 1, 1, "dev1", "2023-01-04 07:01:59", 1, nil, nil, 0).
	// 	AddRow("f12e558d-9308-49d0-ab96-bac908806795", "234213652341", "เกก", 2328, 8, 62, 213, 8, 6, "5", 0, 1, 1, "dev1", "2022-09-15 11:55:35", 1, nil, nil, 0).
	// 	AddRow("f18d2d44-c07d-4d76-95ce-5ede753794d9", "456892347658922", "test1", 5, 5, 30, 87, 7, 8, "5", 0, 1, 1, "dev1", "2022-09-15 11:19:50", 1, nil, nil, 0).
	// 	AddRow("f7dd9cd8-789f-4530-9dbb-214942c7718e", "Cooller", "LIQUID COOLING COOLER MASTER MASTERLIQUID ML240L V.2 WHITE ARGB (MLW-D24M-A18PW-R)", 1338, 1, 1, 0, 99, 1, '0', 0, 1, 1, "songsit3", "2023-03-01 10:57:10", 1, "1000000019", nil, 0).
	// 	AddRow("f96c476e-8143-4360-aa06-971d12ca7761", "bbb", "bbb", 1338, 1, 1, 0, 99, 1, "0", 0, 1, 1, "dev1", "2022-10-26 13:30:57", 1, "1000000000", nil, 0).
	// 	AddRow("fa7c1b6e-18f7-4041-a0b7-f807611e9c00", "testPrice1", "testPost1_dev", 912, 1, 1, 0, 8, 1, "13.0", 0, 1, 0, "dev1", "2023-01-12 12:08:14", 1, "1000000000", nil, 0).
	// 	AddRow("faceef73-4e12-4b3c-8c9a-4010bf83fae9", "prodUnitTest1", "prodUnitTest1", 85, 1, 1, 0, 1, 1, "4.0", 0, 1, 1, "dev1", "2023-01-12 12:07:39", 1, "1000000000", nil, 0).
	// 	AddRow("fad4fe20-dc13-4712-ae35-3d9bfe26c14c", "testBlueFlag", "testBlueFlage", 1, 1, 1, 0, 1, 1, "12.0", 0, 1, 1, "dev1", "2023-01-12 12:06:56", 1, "1000000000", nil, 0).
	// 	AddRow("fd16b16b-1f92-4494-a4b4-f0c42f6d6ea7", "355678", "ดดกปหกกด", 1338, 1, 1, 0, 99, 1, "0", 0, 1, 1, "tccinven3", "2023-02-03 15:26:53", 1, "000026", nil, 0).
	// 	AddRow("fd5d5aeb-6722-4edb-9b85-f1bc5e81fde5", "123", "123", 1338, 1, 1, 0, 99, 1, "0", 0, 1, 1, "tccinven2", "2023-03-13 09:42:54", 1, "000025", nil, 0)

	// Mock the SELECT operation
	mock.ExpectExec("^SELECT skus SET (.+) WHERE sk_uid = \\?$").
		WithArgs("f12e558d-9308-49d0-ab96-bac908806795").
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Call the function with the mock context object
	GetByFeild(c)

	// Verify the response
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	var sku []models.Sku
	err = json.Unmarshal(w.Body.Bytes(), &sku)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}
	if len(sku) != 0 {
		t.Errorf("expected 1 SKUs but got %d", len(sku))
	}
}

func TestDeleteSku(t *testing.T) {
	// Create a mock gin context object
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set the necessary parameters in the context object
	c.Params = gin.Params{
		gin.Param{
			Key:   "sk_uid",
			Value: "ef2b5a9b-a5a3-4db7-919e-4e35bc75c123",
		},
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

	// Mock the select operation

	mock.ExpectQuery("^SELECT (.+) FROM skus WHERE sk_uid = ?$").WithArgs("ef2b5a9b-a5a3-4db7-919e-4e35bc75c123")

	// Mock the delete operation
	mock.ExpectExec("^DELETE FROM skus WHERE sk_uid = ?$").WithArgs("ef2b5a9b-a5a3-4db7-919e-4e35bc75c123").WillReturnResult(sqlmock.NewResult(0, 1))

	// Call the function with the mock context object
	DeleteSku(c)

	// Verify the response
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}
	var response gin.H
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}
	expected := gin.H{"message": "Sku deleted successfully"}
	if !reflect.DeepEqual(response, expected) {
		t.Errorf("expected response %+v but got %+v", expected, response)
	}
}

func TestUpdateSkuHTTP(t *testing.T) {
	// Create a new SKU to be updated
	sku := models.Sku{
		ProductName: "Updated Product Name",
	}

	// Marshal the SKU to JSON
	jsonPayload, err := json.Marshal(sku)
	if err != nil {
		t.Fatalf("error marshalling SKU: %v", err)
	}

	// Create a new HTTP request with the JSON payload
	req, err := http.NewRequest("PUT", "/Sku/f12e558d-9308-49d0-ab96-bac908806795", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the HTTP response
	recorder := httptest.NewRecorder()

	// Create a new Gin router and register the handler function
	r := gin.Default()
	r.PUT("/Sku/:sk_uid", UpdateSku)

	// Serve the HTTP request using the Gin router
	r.ServeHTTP(recorder, req)

	// Check the HTTP response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Unmarshal the response body into a SKU struct
	var updatedSku models.Sku
	err = json.Unmarshal(recorder.Body.Bytes(), &updatedSku)
	if err != nil {
		t.Fatalf("error unmarshalling response: %v", err)
	}

	// Check that the SKU was updated correctly
	assert.Equal(t, sku.ProductName, updatedSku.ProductName)
}
