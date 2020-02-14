package controller

import (
	"net/http"
	"shop-gin/model"
	"shop-gin/model/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductGroup struct {
	Controller
}

func (*ProductGroup) GetGroup(c *gin.Context) {
	pg := product.ProductGroup{}

	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusOK, RetErr(1, "", ""))
		return
	}
	pg.ID = uint(id)
	ok := model.Get(&pg, nil)
	if !ok {
		c.JSON(http.StatusNotFound, RetErr(2, "", ""))
		return
	}
	c.JSON(http.StatusOK, RetOk(pg))
}
