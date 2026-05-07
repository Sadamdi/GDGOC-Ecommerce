package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ecommerce-backend/internal/domain"
	repoMongo "ecommerce-backend/internal/repository/mongo"
)

func TestProductRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db, teardown := repoMongo.SetupTestMongoDB(t)
	defer teardown()

	repo := repoMongo.NewProductRepository(db)
	ctx := context.Background()

	// Seed Data
	products := []*domain.Product{
		{Name: "Baju Hitam", Price: 50000, Stock: 10, CategoryID: "cat1"},
		{Name: "Baju Putih", Price: 40000, Stock: 0, CategoryID: "cat1"},
		{Name: "Celana Hitam", Price: 100000, Stock: 5, CategoryID: "cat2"},
		{Name: "Topi Biru", Price: 20000, Stock: 20, CategoryID: "cat3"},
	}

	for _, p := range products {
		err := repo.Create(ctx, p)
		require.NoError(t, err)
	}

	t.Run("Get All Products - Filter Category", func(t *testing.T) {
		res, total, err := repo.GetAll(ctx, domain.ProductQuery{CategoryID: "cat1"})
		require.NoError(t, err)
		assert.Equal(t, int64(2), total)
		assert.Len(t, res, 2)
	})

	t.Run("Get All Products - Search Text", func(t *testing.T) {
		res, total, err := repo.GetAll(ctx, domain.ProductQuery{Search: "baju"})
		require.NoError(t, err)
		assert.Equal(t, int64(2), total)
		assert.Len(t, res, 2)
	})

	t.Run("Get All Products - Price Range", func(t *testing.T) {
		_, total, err := repo.GetAll(ctx, domain.ProductQuery{MinPrice: 30000, MaxPrice: 60000})
		require.NoError(t, err)
		assert.Equal(t, int64(2), total)
		// Should return Baju Hitam and Baju Putih
	})

	t.Run("Get All Products - In Stock Only", func(t *testing.T) {
		inStock := true
		res, total, err := repo.GetAll(ctx, domain.ProductQuery{InStock: &inStock, CategoryID: "cat1"})
		require.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Equal(t, "Baju Hitam", res[0].Name)
	})

	t.Run("Get All Products - Out of Stock", func(t *testing.T) {
		inStock := false
		res, total, err := repo.GetAll(ctx, domain.ProductQuery{InStock: &inStock})
		require.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Equal(t, "Baju Putih", res[0].Name)
	})

	t.Run("Get All Products - Pagination", func(t *testing.T) {
		res, total, err := repo.GetAll(ctx, domain.ProductQuery{Page: 1, PerPage: 2})
		require.NoError(t, err)
		assert.Equal(t, int64(4), total)
		assert.Len(t, res, 2)
	})

	t.Run("Update Product", func(t *testing.T) {
		p := products[0]
		p.Name = "Baju Hitam Updated"
		p.Price = 55000
		err := repo.Update(ctx, p)
		require.NoError(t, err)

		fetched, err := repo.GetByID(ctx, p.ID)
		require.NoError(t, err)
		assert.Equal(t, "Baju Hitam Updated", fetched.Name)
		assert.Equal(t, float64(55000), fetched.Price)
	})

	t.Run("Delete Product", func(t *testing.T) {
		p := products[1]
		err := repo.Delete(ctx, p.ID)
		require.NoError(t, err)

		_, err = repo.GetByID(ctx, p.ID)
		assert.ErrorIs(t, err, domain.ErrProductNotFound)
	})
}
