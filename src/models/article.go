package models

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// Article ...
type Article struct {
	ID        int64     `json:"id"`
	TenantID  int64     `json:"tenant_id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Article) TableName() string {
	return "article"
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate:")
	fmt.Println(a.TenantID)
	if a.TenantID == 0 {
		err = errors.New("missing tenantID in Article model")
	}

	return
}

// ArticleService represent the article's usecases
type ArticleService interface {
	Fetch(ctx context.Context) ([]Article, error)
	GetByID(ctx context.Context, id int64) (Article, error)
	Update(ctx context.Context, ar *Article) error
	GetByTitle(ctx context.Context, title string) (Article, error)
	Store(context.Context, *Article) error
	Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repositories contract
type ArticleRepository interface {
	Fetch(ctx context.Context) (res []Article, err error)
	GetByID(ctx context.Context, id int64) (Article, error)
	GetByTitle(ctx context.Context, title string) (Article, error)
	Update(ctx context.Context, ar *Article) error
	Store(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int64) error
}
