package main

// import "fmt"
import (
	"fmt"
	"net/http"

	"shop-gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	type ss string
	var 我 ss = "avc"
	fmt.Println(我)

	// Simple group: v1
	v1 := router.Group("/user")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "abc")
		})
		v1.GET("/logout", func(c *gin.Context) {
			c.String(http.StatusOK, "logout")
		})
	}
	p := &controller.Product{}
	pg := &controller.ProductGroup{}
	v1 = router.Group("/product")
	{
		v1.GET("/", p.Index)
		v1.GET("/group/get", pg.GetGroup)
	}

	router.Run(":8080")
}
