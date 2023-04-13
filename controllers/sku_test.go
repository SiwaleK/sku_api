package controllers

import (
	"encoding/json"
	"fmt"
	"goelster/initializers"
	"goelster/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProduct(t *testing.T) {
	// Initialize a new Gin router
	r := gin.Default()

	// Register the GetAllProductAPI handler
	r.GET("/GetAll", GetAllProductAPI)

	// Make a request to the GetAllProductAPI endpoint
	req, _ := http.NewRequest("GET", "/GetAll", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check that the response has a status code of 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Print the response body
	fmt.Println(string(w.Body.Bytes()))

	// Decode the response body into a []models.Sku
	var skus []models.Sku
	err := json.Unmarshal(w.Body.Bytes(), &skus)
	assert.NoError(t, err)

	// Query the database to get all skus
	var expectedSkus []models.Sku
	err = initializers.DB.Find(&expectedSkus).Error
	assert.NoError(t, err)

	// Compare the expected skus to the actual skus returned in the response
	assert.Equal(t, expectedSkus, skus)
}

func getSKUs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getSKUs called"})
}
func getSKUsBranch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getSKUs called"})
}

func TestGetSKUs(t *testing.T) {
	req, err := http.NewRequest("GET", "/GetAll", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	r := gin.Default()

	r.GET("/GetAll", getSKUsBranch)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	//assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestGetFixPrice(t *testing.T) {
	req, err := http.NewRequest("GET", "/GetFixPrice", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	r := gin.Default()

	r.GET("/GetAll", getSKUs)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	//assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestGetSKUsBranch(t *testing.T) {
	req, err := http.NewRequest("GET", "/SkuBranchGetAll", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	r := gin.Default()

	r.GET("/SkuBranchGetAll", getSKUsBranch)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	//assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}
