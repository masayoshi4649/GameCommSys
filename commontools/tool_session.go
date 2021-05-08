package commontools

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginCheck ... ログイン済か否か
// return ログインT/F , 言語番号
func LoginCheck(ctx *gin.Context) (bool, int) {

	session := sessions.Default(ctx)

	userid := session.Get("userid")
	lang := session.Get("lang")
	fmt.Println(userid)
	fmt.Println(lang)

	if userid != nil {
		return false, lang.(int)
	}
	return false, 0

}

// LoginFuncion ... ログイン後のセッション作成
func LoginFuncion(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Set("userid", 114514)

}
