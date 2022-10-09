package controllers

import (
	"net/http"
	"samuelescherini/gintutorial/data"
	"samuelescherini/gintutorial/helpers"

	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data.Products)
}

func AddProduct(ctx *gin.Context) {
	var newProduct data.Product

	if err := ctx.BindJSON(&newProduct); err != nil {
		helpers.BadRequest(ctx, "the new product is nil")
		return
	}

	if newProduct.Quantity < 0 {
		helpers.BadRequest(ctx, "the quantity must be greater than 0")
		return
	}

	if len(data.Products) != 0 {
		newProduct.PID = data.Products[len(data.Products)-1].PID + 1
	} else {
		newProduct.PID = 1
	}

	data.Products = append(data.Products, newProduct)
	ctx.IndentedJSON(http.StatusCreated, newProduct)
}

func UpdateProduct(ctx *gin.Context) {}

func DeleteProduct(ctx *gin.Context) {}
