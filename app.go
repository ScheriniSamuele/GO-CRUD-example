package main

import (
	"net/http"
	"samuelescherini/gintutorial/data"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// create and init a "store"
	products := data.InitProducts()

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, products)
	})

	r.POST("/addProduct", func(ctx *gin.Context) {

		var newProduct data.Product

		if err := ctx.BindJSON(&newProduct); err != nil {
			badRequest(ctx, "the new product is nil")
			return
		}

		if newProduct.Quantity < 0 {
			badRequest(ctx, "the quantity must be greater than 0")
			return
		}

		if len(products) != 0 {
			newProduct.PID = products[len(products)-1].PID + 1
		} else {
			newProduct.PID = 1
		}

		products = append(products, newProduct)
		ctx.IndentedJSON(http.StatusCreated, newProduct)
	})

	r.NoRoute(func(ctx *gin.Context) {
		badRequest(ctx, "No route found")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func badRequest(ctx *gin.Context, errMsg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
}
