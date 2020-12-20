package get

import (
	"github.com/gin-gonic/gin"
)

// Login ... GET login.html FOR DEBUG
func Login(ctx *gin.Context) {
	labelmap := make(map[string]interface{})
	labelmap["discordURL"] = "https://discord.com/api/oauth2/authorize?client_id=764483773567467550&redirect_uri=http%3A%2F%2Fdev.masayoshi4649.top%3A8080%2Floginprocessing&response_type=code&scope=identify%20email"

	ctx.HTML(200, "login.html", labelmap)
}
