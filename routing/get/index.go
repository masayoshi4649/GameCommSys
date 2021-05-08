package get

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Index ... GET index.html
func Index(ctx *gin.Context) {
	var uri string = "index"
	lang := 1
	var label string = Datalabel(uri, lang)

	var labelmap map[string]interface{}
	json.Unmarshal([]byte(label), &labelmap)
	labelmap["test"] = "てすと"

	ctx.HTML(200, "index.html", labelmap)
}
