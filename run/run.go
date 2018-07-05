package run

import (
	"github.com/SunMaybo/go-jewel/context"
	"tokenup-pay-gateway-vendor-server-demo/config"
	"tokenup-pay-gateway-vendor-server-demo/controller"
	"tokenup-pay-gateway-vendor-server-demo/sdk"
	"tokenup-pay-gateway-vendor-server-demo/entity"
	"tokenup-pay-gateway-vendor-server-demo/model"
	"tokenup-pay-gateway-vendor-server-demo/router"
)

func Run() *context.Boot {
	boot := context.NewInstance()
	boot.AddApplyCfg(&config.SystemConfig{})
	boot.AddApply(&controller.NotifyController{},
		&controller.ThirdPartyController{},
		&entity.TransactionDb{},
		&router.Router{},
	)
	boot.AddFun(func() {
		context.Services.Db().MysqlDb.AutoMigrate(&model.Transaction{})
		SystemConfig := boot.GetInject().Service(&config.SystemConfig{}).(config.SystemConfig)
		controller.Client = sdk.ConfigBuilder(sdk.Config{
			Url:                    SystemConfig.Url,
			AppId:                  SystemConfig.AppId,
			AppKey:                 SystemConfig.AppKey,
			NotifyUrl:              SystemConfig.NotifyUrl,
			PrivateKey:             SystemConfig.PrivateKey,
			CallBackPartyPublicKey: SystemConfig.CallBackPartyPublicKey,
		})
	})
	return boot
}
