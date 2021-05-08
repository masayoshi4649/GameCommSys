package get

import (
	"main/commontools"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Login ... GET login.html FOR DEBUG
func Login(ctx *gin.Context) {
	labelmap := make(map[string]interface{})
	var discordURI string = commontools.Fixedparam("discordURI")
	discordURI = strings.Replace(discordURI, "【clientID】", commontools.Fixedparam("clientID"), -1)
	var redirectURI string = commontools.Fixedparam("myhostname") + "loginprocessing"
	discordURI = strings.Replace(discordURI, "【enc_redirect_uri】", url.QueryEscape(redirectURI), -1)

	labelmap["discordURL"] = discordURI

	ctx.HTML(200, "login.html", labelmap)
}
