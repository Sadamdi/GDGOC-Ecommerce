package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ecommerce-backend/internal/domain"
	repoMongo "ecommerce-backend/internal/repository/mongo"
)

func TestCartRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db, teardown := repoMongo.SetupTestMongoDB(t)
	defer teardown()

	repo := repoMongo.NewCartRepository(db)
	ctx := context.Background()

	userID := "test-user-123"

	t.Run("Save and Get Cart", func(t *testing.T) {
		cart := &domain.Cart{
			UserID: userID,
			Items: []domain.CartItem{
				{ProductID: "prod1", Quantity: 2, Price: 100, SubTotal: 200},
			},
			TotalAmount: 200,
		}

		err := repo.Save(ctx, cart)
		require.NoError(t, err)

		fetched, err := repo.GetByUserID(ctx, userID)
		require.NoError(t, err)
		assert.Equal(t, userID, fetched.UserID)
		assert.Len(t, fetched.Items, 1)
		assert.Equal(t, cart.Version, fetched.Version)
	})

	t.Run("Optimistic Locking Conflict", func(t *testing.T) {
		// Fetch the current cart
		fetched1, err := repo.GetByUserID(ctx, userID)
		require.NoError(t, err)

		// Simulate another concurrent fetch
		fetched2, err := repo.GetByUserID(ctx, userID)
		require.NoError(t, err)

		// Both add items and save
		fetched1.TotalAmount = 300
		err = repo.Save(ctx, fetched1)
		require.NoError(t, err)

		fetched2.TotalAmount = 400
		err = repo.Save(ctx, fetched2)
		// Since fetched2 has old version, it should fail
		assert.ErrorIs(t, err, domain.ErrCartConflict)
	})

	t.Run("Delete Cart", func(t *testing.T) {
		err := repo.DeleteByUserID(ctx, userID)
		require.NoError(t, err)

		_, err = repo.GetByUserID(ctx, userID)
		assert.ErrorIs(t, err, domain.ErrCartNotFound)
	})
}
