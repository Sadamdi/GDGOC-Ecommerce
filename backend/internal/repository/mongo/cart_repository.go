package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ecommerce-backend/internal/domain"
)

type cartRepository struct {
	collection *mongo.Collection
}

// NewCartRepository creates a new cart repository instance
func NewCartRepository(db *mongo.Database) domain.CartRepository {
	collection := db.Collection("carts")

	// Create unique index on user_id
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"user_id": 1},
		Options: options.Index().SetUnique(true),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _ = collection.Indexes().CreateOne(ctx, indexModel)

	return &cartRepository{collection: collection}
}

func (r *cartRepository) GetByUserID(ctx context.Context, userID string) (*domain.Cart, error) {
	var cart domain.Cart
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&cart)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrCartNotFound
		}
		return nil, err
	}

	// Ensure items array is not nil for response
	if cart.Items == nil {
		cart.Items = []domain.CartItem{}
	}

	return &cart, nil
}

func (r *cartRepository) Save(ctx context.Context, cart *domain.Cart) error {
	cart.UpdatedAt = time.Now()

	filter := bson.M{"user_id": cart.UserID}
	if cart.Version > 0 {
		filter["version"] = cart.Version
	}

	update := bson.M{
		"$set": bson.M{
			"items":        cart.Items,
			"total_amount": cart.TotalAmount,
			"updated_at":   cart.UpdatedAt,
		},
		"$inc": bson.M{
			"version": 1,
		},
	}

	opts := options.Update()
	if cart.Version == 0 {
		opts.SetUpsert(true)
	}

	res, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 && res.UpsertedCount == 0 {
		return domain.ErrCartConflict
	}

	cart.Version++
	return nil
}

func (r *cartRepository) DeleteByUserID(ctx context.Context, userID string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"user_id": userID})
	return err
}
