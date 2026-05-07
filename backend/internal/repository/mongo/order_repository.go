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

type orderRepository struct {
	collection *mongo.Collection
}

// NewOrderRepository creates a new order repository instance
func NewOrderRepository(db *mongo.Database) domain.OrderRepository {
	return &orderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *orderRepository) Create(ctx context.Context, order *domain.Order) error {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		order.ID = oid.Hex()
	}

	return nil
}

func (r *orderRepository) FindByID(ctx context.Context, id string) (*domain.Order, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrOrderNotFound
	}

	var order domain.Order
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}

	return &order, nil
}

func (r *orderRepository) FindByUserID(ctx context.Context, userID string, page, limit int) ([]*domain.Order, int64, error) {
	filter := bson.M{"user_id": userID}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"created_at": -1})

	if page > 0 && limit > 0 {
		skip := (page - 1) * limit
		findOptions.SetSkip(int64(skip))
		findOptions.SetLimit(int64(limit))
	}

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var orders []*domain.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, 0, err
	}

	if orders == nil {
		orders = make([]*domain.Order, 0)
	}

	return orders, total, nil
}

func (r *orderRepository) FindAll(ctx context.Context, page, limit int) ([]*domain.Order, int64, error) {
	filter := bson.M{}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"created_at": -1})

	if page > 0 && limit > 0 {
		skip := (page - 1) * limit
		findOptions.SetSkip(int64(skip))
		findOptions.SetLimit(int64(limit))
	}

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var orders []*domain.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, 0, err
	}

	if orders == nil {
		orders = make([]*domain.Order, 0)
	}

	return orders, total, nil
}

func (r *orderRepository) UpdateStatus(ctx context.Context, id string, status domain.OrderStatus) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrOrderNotFound
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return domain.ErrOrderNotFound
	}

	return nil
}
