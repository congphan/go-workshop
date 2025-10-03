package primative

import (
	"context"
	"fmt"
)

// query tenant user by tenantUID and tenantUserUID
func QueryTenantUser(tenantUID int, tenantUserUID int) int {
	return tenantUID - tenantUserUID // minus for demonstration purposes
}

const (
	tenantUIDKey string = "tenant_uid"
	userIDKey    string = "user_id"
)

func InitContextLoggger() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, tenantUIDKey, "default-tenant-uid")
	ctx = context.WithValue(ctx, userIDKey, "default-user-id")

	return ctx
}

func PrintLogFromContext(ctx context.Context, message string) {
	tenantUID := ctx.Value(tenantUIDKey)
	userID := ctx.Value(userIDKey)

	// Simulate logging with tenant and user context
	logMessage := fmt.Sprintf("%s: %v, %s: %v - ", tenantUIDKey, tenantUID, userIDKey, userID) + message
	fmt.Println(logMessage)
}
