package ping

import "github.com/gin-gonic/gin"

// Pong returns a "pong" to the browser
// when ever /ping is called
//
// Used as heart beat
func Pong(c *gin.Context) {
	c.String(200, "Pong")
}
