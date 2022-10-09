package controllers

import (
	"fmt"
	"net/http"
	"samuelescherini/gintutorial/data"
	"samuelescherini/gintutorial/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data.Products)
}

func GetProductById(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		helpers.BadRequest(ctx, "the id of the product to update is not valid")
		return
	}

	for _, prod := range data.Products {
		if prod.PID == id {
			ctx.IndentedJSON(http.StatusOK, prod)
			return
		}
	}
	helpers.BadRequest(ctx, fmt.Sprintf("no item found with the given id: %d", id))
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

func UpdateProduct(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		helpers.BadRequest(ctx, "the id of the product to update is not valid")
		return
	}

	var newProduct data.Product
	newProduct.PID = id

	if err := ctx.BindJSON(&newProduct); err != nil {
		helpers.BadRequest(ctx, "the new product is nil")
		return
	}

	if newProduct.Quantity < 0 {
		helpers.BadRequest(ctx, "the quantity must be greater than 0")
		return
	}

	for i, prod := range data.Products {
		if prod.PID == id {
			data.Products[i] = newProduct
			ctx.IndentedJSON(http.StatusOK, data.Products[i])
			return
		}
	}
	helpers.BadRequest(ctx, fmt.Sprintf("no item found with the given id: %d", id))
}

func DeleteProduct(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		helpers.BadRequest(ctx, "the id of the product to update is not valid")
		return
	}

	data.Products = append(data.Products[:id-1], data.Products[id:]...)
	ctx.IndentedJSON(http.StatusCreated, data.Products)

}
