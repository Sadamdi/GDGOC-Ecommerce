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

type blocklistRepository struct {
	collection *mongo.Collection
}

func NewBlocklistRepository(db *mongo.Database) domain.BlocklistRepository {
	collection := db.Collection("token_blocklist")

	// Set TTL Index on expiry
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"expiry": 1},
		Options: options.Index().SetExpireAfterSeconds(0), // Automatically delete when current time > expiry
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _ = collection.Indexes().CreateOne(ctx, indexModel)

	return &blocklistRepository{collection: collection}
}

func (r *blocklistRepository) AddToBlocklist(ctx context.Context, token string, expiry time.Time) error {
	blocklist := domain.TokenBlocklist{
		Token:     token,
		Expiry:    expiry,
		CreatedAt: time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, blocklist)
	return err
}

func (r *blocklistRepository) IsBlacklisted(ctx context.Context, token string) (bool, error) {
	err := r.collection.FindOne(ctx, bson.M{"token": token}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // Not found = not blacklisted
		}
		return false, err // DB Error
	}

	return true, nil // Found = blacklisted
}
