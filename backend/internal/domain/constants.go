package domain

type ContextKey string

//nolint:gosec // false positive
const (
	CtxKeyUserID      ContextKey = "user_id"
	CtxKeyUserRole    ContextKey = "user_role"
	CtxKeyTokenString ContextKey = "token_string"
	CtxKeyTokenExpiry ContextKey = "token_expiry"
)
