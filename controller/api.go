package controller

import (
	"tokenup-pay-gateway-vendor-server-demo/entity"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"tokenup-pay-gateway-vendor-server-demo/sdk"
	"tokenup-pay-gateway-vendor-server-demo/model"
	"time"
	"github.com/cihub/seelog"
	"crypto/md5"
	"encoding/hex"
	"encoding/base64"
	"io"
	"crypto/rand"
)

var Client *sdk.Client

type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ThirdPartyController struct {
	TransactionDb *entity.TransactionDb `inject:""`
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func (tc *ThirdPartyController) Balance(c *gin.Context) {
	userId := c.Query("user_id")
	amount := tc.TransactionDb.FindBalance(userId)
	amountMap := make(map[string]string)
	amountMap["amount"] = amount
	c.JSON(http.StatusOK, Response{
		Data: amountMap,
		Status: Status{
			Code:    200,
			Message: "api invoke success",
		},
	})
}
func (tc *ThirdPartyController) Txs(c *gin.Context) {
	userId := c.Query("user_id")
	pageIndexStr := c.DefaultQuery("page_index", "0")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	pageIndex, _ := strconv.Atoi(pageIndexStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	txs := tc.TransactionDb.FindPage(userId, pageIndex, pageSize)
	c.JSON(http.StatusOK, Response{
		Data: txs,
		Status: Status{
			Code:    200,
			Message: "api invoke success",
		},
	})
}
func (tc *ThirdPartyController) Withdraw(c *gin.Context) {
	var req map[string]string
	c.BindJSON(&req)
	orderId := UniqueId()
	userId := req["user_id"]
	amount, _ := strconv.ParseFloat(req["amount"], 10)
	tx := model.Transaction{
		UserId:       userId,
		OrderID:      orderId,
		Timestamp:    time.Now().Unix(),
		Status:       model.Pending.String(),
		Type:         1,
		Amount:       amount,
		CurrencyType: "ETH",
		Address:      "0x14f96915220ce4ca498c5ec00f4d0904515e1fbd",
	}
	tc.TransactionDb.Save(tx)
	result, err := Client.Withdraw(sdk.Withdraw{
		OrderId: orderId,
		Amount:  req["amount"],
		Extras:  "tokenup-test",
		UserId:  userId,
	})
	if err != nil {
		seelog.Error(err)
		c.JSON(http.StatusOK, Response{
			Status: Status{
				Code:    -201,
				Message: err.Error(),
			},
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Data: result,
			Status: Status{
				Code:    200,
				Message: "api invoke success",
			},
		})
	}
}
func (tc *ThirdPartyController) Recharge(c *gin.Context) {
	var req map[string]string
	c.BindJSON(&req)
	orderId := UniqueId()
	userId := req["user_id"]
	amount, _ := strconv.ParseFloat(req["amount"], 10)
	tx := model.Transaction{
		UserId:       userId,
		OrderID:      orderId,
		Timestamp:    time.Now().Unix(),
		Status:       model.Pending.String(),
		Type:         0,
		Amount:       amount,
		CurrencyType: "ETH",
		Address:      "0x949e36c7a8600ade99ae50f68a1a2d91a025ea57",
	}
	tc.TransactionDb.Save(tx)
	result, err := Client.ReCharge(sdk.Charge{
		OrderId: orderId,
		Amount:  req["amount"],
		Extras:  "tokenup-test",
		UserId:  userId,
	})
	if err != nil {
		seelog.Error(err)
		c.JSON(http.StatusOK, Response{
			Status: Status{
				Code:    -201,
				Message: err.Error(),
			},
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Data: result,
			Status: Status{
				Code:    200,
				Message: "api invoke success",
			},
		})
	}
}
