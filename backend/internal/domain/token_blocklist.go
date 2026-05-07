package domain

import (
	"context"
	"time"
)

// TokenBlocklist menyimpan JWT token yang sudah di-logout atau diblacklist
type TokenBlocklist struct {
	ID        string    `bson:"_id,omitempty"`
	Token     string    `bson:"token"`
	Expiry    time.Time `bson:"expiry"` // Untuk TTL index (auto hapus ketika expired)
	CreatedAt time.Time `bson:"created_at"`
}

// BlocklistRepository mendefinisikan operasi ke database untuk token blocklist
type BlocklistRepository interface {
	AddToBlocklist(ctx context.Context, token string, expiry time.Time) error
	IsBlacklisted(ctx context.Context, token string) (bool, error)
}
