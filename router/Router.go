package router

import (
	"github.com/gin-gonic/gin"
	"tokenup-pay-gateway-vendor-server-demo/controller"
)

type Router struct {
	NotifyController     *controller.NotifyController     `inject:""`
	ThirdPartyController *controller.ThirdPartyController `inject:""`
}

func (router *Router) Router(context *gin.Engine) {
	context.POST("/test/notify", func(context *gin.Context) {
		router.NotifyController.Test(context)
	})
	context.GET("/verdor/balances", func(context *gin.Context) {
		router.ThirdPartyController.Balance(context)
	})
	context.POST("/verdor/recharge", func(context *gin.Context) {
		router.ThirdPartyController.Recharge(context)
	})
	context.POST("/verdor/withdraw", func(context *gin.Context) {
		router.ThirdPartyController.Withdraw(context)
	})
	context.GET("/verdor/transactions", func(context *gin.Context) {
		router.ThirdPartyController.Txs(context)
	})
}
