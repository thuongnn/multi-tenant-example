package repositories

import (
	"context"
	"gorm.io/gorm"
)

type baseRepository struct {
	dbCtx *gorm.DB
}

func (b baseRepository) tenantCtx(ctx context.Context) *gorm.DB {
	return b.dbCtx.Scopes(func (db *gorm.DB) *gorm.DB{
		return db.Where("tenant_id = ?", ctx.Value("tenantId").(int64))
	})
}

func (b baseRepository) TenantID(ctx context.Context) int64 {
	return ctx.Value("tenantId").(int64)
}