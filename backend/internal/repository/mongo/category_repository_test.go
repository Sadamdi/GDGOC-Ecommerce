package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ecommerce-backend/internal/domain"
	repoMongo "ecommerce-backend/internal/repository/mongo"
)

func TestCategoryRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db, teardown := repoMongo.SetupTestMongoDB(t)
	defer teardown()

	repo := repoMongo.NewCategoryRepository(db)
	ctx := context.Background()

	t.Run("Create and Get Category", func(t *testing.T) {
		cat := &domain.Category{
			Name:        "Electronics",
			Description: "Gadgets and more",
		}

		err := repo.Create(ctx, cat)
		require.NoError(t, err)
		assert.NotEmpty(t, cat.ID)

		// Get By ID
		fetched, err := repo.GetByID(ctx, cat.ID)
		require.NoError(t, err)
		assert.Equal(t, "Electronics", fetched.Name)

		// Get By Name
		fetchedByName, err := repo.GetByName(ctx, "Electronics")
		require.NoError(t, err)
		assert.Equal(t, cat.ID, fetchedByName.ID)
	})

	t.Run("Create Duplicate Name", func(t *testing.T) {
		cat1 := &domain.Category{Name: "Fashion"}
		err := repo.Create(ctx, cat1)
		require.NoError(t, err)

		cat2 := &domain.Category{Name: "Fashion"}
		err = repo.Create(ctx, cat2)
		assert.ErrorIs(t, err, domain.ErrCategoryAlreadyExists)
	})

	t.Run("Update Category", func(t *testing.T) {
		cat := &domain.Category{Name: "Books", Description: "Old desc"}
		_ = repo.Create(ctx, cat)

		cat.Description = "New desc"
		// Simulate update context
		time.Sleep(10 * time.Millisecond)
		err := repo.Update(ctx, cat)
		require.NoError(t, err)

		fetched, _ := repo.GetByID(ctx, cat.ID)
		assert.Equal(t, "New desc", fetched.Description)
		assert.True(t, fetched.UpdatedAt.After(fetched.CreatedAt))
	})

	t.Run("Get All Categories", func(t *testing.T) {
		categories, err := repo.GetAll(ctx)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(categories), 3) // From previous tests
	})
}
