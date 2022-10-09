package controllers

import (
	"fmt"
	"net/http"
	"samuelescherini/gintutorial/data"
	"samuelescherini/gintutorial/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Route: /products
// @Description: the function returns all products in the store array
func GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data.Products)
}

// @Route: /products/:id
// @Description: the function returns one product matching the requested id
func GetProductById(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid) // the id in the params is a string but we need an int to compare

	if err != nil {
		helpers.BadRequest(ctx, "the id of the product to update is not valid") // if Atoi fails
		return
	}

	for _, prod := range data.Products {
		// when the id matches we return the item
		if prod.PID == id {
			ctx.IndentedJSON(http.StatusOK, prod)
			return
		}
	}
	// if we do not find the item the id is not present in the store
	helpers.BadRequest(ctx, fmt.Sprintf("no item found with the given id: %d", id))
}

// @Route: /addProduct
// @Description: the function register a new product to the list
func AddProduct(ctx *gin.Context) {
	var newProduct data.Product

	// bind the json body to newProduct
	if err := ctx.BindJSON(&newProduct); err != nil {
		helpers.BadRequest(ctx, "the new product is nil")
		return
	}

	// quantity < 0 is not valid
	if newProduct.Quantity < 0 {
		helpers.BadRequest(ctx, "the quantity must be greater than 0")
		return
	}

	// the new id is the last product id + 1
	if len(data.Products) != 0 {
		newProduct.PID = data.Products[len(data.Products)-1].PID + 1
	} else {
		newProduct.PID = 1
	}

	data.Products = append(data.Products, newProduct)
	ctx.IndentedJSON(http.StatusCreated, newProduct)
}

// @Route: /updateProduct/:id
// @Description: the function updates one product matching the requested id
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
			data.Products[i] = newProduct // replace existing content
			ctx.IndentedJSON(http.StatusOK, data.Products[i])
			return
		}
	}
	helpers.BadRequest(ctx, fmt.Sprintf("no item found with the given id: %d", id))
}

// @Route: /product/:id
// @Description: the function deletes one product matching the requested id
func DeleteProduct(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		helpers.BadRequest(ctx, "the id of the product to update is not valid")
		return
	}

	// delete the item using Go append method and slice's properties
	data.Products = append(data.Products[:id-1], data.Products[id:]...)
	ctx.IndentedJSON(http.StatusCreated, data.Products)

}
