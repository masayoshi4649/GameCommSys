package get

import (
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Login ... GET login.html FOR DEBUG
func Login(ctx *gin.Context) {
	labelmap := make(map[string]interface{})
	var discordURI string = Fixedparam("discordURI")
	discordURI = strings.Replace(discordURI, "【clientID】", Fixedparam("clientID"), -1)
	var redirectURI string = Fixedparam("myhostname") + "loginprocessing"
	discordURI = strings.Replace(discordURI, "【enc_redirect_uri】", url.QueryEscape(redirectURI), -1)

	labelmap["discordURL"] = discordURI

	ctx.HTML(200, "login.html", labelmap)
}
