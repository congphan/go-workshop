package tenant

import (
	"github.com/congphan/go-workshop/good/id"
)

type Tenant struct {
	ID   id.ID
	Name string
}

func (t *Tenant) String() string {
	return `{"id":"` + string(t.ID) + `","name":"` + t.Name + `"}`
}

type TenantUser struct {
	ID        id.ID
	TenantUID id.ID
	Name      string
}

func (tu *TenantUser) String() string {
	return `{"id":"` + string(tu.ID) + `","tenant_uid":"` + string(tu.TenantUID) + `","name":"` + tu.Name + `"}`
}

type TenantRepo interface {
	FindTenant(tenantUID id.ID) (*Tenant, error)
}

type TenantUserRepo interface {
	FindTenantUsers(tenantUID id.ID) ([]*TenantUser, error)
	FindTenantUser(tenantUserID id.ID) (*TenantUser, error)
}

var (
	// Interface Satisfaction Verification
	_ TenantRepo     = (*TenantRepoMock)(nil)
	_ TenantUserRepo = (*TenantUserRepoMock)(nil)
)

// Mocks for demonstration purposes
type TenantRepoMock struct{}
type TenantUserRepoMock struct{}

func (r *TenantRepoMock) FindTenant(tenantUID id.ID) (*Tenant, error) {
	return &Tenant{ID: tenantUID, Name: "Mock Tenant"}, nil
}

func (r *TenantUserRepoMock) FindTenantUsers(tenantUID id.ID) ([]*TenantUser, error) {
	return []*TenantUser{
		{ID: "01F8MECHZX3TBDSZ7XRADM79XE", TenantUID: tenantUID, Name: "Mock User 1"},
		{ID: "01F8MECHZX3TBDSZ7XRADM79XF", TenantUID: tenantUID, Name: "Mock User 2"},
	}, nil
}

func (r *TenantUserRepoMock) FindTenantUser(tenantUserID id.ID) (*TenantUser, error) {
	return &TenantUser{ID: tenantUserID, TenantUID: "01F8MECHZX3TBDSZ7XRADM79XE", Name: "Mock User"}, nil
}
