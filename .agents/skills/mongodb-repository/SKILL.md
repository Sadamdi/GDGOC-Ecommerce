---
name: MongoDB Repository
description: >
  Membuat repository implementation untuk MongoDB menggunakan Official MongoDB
  Go Driver. Mengikuti pattern dan best practices proyek E-Commerce.
---

# MongoDB Repository

## Kapan Skill Ini Digunakan

- Membuat repository implementation baru untuk MongoDB
- Menambah query atau operasi database baru
- Mengoptimasi query MongoDB

## Setup Connection

```go
// internal/pkg/mongodb/connection.go
func Connect(uri string) (*mongo.Database, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, fmt.Errorf("mongodb connect: %w", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        return nil, fmt.Errorf("mongodb ping: %w", err)
    }

    return client.Database(dbName), nil
}
```

## Repository Pattern

```go
// internal/repository/mongo/<entity>_repository.go
package mongo

type entityRepository struct {
    collection *mongo.Collection
}

func NewEntityRepository(db *mongo.Database) domain.EntityRepository {
    return &entityRepository{
        collection: db.Collection("entities"),
    }
}
```

## Common Operations

### Create
```go
func (r *entityRepository) Create(ctx context.Context, entity *domain.Entity) error {
    entity.ID = primitive.NewObjectID().Hex()
    entity.CreatedAt = time.Now()
    entity.UpdatedAt = time.Now()

    _, err := r.collection.InsertOne(ctx, entity)
    if err != nil {
        return fmt.Errorf("entityRepository.Create: %w", err)
    }
    return nil
}
```

### FindByID
```go
func (r *entityRepository) FindByID(ctx context.Context, id string) (*domain.Entity, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, domain.ErrInvalidID
    }

    var entity domain.Entity
    err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&entity)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, domain.ErrEntityNotFound
        }
        return nil, fmt.Errorf("entityRepository.FindByID: %w", err)
    }
    return &entity, nil
}
```

### List (Paginated)
```go
func (r *entityRepository) List(ctx context.Context, page, perPage int) ([]*domain.Entity, int64, error) {
    skip := (page - 1) * perPage
    opts := options.Find().
        SetSkip(int64(skip)).
        SetLimit(int64(perPage)).
        SetSort(bson.D{{Key: "created_at", Value: -1}})

    total, err := r.collection.CountDocuments(ctx, bson.M{})
    if err != nil {
        return nil, 0, fmt.Errorf("entityRepository.List count: %w", err)
    }

    cursor, err := r.collection.Find(ctx, bson.M{}, opts)
    if err != nil {
        return nil, 0, fmt.Errorf("entityRepository.List find: %w", err)
    }
    defer cursor.Close(ctx)

    var entities []*domain.Entity
    if err := cursor.All(ctx, &entities); err != nil {
        return nil, 0, fmt.Errorf("entityRepository.List decode: %w", err)
    }

    return entities, total, nil
}
```

### Update
```go
func (r *entityRepository) Update(ctx context.Context, entity *domain.Entity) error {
    entity.UpdatedAt = time.Now()
    objectID, _ := primitive.ObjectIDFromHex(entity.ID)

    result, err := r.collection.UpdateOne(ctx,
        bson.M{"_id": objectID},
        bson.M{"$set": entity},
    )
    if err != nil {
        return fmt.Errorf("entityRepository.Update: %w", err)
    }
    if result.MatchedCount == 0 {
        return domain.ErrEntityNotFound
    }
    return nil
}
```

### Delete
```go
func (r *entityRepository) Delete(ctx context.Context, id string) error {
    objectID, _ := primitive.ObjectIDFromHex(id)

    result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
    if err != nil {
        return fmt.Errorf("entityRepository.Delete: %w", err)
    }
    if result.DeletedCount == 0 {
        return domain.ErrEntityNotFound
    }
    return nil
}
```

## Best Practices

- ✅ Selalu gunakan `context.Context` dengan timeout
- ✅ Wrap error dengan context: `fmt.Errorf("repoName.Method: %w", err)`
- ✅ Map `mongo.ErrNoDocuments` ke domain error
- ✅ Buat indexes untuk field yang sering di-query
- ✅ Gunakan `bson.M` untuk filter, `bson.D` untuk ordered operations
- ❌ Jangan taruh business logic di repository
- ❌ Jangan return `mongo.*` types ke layer luar

## Index Creation

```go
func (r *entityRepository) createIndexes(ctx context.Context) error {
    indexes := []mongo.IndexModel{
        {
            Keys:    bson.D{{Key: "email", Value: 1}},
            Options: options.Index().SetUnique(true),
        },
        {
            Keys: bson.D{{Key: "created_at", Value: -1}},
        },
    }
    _, err := r.collection.Indexes().CreateMany(ctx, indexes)
    return err
}
```
