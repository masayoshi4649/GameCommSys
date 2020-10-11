package routing

import (
	"./get"
	"github.com/gin-gonic/gin"
)

// LoadInitial ... "common"
func LoadInitial(router *gin.Engine) {
	router.StaticFile("/favicon.ico", "common/favicon.ico")
	router.LoadHTMLGlob("view/*.html")

}

// LoadGet ... GET Request
func LoadGet(router *gin.Engine) {
	router.GET("/", get.Index)
	router.GET("/login", get.Login)
	router.GET("/loginprocessing", get.LoginProcessing)
}

// LoadPost ... POST Request
func LoadPost(router *gin.Engine) {

}

// LoadPut ... PUT Request
func LoadPut(router *gin.Engine) {

}
