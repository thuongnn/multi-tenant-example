package controllers

import (
	"example/src/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleController  represent the httpHandler for article
type ArticleController struct {
	base BaseController
	AUsecase models.ArticleService
}

// NewArticleController will initialize the articles/ resources endpoint
func NewArticleController(e *gin.Engine, us models.ArticleService) {
	handler := &ArticleController{
		AUsecase: us,
	}

	grp1 := e.Group("/v1")
	{
		grp1.GET("/articles", handler.FetchArticle)
		grp1.POST("/articles", handler.Store)
		grp1.GET("/articles/:id", handler.GetByID)
		grp1.DELETE("/articles/:id", handler.Delete)
	}
}

// FetchArticle will fetch the article based on given params
func (a *ArticleController) FetchArticle(c *gin.Context) {
	ctx := a.base.Ctx(c)

	listAr, err := a.AUsecase.Fetch(ctx)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, listAr)
}

// GetByID will get article by given id
func (a *ArticleController) GetByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := a.base.Ctx(c)

	art, err := a.AUsecase.GetByID(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, art)
}

func isRequestValid(m *models.Article) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the article by given request body
func (a *ArticleController) Store(c *gin.Context) {
	var article models.Article
	err := c.Bind(&article)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var ok bool
	if ok, err = isRequestValid(&article); !ok {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := a.base.Ctx(c)
	err = a.AUsecase.Store(ctx, &article)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}

// Delete will delete article by given param
func (a *ArticleController) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
		return
	}

	id := int64(idP)
	ctx := a.base.Ctx(c)

	err = a.AUsecase.Delete(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
