package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
)

type BaseController struct {}

func (b BaseController) Ctx(c *gin.Context) context.Context {
	tenantId := c.MustGet("tenantId")
	return context.WithValue(c.Request.Context(), "tenantId", tenantId)
}