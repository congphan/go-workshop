package main

import (
	"fmt"

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
		tenantUID    tenant.TenantUID    = "01F8MECHZX3TBDSZ7XRADM79XE"
		tenantUserID tenant.TenantUserID = "01F8MECHZX3TBDSZ7XRADM79XF"
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
	primative.PrintLogFromContext(ctx, "This is a log message.")
}
