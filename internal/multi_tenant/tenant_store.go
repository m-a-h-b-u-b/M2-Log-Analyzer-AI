//! Module Name: tenant_store.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Manages all tenants, provides lookup by tenant ID.

package multi_tenant

import "sync"

type TenantManager struct {
	tenants map[string]*Tenant
	mu      sync.RWMutex
}

func NewTenantManager() *TenantManager {
	return &TenantManager{
		tenants: make(map[string]*Tenant),
	}
}

func (m *TenantManager) AddTenant(name string, proc *pipeline.Processor, store Storage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tenants[name] = &Tenant{
		Name:     name,
		Pipeline: proc,
		Store:    store,
	}
}

func (m *TenantManager) GetTenant(name string) (*Tenant, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	t, ok := m.tenants[name]
	return t, ok
}
