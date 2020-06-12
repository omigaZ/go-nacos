package bus

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

var (
	body *interface{}
)


func Do(ctx *fasthttp.RequestCtx)  {

	bodyByte := ctx.PostBody()
	err := json.Unmarshal(bodyByte,&body)
	if err != nil {
		fmt.Println(err)
	}
	ctx.Write(bodyByte)

}
