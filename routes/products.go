package routes

import (
	"samuelescherini/gintutorial/controllers"
	"samuelescherini/gintutorial/helpers"

	"github.com/gin-gonic/gin"
)

func SetProductRoutes(r *gin.Engine) {

	// all products routes match a controller's function
	r.GET("/products", controllers.GetProducts)
	r.GET("/product/:id", controllers.GetProductById)
	r.POST("/addProduct", controllers.AddProduct)
	r.PUT("/updateProduct/:id", controllers.UpdateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)

	// default route for 404 error
	r.NoRoute(func(ctx *gin.Context) {
		helpers.Error404(ctx, "No route found")
	})

}
