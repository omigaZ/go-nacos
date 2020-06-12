package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"go-nacos/bus"
	"go-nacos/bus/data"
	"log"
)

func init()  {

}

func main()  {
	defer data.XormOp().Close()
	defer data.CacheOp().Close()
	router := fasthttprouter.New()
	router.POST("/bus/hello", bus.Do)
	err := fasthttp.ListenAndServe(":8800", router.Handler)
	if err != nil {
		log.Fatal(err)
	}
}
