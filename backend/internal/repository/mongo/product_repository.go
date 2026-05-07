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

type productRepository struct {
	collection *mongo.Collection
}

// NewProductRepository creates a new product repository instance.
func NewProductRepository(db *mongo.Database) domain.ProductRepository {
	collection := db.Collection("products")

	// Set Index on Name for searching and CategoryID for filtering
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.M{"name": "text"},
		},
		{
			Keys: bson.M{"category_id": 1},
		},
		{
			Keys: bson.M{"price": 1},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _ = collection.Indexes().CreateMany(ctx, indexModels)

	return &productRepository{collection: collection}
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	res, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		product.ID = oid.Hex()
	}

	return nil
}

func (r *productRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrProductNotFound
	}

	var product domain.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetAll(ctx context.Context, query domain.ProductQuery) ([]*domain.Product, int64, error) {
	filter := bson.M{}

	if query.CategoryID != "" {
		filter["category_id"] = query.CategoryID
	}

	if query.Search != "" {
		filter["name"] = bson.M{"$regex": query.Search, "$options": "i"}
	}

	if query.MinPrice > 0 || query.MaxPrice > 0 {
		priceFilter := bson.M{}
		if query.MinPrice > 0 {
			priceFilter["$gte"] = query.MinPrice
		}
		if query.MaxPrice > 0 {
			priceFilter["$lte"] = query.MaxPrice
		}
		filter["price"] = priceFilter
	}

	if query.InStock != nil {
		if *query.InStock {
			filter["stock"] = bson.M{"$gt": 0}
		} else {
			filter["stock"] = 0
		}
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	page := query.Page
	if page < 1 {
		page = 1
	}
	limit := query.PerPage
	if limit < 1 {
		limit = 20
	}
	skip := (page - 1) * limit

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.M{"created_at": -1}) // Default sort descending

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []*domain.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}

	if products == nil {
		products = make([]*domain.Product, 0)
	}

	return products, total, nil
}

func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
	oid, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return domain.ErrProductNotFound
	}

	product.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"stock":       product.Stock,
			"category_id": product.CategoryID,
			"images":      product.Images,
			"updated_at":  product.UpdatedAt,
		},
	}

	res, err := r.collection.UpdateByID(ctx, oid, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}

func (r *productRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrProductNotFound
	}

	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}
