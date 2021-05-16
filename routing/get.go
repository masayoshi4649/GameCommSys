package routing

import (
	"encoding/json"
	"fmt"
	"main/commontools"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Index ... GET index.html
func Index(ctx *gin.Context) {
	var uri string = "index"
	lang := 1
	var label string = commontools.Datalabel(uri, lang)

	var labelmap map[string]interface{}
	json.Unmarshal([]byte(label), &labelmap)

	isLogin := commontools.LoginCheck(ctx)
	labelmap["test"] = "てすと"
	labelmap["isLogin"] = isLogin

	ctx.HTML(200, "index.html", labelmap)
	fmt.Println(labelmap)
}

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

func Logout(ctx *gin.Context) {
	commontools.LogoutFunction(ctx)
}
