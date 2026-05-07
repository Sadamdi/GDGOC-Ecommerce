package mongo

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ecommerce-backend/internal/domain"
)

type categoryRepository struct {
	collection *mongo.Collection
}

// NewCategoryRepository creates a new category repository instance and sets up indexes.
func NewCategoryRepository(db *mongo.Database) domain.CategoryRepository {
	collection := db.Collection("categories")

	// Set Unique Index on Name
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _ = collection.Indexes().CreateOne(ctx, indexModel)

	return &categoryRepository{collection: collection}
}

func (r *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return domain.ErrCategoryAlreadyExists
		}
		return err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		category.ID = oid.Hex()
	}

	return nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrCategoryNotFound
	}

	var category domain.Category
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&category)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrCategoryNotFound
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetByName(ctx context.Context, name string) (*domain.Category, error) {
	var category domain.Category
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(&category)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrCategoryNotFound
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]*domain.Category, error) {
	var categories []*domain.Category

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	if categories == nil {
		categories = make([]*domain.Category, 0)
	}

	return categories, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	oid, err := primitive.ObjectIDFromHex(category.ID)
	if err != nil {
		return domain.ErrCategoryNotFound
	}

	category.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"name":        category.Name,
			"description": category.Description,
			"updated_at":  category.UpdatedAt,
		},
	}

	res, err := r.collection.UpdateByID(ctx, oid, update)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return domain.ErrCategoryAlreadyExists
		}
		return err
	}

	if res.MatchedCount == 0 {
		return domain.ErrCategoryNotFound
	}

	return nil
}
