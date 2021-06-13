package services

import (
	"context"
	"example/src/models"
	"time"
)

type articleService struct {
	articleRepo    models.ArticleRepository
	contextTimeout time.Duration
}

// NewArticleService will create new an articleService object representation of models.NewArticleService interface
func NewArticleService(a models.ArticleRepository, timeout time.Duration) models.ArticleService {
	return &articleService{
		articleRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *articleService) Fetch(c context.Context) (res []models.Article, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (a *articleService) GetByID(c context.Context, id int64) (res models.Article, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *articleService) Update(c context.Context, ar *models.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.articleRepo.Update(ctx, ar)
}

func (a *articleService) GetByTitle(c context.Context, title string) (res models.Article, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.articleRepo.GetByTitle(ctx, title)
	if err != nil {
		return
	}

	return
}

func (a *articleService) Store(c context.Context, m *models.Article) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.GetByTitle(ctx, m.Title)
	if existedArticle != (models.Article{}) {
		return models.ErrConflict
	}

	err = a.articleRepo.Store(ctx, m)
	return
}

func (a *articleService) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, err := a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedArticle == (models.Article{}) {
		return models.ErrNotFound
	}
	return a.articleRepo.Delete(ctx, id)
}
