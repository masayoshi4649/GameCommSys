package get

import "github.com/gin-gonic/gin"

// Login ... GET login.html FOR DEBUG
func Login(ctx *gin.Context) {
	ctx.HTML(200, "login.html", gin.H{})
}
