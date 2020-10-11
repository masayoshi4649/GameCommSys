package get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// Content ... labels
type Content struct {
	Label1 string `json:"label1"`
	Label2 string `json:"label2"`
	Label3 string `json:"label3"`
}

// LangData ... languages
type LangData struct {
	Ja Content `json:"ja"`
	En Content `json:"en"`
}

var langData LangData

// Index ... GET index.html
func Index(ctx *gin.Context) {

	// デバッグ用
	var debuglang = "ja"

	labeljson, err := ioutil.ReadFile("./lang/index.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(labeljson, &langData)

	fmt.Println(langData.Ja.Label1)
	fmt.Println(langData.En.Label1)

	switchLang(debuglang, ctx, langData)
}

func switchLang(lang string, ctx *gin.Context, langData LangData) {
	switch lang {
	case "ja":
		ctx.HTML(200, "index.html", gin.H{"label1": langData.Ja.Label1, "label2": langData.Ja.Label2, "label3": langData.Ja.Label3})
	case "en":
		ctx.HTML(200, "index.html", gin.H{"label1": langData.En.Label1, "label2": langData.En.Label2, "label3": langData.En.Label3})
	default:
		ctx.HTML(200, "index.html", gin.H{"label1": langData.En.Label1, "label2": langData.En.Label2, "label3": langData.En.Label3})
	}
}
