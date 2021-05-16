package main

import (
	"main/commontools"
	"main/routing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// ミドルウェアGin宣言
var router *gin.Engine = gin.Default()

func main() {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions(commontools.Fixedparam("sessionName"), store))
	// router.Use(commontools.SampleMiddleware())

	routing.LoadInitial(router)
	routing.LoadGet(router)
	routing.LoadPost(router)
	routing.LoadPut(router)

	router.Run(":" + commontools.Fixedparam("port"))
}
