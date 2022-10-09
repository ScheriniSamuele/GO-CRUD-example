package main

import (
	"samuelescherini/gintutorial/data"
	"samuelescherini/gintutorial/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// insert some fake data in the global array
	data.InitProducts()

	// use Routes
	routes.SetProductRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
