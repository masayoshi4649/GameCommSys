package routing

import (
	"github.com/gin-gonic/gin"
)

// LoadInitial ... "common"
func LoadInitial(router *gin.Engine) {
	router.StaticFile("/favicon.ico", "view/favicon.ico")
	router.Static("/src/css", "view/css")
	router.LoadHTMLGlob("view/*.html")

}

// LoadGet ... GET Request
func LoadGet(router *gin.Engine) {
	router.GET("/", Index)
	router.GET("/login", Login)
	router.GET("/loginprocessing", LoginProcessing)
	router.GET("/logout", Logout)
}

// LoadPost ... POST Request
func LoadPost(router *gin.Engine) {

}

// LoadPut ... PUT Request
func LoadPut(router *gin.Engine) {

}
