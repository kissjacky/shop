package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Controller
}

func (*Product) Index(c *gin.Context) {

	c.String(http.StatusOK, "controller product")
}
func (*Product) Addd(c *gin.Context) {
}
