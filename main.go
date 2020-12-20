package main

import (
	"./routing"
	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// ServerGeneral ... "./conf/serverGeneral.tml"
type ServerGeneral struct {
	Port        string
	SessionName string
}

var serverGeneral ServerGeneral

// ミドルウェアGin宣言
var router *gin.Engine = gin.Default()

func init() {
	// "./conf/serverGeneral.tml"
	_, err := toml.DecodeFile("./conf/serverGeneral.tml", &serverGeneral)
	if err != nil {
		panic(err)
	}
}

func main() {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions(serverGeneral.SessionName, store))

	routing.LoadInitial(router)
	routing.LoadGet(router)
	routing.LoadPost(router)
	routing.LoadPut(router)

	router.Run(":" + serverGeneral.Port)
}

// リフレッシュトークンのためのバッチ処理がいるかも
