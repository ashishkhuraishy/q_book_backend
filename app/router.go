package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApp will be the entry point to out router
// which will handle all the calls to our api
func StartApp() {
	urlMappings()
	router.Run("localhost:8080")
}
