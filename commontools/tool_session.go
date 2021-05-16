package commontools

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginFuncion ... ログイン後のセッション作成
func LoginFuncion(ctx *gin.Context, userid string) {
	session := sessions.Default(ctx)
	session.Set("userid", userid)
	session.Save()
}

// LoginCheck ... ログイン済か否か
func LoginCheck(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	userid := session.Get("userid")
	fmt.Println(userid)
	if userid != nil {
		return true
	}
	return false
}

// func SampleMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// ルーティング内の処理の前に実行される
// 		fmt.Println("test1")
// 		ctx.Next()
// 		// ルーティング内の処理の後に実行される
// 		fmt.Println("test2")
// 	}
// }

func LogoutFunction(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	println("Session clear")
}
