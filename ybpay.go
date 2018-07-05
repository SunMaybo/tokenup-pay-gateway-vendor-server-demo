package main

import (
	"tokenup-pay-gateway-vendor-server-demo/run"
	"github.com/gin-gonic/gin"
	"tokenup-pay-gateway-vendor-server-demo/router"
)

func main() {
	boot := run.Run().Start()
	boot.BindHttp(func(engine *gin.Engine) {
		router := boot.GetInject().Service(&router.Router{}).(router.Router)
		router.Router(engine)
	})
}
