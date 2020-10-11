package get

import "github.com/gin-gonic/gin"

// Index ... GET index.html
func Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}
