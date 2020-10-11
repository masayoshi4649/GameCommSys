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

// Lang ... languages
type Lang struct {
	Ja Content `json:"ja"`
	En Content `json:"en"`
}

var lang Lang

// Index ... GET index.html
func Index(ctx *gin.Context) {

	labeljson, err := ioutil.ReadFile("./lang/index.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(labeljson, &lang)

	fmt.Println(lang.Ja.Label1)
	fmt.Println(lang.En.Label1)

	ctx.HTML(200, "index.html", gin.H{})
}
