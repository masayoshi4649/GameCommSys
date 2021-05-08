package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/masayoshi4649/GameCommSys/commontools"
	"github.com/masayoshi4649/GameCommSys/routing"
)

// ミドルウェアGin宣言
var router *gin.Engine = gin.Default()

func main() {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions(commontools.Fixedparam("sessionName"), store))

	routing.LoadInitial(router)
	routing.LoadGet(router)
	routing.LoadPost(router)
	routing.LoadPut(router)

	router.Run(":" + commontools.Fixedparam("port"))
}
