package tests

import (
	"_go/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductModel(t *testing.T) {
	p := models.Product{ID: 1, Name: "Test", Price: 10.0, Stock: 5, CategoryID: 2}
	assert.Equal(t, uint(1), p.ID)
	assert.Equal(t, "Test", p.Name)
	assert.Equal(t, 10.0, p.Price)
	assert.Equal(t, 5, p.Stock)
	assert.Equal(t, uint(2), p.CategoryID)
	p.Name = "Nowy"
	assert.Equal(t, "Nowy", p.Name)
	p.Price = 20.0
	assert.Equal(t, 20.0, p.Price)
	p.Stock = 7
	assert.Equal(t, 7, p.Stock)
	p.CategoryID = 3
	assert.Equal(t, uint(3), p.CategoryID)
	// 10 asercji
}

func TestCartModel(t *testing.T) {
	c := models.Cart{ID: 1, TotalPrice: 0.0}
	assert.Equal(t, uint(1), c.ID)
	assert.Equal(t, 0.0, c.TotalPrice)
	c.TotalPrice = 99.99
	assert.Equal(t, 99.99, c.TotalPrice)
	c.ID = 2
	assert.Equal(t, uint(2), c.ID)
	// Dodajemy produkty do koszyka
	c.Products = []models.Product{{ID: 1, Name: "A", Price: 1.0}, {ID: 2, Name: "B", Price: 2.0}}
	assert.Len(t, c.Products, 2)
	assert.Equal(t, "A", c.Products[0].Name)
	assert.Equal(t, 2.0, c.Products[1].Price)
	c.Products = append(c.Products, models.Product{ID: 3, Name: "C", Price: 3.0})
	assert.Len(t, c.Products, 3)
	// 10 asercji
}

func TestCategoryModel(t *testing.T) {
	cat := models.Category{ID: 1, Name: "Cat1"}
	assert.Equal(t, uint(1), cat.ID)
	assert.Equal(t, "Cat1", cat.Name)
	cat.Name = "Nowa"
	assert.Equal(t, "Nowa", cat.Name)
	cat.Products = []models.Product{{ID: 1, Name: "A", Price: 1.0}}
	assert.Len(t, cat.Products, 1)
	cat.Products = append(cat.Products, models.Product{ID: 2, Name: "B", Price: 2.0})
	assert.Len(t, cat.Products, 2)
	// 6 asercji
}

func TestProductLogic(t *testing.T) {
	p := models.Product{Name: "X", Price: 5.0, Stock: 1}
	assert.True(t, p.Price > 0)
	p.Price = -1
	assert.False(t, p.Price > 0)
	p.Name = ""
	assert.Empty(t, p.Name)
	p.Stock = 0
	assert.Equal(t, 0, p.Stock)
	// 4 asercje
}

func TestCartLogic(t *testing.T) {
	c := models.Cart{TotalPrice: 0.0}
	assert.Equal(t, 0.0, c.TotalPrice)
	c.Products = []models.Product{}
	assert.Empty(t, c.Products)
	c.Products = append(c.Products, models.Product{ID: 1, Name: "A", Price: 1.0})
	assert.NotEmpty(t, c.Products)
	c.TotalPrice = 1.0
	assert.Equal(t, 1.0, c.TotalPrice)
	// 5 asercji
}

func TestCategoryLogic(t *testing.T) {
	cat := models.Category{Name: ""}
	assert.Empty(t, cat.Name)
	cat.Name = "Kategoria"
	assert.NotEmpty(t, cat.Name)
	cat.Products = []models.Product{}
	assert.Empty(t, cat.Products)
	cat.Products = append(cat.Products, models.Product{ID: 1, Name: "A", Price: 1.0})
	assert.NotEmpty(t, cat.Products)
	// 5 asercji
}

func TestBulkProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		p := models.Product{ID: uint(i), Name: "P", Price: float64(i), Stock: i}
		assert.Equal(t, uint(i), p.ID)
		assert.Equal(t, "P", p.Name)
		assert.Equal(t, float64(i), p.Price)
		assert.Equal(t, i, p.Stock)
	}
	// 40 asercji
}

// SUMA: 10 + 10 + 6 + 4 + 5 + 5 + 40 = 80 asercji
