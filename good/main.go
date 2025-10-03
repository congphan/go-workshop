package main

import (
	"context"
	"fmt"

	"github.com/congphan/go-workshop/good/id"
	"github.com/congphan/go-workshop/good/primative"
	"github.com/congphan/go-workshop/good/tenant"
)

func main() {
	// primative type usage
	var (
		primativeTenantUID     primative.TenantUID     = 5
		primativeTenantUserUID primative.TenantUserUID = 2
	)

	fmt.Println(primative.QueryTenantUser(primativeTenantUID, primativeTenantUserUID)) // Output: 3

	var (
		tenantRepo     tenant.TenantRepo     = &tenant.TenantRepoMock{}
		tenantUserRepo tenant.TenantUserRepo = &tenant.TenantUserRepoMock{}
	)

	// spefic identifer technique in your domain model
	var (
		tenantUID    id.ID = "01F8MECHZX3TBDSZ7XRADM79XE"
		tenantUserID id.ID = "01F8MECHZX3TBDSZ7XRADM79XF"
	)

	tenant, err := tenantRepo.FindTenant(tenantUID)
	if err != nil {
		fmt.Println("Error finding tenant:", err)
	}

	tenantUser, err := tenantUserRepo.FindTenantUser(tenantUserID)
	if err != nil {
		fmt.Println("Error finding tenant user:", err)
	}

	tenantUsers, err := tenantUserRepo.FindTenantUsers(tenantUID)
	if err != nil {
		fmt.Println("Error finding tenant users:", err)
	}

	fmt.Println("Tenant:", tenant)
	fmt.Println("Tenant User:", tenantUser)
	fmt.Println("Tenant Users:", tenantUsers)

	// context usage
	ctx := primative.InitContextLoggger()
	ctx = context.WithValue(ctx, "tenant_uid", "other_tenant_uid") // mistakenly override tenant_uid in context
	ctx = context.WithValue(ctx, "user_id", "other_user_id")       // mistakenly override user_id in context
	primative.PrintLogFromContext(ctx, "This is a log message.")
}
