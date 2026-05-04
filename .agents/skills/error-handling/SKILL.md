---
name: Error Handling
description: >
  Mengimplementasikan error handling yang konsisten di semua layer aplikasi.
  Termasuk domain errors, error wrapping, error-to-HTTP mapping, dan logging.
---

# Error Handling

## Kapan Skill Ini Digunakan

- Menambah error type baru
- Mengimplementasikan error handling di use case
- Mapping error ke HTTP response di handler
- Setup logging untuk errors

## Domain Errors (`internal/domain/errors.go`)

```go
package domain

import "errors"

// Sentinel errors — gunakan untuk errors.Is() matching
var (
    // Not Found
    ErrUserNotFound    = errors.New("user not found")
    ErrProductNotFound = errors.New("product not found")
    ErrOrderNotFound   = errors.New("order not found")
    ErrCartEmpty       = errors.New("cart is empty")

    // Validation
    ErrInvalidInput    = errors.New("invalid input")
    ErrInvalidEmail    = errors.New("invalid email format")
    ErrWeakPassword    = errors.New("password too weak")
    ErrInvalidID       = errors.New("invalid id format")

    // Conflict
    ErrEmailExists     = errors.New("email already exists")

    // Business Logic
    ErrInsufficientStock = errors.New("insufficient stock")

    // Auth
    ErrUnauthorized    = errors.New("unauthorized")
    ErrForbidden       = errors.New("forbidden")
    ErrInvalidToken    = errors.New("invalid token")
    ErrTokenExpired    = errors.New("token expired")
)
```

## Error Wrapping in UseCase

```go
func (uc *useCase) Method(ctx context.Context, ...) error {
    result, err := uc.repo.FindByID(ctx, id)
    if err != nil {
        // Wrap dengan context tapi preserve original error
        return fmt.Errorf("useCaseName.Method: %w", err)
    }
    return nil
}
```

## Error-to-HTTP Mapping in Handler

```go
func writeError(w http.ResponseWriter, err error) {
    status := http.StatusInternalServerError
    code := "INTERNAL_ERROR"
    message := "An unexpected error occurred"

    switch {
    case errors.Is(err, domain.ErrUserNotFound),
         errors.Is(err, domain.ErrProductNotFound),
         errors.Is(err, domain.ErrOrderNotFound):
        status = http.StatusNotFound
        code = "NOT_FOUND"
        message = err.Error()

    case errors.Is(err, domain.ErrInvalidInput),
         errors.Is(err, domain.ErrInvalidEmail),
         errors.Is(err, domain.ErrInvalidID):
        status = http.StatusBadRequest
        code = "VALIDATION_ERROR"
        message = err.Error()

    case errors.Is(err, domain.ErrEmailExists):
        status = http.StatusConflict
        code = "CONFLICT"
        message = err.Error()

    case errors.Is(err, domain.ErrUnauthorized),
         errors.Is(err, domain.ErrInvalidToken),
         errors.Is(err, domain.ErrTokenExpired):
        status = http.StatusUnauthorized
        code = "UNAUTHORIZED"
        message = err.Error()

    case errors.Is(err, domain.ErrForbidden):
        status = http.StatusForbidden
        code = "FORBIDDEN"
        message = err.Error()

    case errors.Is(err, domain.ErrInsufficientStock):
        status = http.StatusUnprocessableEntity
        code = "BUSINESS_ERROR"
        message = err.Error()
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(ErrorResponse{
        Success: false,
        Message: message,
        Error:   ErrorDetail{Code: code},
    })
}
```

## Rules

- ✅ Gunakan sentinel errors di domain layer
- ✅ Wrap error dengan `fmt.Errorf("context: %w", err)`
- ✅ Gunakan `errors.Is()` untuk matching
- ✅ Log error di handler layer (bukan di domain/usecase)
- ✅ Jangan expose internal error details ke client (500 errors)
- ❌ Jangan ignore error (no `_` for errors)
- ❌ Jangan buat `context.Background()` di tengah chain
- ❌ Jangan log sensitive data (passwords, tokens)
