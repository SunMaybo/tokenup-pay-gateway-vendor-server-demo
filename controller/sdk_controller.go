package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/cihub/seelog"
	"fmt"
	"net/http"
	"tokenup-pay-gateway-vendor-server-demo/entity"
	"tokenup-pay-gateway-vendor-server-demo/sdk/event"
	"tokenup-pay-gateway-vendor-server-demo/model"
	"tokenup-pay-gateway-vendor-server-demo/sdk"
)

type NotifyController struct {
	TransactionDb *entity.TransactionDb `inject:""`
}

func (nc *NotifyController) Test(context *gin.Context) {
	var confirm sdk.Confirm
	context.BindJSON(&confirm)
	received, err := Client.ValidReceivedCallBack(&confirm, "ok")
	if err != nil {
		seelog.Error(err)
	} else {
		fmt.Printf("%+v", received)
		received := confirm.Received
		if confirm.Event == event.UserWithdrawEvent.String() || confirm.Event == event.RechargeEvent.String() {
			status := confirm.Result["status"]
			orderId := received["order_id"].(string)
			if status == sdk.BlockConfirm {
				nc.TransactionDb.UpdateStatus(orderId, model.Confirms.String())
			}
			context.JSON(http.StatusOK, received)
		}
	}

}
