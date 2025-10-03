package primative

import (
	"context"
	"fmt"
)

type (
	TenantUID     int
	TenantUserUID int
)

// query tenant user by tenantUID and tenantUserUID
func QueryTenantUser(tenantUID TenantUID, tenantUserUID TenantUserUID) int {
	return int(tenantUID) - int(tenantUserUID) // minus for demonstration purposes
}

type (
	contextKey string
)

const (
	contextKeyTenantUID contextKey = "tenant_uid"
	contextKeyUserID    contextKey = "user_id"
)

func InitContextLoggger() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextKeyTenantUID, "default-tenant-uid")
	ctx = context.WithValue(ctx, contextKeyUserID, "default-user-id")

	return ctx
}

func PrintLogFromContext(ctx context.Context, message string) {
	tenantUID := ctx.Value(contextKeyTenantUID)
	userID := ctx.Value(contextKeyUserID)

	// Simulate logging with tenant and user context
	logMessage := fmt.Sprintf("%s: %v, %s: %v - ", contextKeyTenantUID, tenantUID, contextKeyUserID, userID) + message
	fmt.Println(logMessage)
}
