package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetupTestMongoDB starts a MongoDB testcontainer and returns the db instance and a teardown function.
func SetupTestMongoDB(t *testing.T) (*mongo.Database, func()) {
	ctx := context.Background()

	// Create MongoDB container
	mongodbContainer, err := mongodb.Run(ctx, "mongo:7.0")
	require.NoError(t, err)

	endpoint, err := mongodbContainer.ConnectionString(ctx)
	require.NoError(t, err)

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(endpoint)
	client, err := mongo.Connect(ctx, clientOptions)
	require.NoError(t, err)

	// Ping to ensure connection is established
	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = client.Ping(ctxPing, nil)
	require.NoError(t, err)

	db := client.Database("test_ecommerce_db")

	// Teardown function
	teardown := func() {
		_ = client.Disconnect(ctx)
		_ = testcontainers.TerminateContainer(mongodbContainer)
	}

	return db, teardown
}
