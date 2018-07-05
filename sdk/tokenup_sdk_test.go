package sdk

import (
	"testing"
	"github.com/satori/go.uuid"
	"log"
	"fmt"
)

var client *Client

func init() {
	client = ConfigBuilder(Config{
		Url:                    "http://api.staging.paygateway.tokenup.io",
		AppId:                  "cf9a6bff34289f7722ce59720710d351b7491d05",
		AppKey:                 "9b3339ab9b9c9b161d23548bf4959cb50bd8ad2b",
		NotifyUrl:              "http://127.0.0.1:8597/test/notify",
		PrivateKey:             "MIIEowIBAAKCAQEAxJ1ULffXG3H4v6h2FnHJPrmcljS39JkC5ffjVBojIi1L82vxswmJSJajMOxO/jsFdpvZa9EpxuC0TA30BfHeRNsXXOjE1MVbpDp99sgj5+KZKkjoeE/dWd7/jZgb+4yMwgZYlGt716TJ73PgMu1eq5+tw7QOVaAs5vc0EGLKDd0YrwxI4aaRrnIpg34g/p8WnTcfaexUikHWRlBoOgypVexEPdeA8NC7F78h5pykWpk4Whko8Hxs1JebGX8YOtepCpWv6llwQZOU9RvDdDkrFe/velVrojBmb7IPNJCyNKsxwJ2PtpP+NLMVvaXqsGtjcPtcC2zo0y8WUCTavDsW3wIDAQABAoIBAD+m85bCMvimqDJcNobDpbRR4Pjb7mYYl1CeNRGIOLGa2oje/GvK/Y/rfL+c8WHq97TTdcsq9wx0uMoahlLaX+wIxgKFNRvxHN8JNLiNSNqMiKug2OoCaRXsVO2hPgXtFbDG3yyFs503s0x7Ri0WndyQIHBIPY/JAGBxzYA0i1d7F2qSoG1i56HoRQ0CtQ+sbf2TPAXC8y+nkSI/4WJDYjqPNWf9Ak5IkPwn21C3YcIgVzxLZBYc4qBdxWhhXIC8GoYA4h3UMtKX7D27FyJDEUoD7HEUJod3puygPsUrCS7AQN9B0v1omQCLfBe41YJBZEfQgW9nOd6o4H3eRi/TSUkCgYEA/T4kGznRXVNQuI4C4C+n8AF8lhqBCUDlGQIbwjP2wrOkycAyY5F0S/rsHyagyki4kkUeikcHVyK5wVSwM0ew5jIJX9a6jpgFlKmV1wquICRZe9pRV5KElfnAFfyikywV7GJRpEv7ScKdhgqWHhq/QabK9Eh3K2TJdPD4mr9FfAUCgYEAxsFZXhmALTcdi2RFw73sL+MI6JgICp5hfuy3Zd5CWGpAoCI6wE40fjXTlY9ECSgQugHJbETjqKpL6uZ/qN+AQADLcIl/6a0wh4ha2ahkNVeHb7ELslU3Wbrc7Q0ZUxmaSGcbtrjlfJRMBMV5mfGEt/hkP8Cio7hseEdbmHD3YJMCgYAdbdtETrPF5Ki8ycQLyX36pjGUQAA+0wvMnDIdn2xNtBKyX2N7rquVKNPHyvVkjI7mcKHb7+Uqex6bGPxg+TPVjHsKaCnF6GS9ofeHxfX2RkMf4X8SbjR6OUvZQkKiV700eziBn1LUf4lOymwnk3QmbPuo58LxiAThUh+R3Ch3AQKBgCdEg/eHaj+EqB2mDfKCT2uWm2f4wX33lKOS+RjzNIBrXaFFof3kdZKJ5+eginyUodleCQGPCruECcO7DnW60ofSoF73i4ILaY8dbXWbQ1EWnfd/LyRomarstEFRWTOF12l+lYcgOJbIZcx7h27WvLXsKUI/OOLHyQZqcrpHd1hpAoGBAIpKEyYejSeggIZEt2G4dSdf3QG9Iq49jgANJyFnYU+zCrbdIPoBgRlayDN7+yOkfkIju2HPckxoS3OIXxzO6uW6nftPW20G/h+ii33ODENzFmhrz6Q9khk1I0HJgEoVLuqbPUcgdGCHajKQKA17OcKa8Ke0oSV/duYT65fawdhF",
		CallBackPartyPublicKey: "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDawVhV65uzU5oLC98O9+em6xEqzHz/9JW3udZ8M1ns5iRFV8p8iga7e2hNHiWP7tu6SnPqvoyYEFRwtv6U9/GkpIO+6gBKjjEBDVk2V1GouMTVhnlGc1EdcBSORuWRdDuFTa9FjB5U863rAOaxwtvrnPghJdQQzL67JcwgWXjgLwIDAQAB",
	})
}

/**
 充值地址申请测试
 */
func TestClient_AddressApply(t *testing.T) {
	u1, _ := uuid.NewV4()
	result, err := client.AddressApply(RechargeAddress{
		UserId: u1.String(),
		Extras: "tokenup-address_apply-test",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Sprintf("%+v", result)
}

/**
充值接口测试
 */
func TestClient_ReCharge(t *testing.T) {
	u1, _ := uuid.NewV4()
	result, err := client.ReCharge(Charge{
		UserId:  "90fdc761-9633-41a4-80dd-c48ed3ea0ed0",
		OrderId: u1.String(),
		Amount:  "0.25",
		Extras:  "tokenup-recharge-test",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Sprintf("%+v", result)
}

/**
提现接口测试
 */
func TestClient_Withdraw(t *testing.T) {
	u1, _ := uuid.NewV4()
	result, err := client.Withdraw(Withdraw{
		OrderId: u1.String(),
		Amount:  "0.16",
		UserId:  "90fdc761-9633-41a4-80dd-c48ed3ea0ed0",
		Extras:  "tokenup-withdraw-test",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Sprintf("%+v", result)
}

/**
提现地址bind测试
 */
func TestClient_UserBind(t *testing.T) {
	result, err := client.UserBind(User{
		UserId:  "90fdc761-9633-41a4-80dd-c48ed3ea0ed0",
		Address: "0x14f96915220ce4ca498c5ec00f4d0904515e1fbd",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Sprintf("%+v", result)
}
