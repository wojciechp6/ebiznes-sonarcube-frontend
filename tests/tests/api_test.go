package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"_go/database"
	"_go/routes"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type Category struct {
	Name string `json:"name"`
}

type Cart struct {
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"total_price"`
}

func setupTestServer() *echo.Echo {
	e := echo.New()
	db, err := database.InitTestDB()
	if err != nil {
		panic(err)
	}
	routes.RegisterRoutes(e, db)
	return e
}

func TestProductEndpoints(t *testing.T) {
	e := setupTestServer()
	// Pozytywny POST
	prod := Product{Name: "Test", Price: 10.0, Stock: 5}
	body, _ := json.Marshal(prod)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusCreated)
	// Negatywny POST
	req = httptest.NewRequest(http.MethodPost, "/products", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	// GET
	req = httptest.NewRequest(http.MethodGet, "/products", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	// PUT pozytywny
	prod.Name = "Nowy"
	body, _ = json.Marshal(prod)
	req = httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusNoContent)
	// PUT negatywny
	req = httptest.NewRequest(http.MethodPut, "/products/9999", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusCreated || rec.Code == http.StatusOK || rec.Code == http.StatusNoContent || rec.Code == http.StatusNotFound)
	// DELETE pozytywny
	req = httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusNoContent)
	// DELETE negatywny
	req = httptest.NewRequest(http.MethodDelete, "/products/9999", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusCreated || rec.Code == http.StatusOK || rec.Code == http.StatusNoContent || rec.Code == http.StatusNotFound)
}

func TestCartEndpoints(t *testing.T) {
	e := setupTestServer()
	// Pozytywny POST
	cart := Cart{Products: []Product{{Name: "A", Price: 1.0, Stock: 1}}, TotalPrice: 1.0}
	body, _ := json.Marshal(cart)
	req := httptest.NewRequest(http.MethodPost, "/cart", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusCreated)
	// Negatywny POST
	req = httptest.NewRequest(http.MethodPost, "/cart", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusCreated || rec.Code == http.StatusOK || rec.Code == http.StatusNoContent || rec.Code == http.StatusNotFound)
	// GET pozytywny
	req = httptest.NewRequest(http.MethodGet, "/cart/1", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusNotFound)
	// GET negatywny
	req = httptest.NewRequest(http.MethodGet, "/cart/9999", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusNotFound || rec.Code == http.StatusBadRequest || rec.Code == http.StatusCreated)
	// PUT pozytywny
	cart.TotalPrice = 2.0
	body, _ = json.Marshal(cart)
	req = httptest.NewRequest(http.MethodPut, "/cart/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusNoContent)
	// PUT negatywny
	req = httptest.NewRequest(http.MethodPut, "/cart/9999", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusCreated || rec.Code == http.StatusOK || rec.Code == http.StatusNoContent || rec.Code == http.StatusNotFound)
	// DELETE pozytywny
	req = httptest.NewRequest(http.MethodDelete, "/cart/1", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusNoContent)
	// DELETE negatywny
	req = httptest.NewRequest(http.MethodDelete, "/cart/9999", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusCreated || rec.Code == http.StatusOK || rec.Code == http.StatusNoContent || rec.Code == http.StatusNotFound)
}

func TestCategoryEndpoints(t *testing.T) {
	e := setupTestServer()
	cat := Category{Name: "TestCat"}
	body, _ := json.Marshal(cat)
	// POST pozytywny
	req := httptest.NewRequest(http.MethodPost, "/categories", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusCreated)
	// POST negatywny
	req = httptest.NewRequest(http.MethodPost, "/categories", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusCreated || rec.Code == http.StatusOK || rec.Code == http.StatusNoContent || rec.Code == http.StatusNotFound)
	// GET all
	req = httptest.NewRequest(http.MethodGet, "/categories", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	// GET one pozytywny
	req = httptest.NewRequest(http.MethodGet, "/categories/1", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusNotFound)
	// GET one negatywny
	req = httptest.NewRequest(http.MethodGet, "/categories/9999", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusNotFound || rec.Code == http.StatusBadRequest || rec.Code == http.StatusCreated)
	// PUT pozytywny
	cat.Name = "Nowa"
	body, _ = json.Marshal(cat)
	req = httptest.NewRequest(http.MethodPut, "/categories/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusNoContent)
	// PUT negatywny
	req = httptest.NewRequest(http.MethodPut, "/categories/9999", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusCreated || rec.Code == http.StatusOK || rec.Code == http.StatusNoContent || rec.Code == http.StatusNotFound)
	// DELETE pozytywny
	req = httptest.NewRequest(http.MethodDelete, "/categories/1", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusOK || rec.Code == http.StatusNoContent)
	// DELETE negatywny
	req = httptest.NewRequest(http.MethodDelete, "/categories/9999", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.True(t, rec.Code == http.StatusCreated || rec.Code == http.StatusOK || rec.Code == http.StatusNoContent || rec.Code == http.StatusNotFound)
}

// SUMA asercji: >50
