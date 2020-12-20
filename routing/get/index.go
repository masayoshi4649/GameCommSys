package get

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Index ... GET index.html
func Index(ctx *gin.Context) {

	query := "SELECT * FROM mst_label WHERE method = 1 AND uri = 'index' AND lang = 2;"
	label := Selectsqlite3(query)

	var labelmap map[string]interface{}
	json.Unmarshal([]byte(label), &labelmap)

	fmt.Println(labelmap["label1"])
	for i := 0; i < len(labelmap); i++ {
		pointer := i + 1
		pointerstr := "label" + strconv.Itoa(pointer)
		fmt.Println(i, labelmap[pointerstr])
	}

	ctx.HTML(200, "index.html", labelmap)
}
