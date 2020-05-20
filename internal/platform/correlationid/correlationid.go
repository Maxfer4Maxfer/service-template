package correlationid

import (
	"context"

	"github.com/lucsky/cuid"
)

// ContextKey is used for context.Context value.
// The value requires a key that is not primitive type.
type ContextKey string // can be unexported

// ContextKeyCorrelationID is the ContextKey for CorrelationID.
const ContextKeyCorrelationID ContextKey = "correlationID"

// Assign will attach a brand new correlation ID to a context.
func Assign(ctx context.Context) (context.Context, string) {
	cID := cuid.New()
	return context.WithValue(ctx, ContextKeyCorrelationID, cID), cID
}

// Extract will get cID from a http correlation and return it as a string.
func Extract(ctx context.Context) string {
	cID := ctx.Value(ContextKeyCorrelationID)
	if ret, ok := cID.(string); ok {
		return ret
	}

	return ""
}
