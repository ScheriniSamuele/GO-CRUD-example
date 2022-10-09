package routes

import (
	"samuelescherini/gintutorial/controllers"
	"samuelescherini/gintutorial/helpers"

	"github.com/gin-gonic/gin"
)

func SetProductController(r *gin.Engine) {

	r.GET("/products", controllers.GetProducts)
	r.POST("/addProduct", controllers.AddProduct)
	r.PUT("/updateProduct/{id}", controllers.UpdateProduct)
	r.DELETE("/deleteProduct/{id}", controllers.DeleteProduct)

	r.NoRoute(func(ctx *gin.Context) {
		helpers.BadRequest(ctx, "No route found")
	})

}