package sdk

import (
	"net/url"
	"strconv"
	"reflect"
	"strings"
	"sort"
	"errors"
	"encoding/base64"
	"crypto/x509"
	"crypto/sha256"
	"encoding/hex"
	"crypto/rsa"
	"crypto"
	"time"
	"crypto/rand"
	rand2 "math/rand"
	"github.com/SunMaybo/jewel-template/template/rest"
	"fmt"
)

type Authorize struct {
	AppId                  string
	AppKey                 string
	NotifyUrl              string
	PrivateKey             string
	CallBackPartyPublicKey string
}
type WithdrawConfirm struct {
	Nonce     string                 `json:"nonce" sign:"nonce"`
	Received  map[string]interface{} `json:"received" sign:"received"`
	Event     string                 `json:"event" sign:"event"`
	Signature string                 `json:"signature"`
	Result struct {
		TransactionHash string `json:"transaction_hash" sign:"transaction_hash"`
		Status          string `json:"status" sign:"status"`
		Message         string `json:"message" sign:"message"`
	} `json:"result" sign:"result"`
}
type Confirm struct {
	Nonce     string                 `json:"nonce" sign:"nonce"`
	Received  map[string]interface{} `json:"received" sign:"received"`
	Event     string                 `json:"event" sign:"event"`
	Signature string                 `json:"signature"`
	Result    map[string]interface{} `json:"result" sign:"result"`
}
type WithdrawAddressBindConfirm struct {
	Nonce     string                 `json:"nonce" sign:"nonce"`
	Received  map[string]interface{} `json:"received" sign:"received"`
	Event     string                 `json:"event" sign:"event"`
	Signature string                 `json:"signature"`
	Result struct {
		Status  string `json:"status" sign:"status"`
		Message string `json:"message" sign:"message"`
	} `json:"result" sign:"result"`
}
type RechargeAddressApplyConfirm struct {
	Nonce     string                 `json:"nonce" sign:"nonce"`
	Received  map[string]interface{} `json:"received" sign:"received"`
	Event     string                 `json:"event" sign:"event"`
	Signature string                 `json:"signature"`
	Result struct {
		Address string `json:"address" sign:"address"`
		Status  string `json:"status" sign:"status"`
		Message string `json:"message" sign:"message"`
	} `json:"result" sign:"result"`
}
type RechargeConfirm struct {
	Nonce     string                 `json:"nonce" sign:"nonce"`
	Received  map[string]interface{} `json:"received" sign:"received"`
	Event     string                 `json:"event" sign:"event"`
	Signature string                 `json:"signature"`
	Result struct {
		Status  string `json:"status" sign:"status"`
		Message string `json:"message" sign:"message"`
	} `json:"result" sign:"result"`
}
type ReceivedConfirm struct {
	Message   string `json:"message" sign:"message"`
	Nonce     string `json:"nonce" sign:"nonce"`
	AppKey    string `json:"-" sign:"app_key"`
	Signature string `json:"signature"`
}
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Result struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}
type Client struct {
	Url       string
	Authorize Authorize
	Rest      *rest.RestTemplate
}

type Charge struct {
	OrderId string
	UserId  string
	Amount  string
	Extras  string
}
type ChargeSafe struct {
	OrderId   string `json:"order_id" sign:"order_id"`
	Nonce     string `json:"nonce" sign:"nonce"`
	UserId    string `json:"user_id" sign:"user_id"`
	Amount    string `json:"amount" sign:"amount"`
	Extras    string `json:"extras" sign:"extras"`
	AppId     string `json:"app_id" sign:"app_id"`
	Timestamp int64  `json:"timestamp" sign:"timestamp"`
	NotifyUrl string `json:"notify_url" sign:"notify_url"`
	Signature string `json:"signature"`
	AppKey    string `sign:"app_key" json:"-"`
}

type Withdraw struct {
	OrderId string
	Extras  string
	UserId  string
	Amount  string
}

type WithdrawSafe struct {
	OrderId   string `json:"order_id" sign:"order_id"`
	Nonce     string `json:"nonce" sign:"nonce"`
	Extras    string `json:"extras" sign:"extras"`
	UserId    string `json:"user_id" sign:"user_id"`
	AppId     string `json:"app_id" sign:"app_id"`
	Timestamp int64  `json:"timestamp" sign:"timestamp"`
	NotifyUrl string `json:"notify_url" sign:"notify_url"`
	Signature string `json:"signature"`
	Amount    string `json:"amount" sign:"amount"`
	AppKey    string `sign:"app_key" json:"-"`
}
type User struct {
	Address string
	UserId  string
	Extras  string
}
type UserSafe struct {
	AppId     string `json:"app_id" sign:"app_id"`
	Timestamp int64  `json:"timestamp" sign:"timestamp"`
	NotifyUrl string `json:"notify_url" sign:"notify_url"`
	Signature string `json:"signature"`
	Address   string `json:"address" sign:"address"`
	Nonce     string `json:"nonce" sign:"nonce"`
	AppKey    string `sign:"app_key" json:"-"`
	UserId    string `json:"user_id" sign:"user_id"`
	Extras    string `json:"extras" sign:"extras"`
}

type RechargeAddress struct {
	UserId string
	Extras string
}

type RechargeAddressSafe struct {
	UserId    string `json:"user_id" sign:"user_id"`
	Extras    string `json:"extras" sign:"extras"`
	AppId     string `json:"app_id" sign:"app_id"`
	AppKey    string `sign:"app_key" json:"-"`
	Nonce     string `json:"nonce" sign:"nonce"`
	NotifyUrl string `json:"notify_url" sign:"notify_url"`
	Timestamp int64  `json:"timestamp" sign:"timestamp"`
	Signature string `json:"signature"`
}

type Config struct {
	Url                    string
	AppId                  string
	AppKey                 string
	NotifyUrl              string
	PrivateKey             string
	CallBackPartyPublicKey string
	RestConfig             *rest.ClientConfig
}

func ConfigBuilder(cfg Config) *Client {
	var restClient *rest.RestTemplate
	if cfg.RestConfig == nil {
		restClient = rest.Default()
	} else {
		restClient = rest.Config(*cfg.RestConfig)
	}
	return &Client{
		Url: cfg.Url,
		Authorize: Authorize{
			AppId:                  cfg.AppId,
			AppKey:                 cfg.AppKey,
			NotifyUrl:              cfg.NotifyUrl,
			PrivateKey:             cfg.PrivateKey,
			CallBackPartyPublicKey: cfg.CallBackPartyPublicKey,
		},
		Rest: restClient,
	}
}

func (client *Client) ValidReceivedCallBack(confirm interface{}, message string) (ReceivedConfirm, error) {
	signature := reflect.ValueOf(confirm).Elem().FieldByName("Signature").Interface().(string)
	nonce := reflect.ValueOf(confirm).Elem().FieldByName("Nonce").Interface().(string)
	reflect.ValueOf(confirm).Elem().FieldByName("Received").SetMapIndex(reflect.ValueOf("app_key"), reflect.ValueOf(client.Authorize.AppKey))
	timeValue := reflect.ValueOf(confirm).Elem().FieldByName("Received").MapIndex(reflect.ValueOf("timestamp"))
	var time float64
	fmt.Sscanf(fmt.Sprint(timeValue.Interface()), "%e", &time)
	reflect.ValueOf(confirm).Elem().FieldByName("Received").SetMapIndex(reflect.ValueOf("timestamp"), reflect.ValueOf(uint64(time)))
	fmt.Printf("sdk---:%s", EncodeString(confirm))
	err := RsaSignVerAndPublicHex([]byte(EncodeString(confirm)), signature, client.Authorize.CallBackPartyPublicKey)
	if err != nil {
		return ReceivedConfirm{}, err
	}
	rc := ReceivedConfirm{
		Nonce:   nonce,
		Message: message,
		AppKey:  client.Authorize.AppKey,
	}
	signReceived, _ := RsaSignAndPrivate([]byte(EncodeString(rc)), client.Authorize.PrivateKey)
	rc.Signature = signReceived
	return rc, nil
}
func (client *Client) Withdraw(withdraw Withdraw) (Result, error) {
	var withdrawSafe WithdrawSafe
	StructCopy(&withdraw, &withdrawSafe)
	return client.callPost(client.Url+"/vendor/withdraw", &withdrawSafe)
}
func (client *Client) UserBind(user User) (Result, error) {
	var userSafe UserSafe
	StructCopy(&user, &userSafe)
	return client.callPost(client.Url+"/vendor/user/bind", &userSafe)
}
func (client *Client) ReCharge(charge Charge) (Result, error) {
	var chargeSafe ChargeSafe
	StructCopy(&charge, &chargeSafe)
	return client.callPost(client.Url+"/vendor/charge", &chargeSafe)
}
func (client *Client) AddressApply(rechargeAddress RechargeAddress) (Result, error) {
	var rechargeAddressSafe RechargeAddressSafe
	StructCopy(&rechargeAddress, &rechargeAddressSafe)
	return client.callPost(client.Url+"/vendor/charge/address/apply", &rechargeAddressSafe)
}

func (client *Client) callPost(url string, data interface{}) (Result, error) {
	value := reflect.ValueOf(data).Elem()
	value.FieldByName("Timestamp").Set(reflect.ValueOf(time.Now().Unix()))
	value.FieldByName("AppKey").Set(reflect.ValueOf(client.Authorize.AppKey))
	value.FieldByName("NotifyUrl").Set(reflect.ValueOf(client.Authorize.NotifyUrl))
	value.FieldByName("AppId").Set(reflect.ValueOf(client.Authorize.AppId))
	value.FieldByName("Nonce").Set(reflect.ValueOf(strconv.FormatInt(RandInt64(), 10)))
	fmt.Println(EncodeString(data))
	fmt.Printf("%+v", data)
	var Signature string
	var err error
	Signature, err = RsaSignAndPrivate([]byte(EncodeString(data)), client.Authorize.PrivateKey)
	if err != nil {
		return Result{}, err
	}
	signValue := value.FieldByName("Signature")
	signValue.Set(reflect.ValueOf(Signature))
	var result Result
	err = client.Rest.PostForObject(url, data, &result)
	if err != nil {
		return Result{}, err
	}
	return result, nil
}

func DeepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}

var channel = make(chan int64, 32)

func init() {
	go func() {
		var old int64
		for {
			o := rand2.New(rand2.NewSource(time.Now().UnixNano())).Int63()
			if old != o {
				old = o
				select {
				case channel <- o:
				}
			}
		}
	}()
}
func RandInt64() (r int64) {
	select {
	case rand := <-channel:
		r = rand
	}
	return
}

func StructCopy(srcPtr interface{}, desPtr interface{}) {
	srcv := reflect.ValueOf(srcPtr)
	dstv := reflect.ValueOf(desPtr)
	srct := reflect.TypeOf(srcPtr)
	dstt := reflect.TypeOf(desPtr)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := DeepFields(reflect.ValueOf(srcPtr).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}

type Dec []string

func (p Dec) Len() int           { return len(p) }
func (p Dec) Less(i, j int) bool { return p[i] < p[j] }
func (p Dec) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func EncodeString(o interface{}) string {
	dec := encodeField("", o)
	sort.Stable(dec)
	return strings.Join(dec, "&")
}

func encodeField(prefix string, o interface{}) Dec {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	valStr, err := getValueString(v)
	var dec Dec
	if err == nil {
		values := url.Values{}
		values.Set("url", valStr)
		return append(dec, prefix+"="+strings.Split(values.Encode(), "=")[1])
	}
	if prefix != "" {
		prefix += "."
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		prefix2 := t.Field(i).Tag.Get("sign")
		value := v.Field(i)
		if prefix2 == "" {
			continue
		}
		arrs := strings.Split(prefix2, ",")
		if len(arrs) >= 1 {
			prefix2 = arrs[0]
		}
		if isArray(value) {
			for j := 0; j < value.Len(); j++ {
				dec = append(dec, encodeField(prefix+prefix2+"[]", value.Index(j).Interface())...)
			}
		} else if value.Kind() == reflect.Map {
			dataMap := value.Interface().(map[string]interface{})
			for k, v := range dataMap {
				dec = append(dec, encodeField(prefix+prefix2+"."+k, v)...)
			}

		} else {
			dec = append(dec, encodeField(prefix+prefix2, value.Interface())...)
		}

	}
	return dec
}

func getValueString(value reflect.Value) (string, error) {
	if value.Type().String() == "string" {
		return value.Interface().(string), nil
	}
	if value.Type().String() == "bool" {
		val := value.Interface().(bool)
		if val {
			return "true", nil
		} else {
			return "false", nil
		}
	}
	if value.Type().String() == "int8" {
		val := value.Interface().(int8)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "int16" {
		val := value.Interface().(int16)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "int32" {
		val := value.Interface().(int32)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "int64" {
		val := value.Interface().(int64)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "uint8" {
		val := value.Interface().(uint8)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "uint16" {
		val := value.Interface().(uint16)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "uint32" {
		val := value.Interface().(uint32)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "int" {
		val := value.Interface().(int)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "uint" {
		val := value.Interface().(uint)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "uint64" {
		val := value.Interface().(uint64)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "float32" {
		val := value.Interface()
		var time float64
		fmt.Sscanf(fmt.Sprint(val), "%e", &time)
		return strconv.FormatFloat(time, 'E', -1, 32), nil
	}
	if value.Type().String() == "float64" {
		val := value.Interface()
		var time float64
		fmt.Sscanf(fmt.Sprint(val), "%e", &time)
		return strconv.FormatFloat(time, 'E', -1, 64), nil
	}
	if value.Type().String() == "uint32" {
		val := value.Interface().(uint32)
		return strconv.FormatUint(uint64(val), 10), nil
	}

	return "", errors.New("invalid field")

}

func isArray(value reflect.Value) bool {
	if strings.HasPrefix(value.Type().String(), "[]") {
		return true
	}
	return false
}

func RsaSignVerAndPublicHex(data []byte, signature, public string) error {
	signatureDecode, err := hex.DecodeString(signature)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256(data)
	buff, _ := base64.StdEncoding.DecodeString(public)
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(buff)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signatureDecode)
}
func RsaSignAndPrivate(data []byte, privateKey string) (string, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	buff, _ := base64.StdEncoding.DecodeString(privateKey)
	//获取私钥
	priv, err := x509.ParsePKCS1PrivateKey(buff)
	if err != nil {
		return "", err
	}
	sign, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
	return hex.EncodeToString(sign), err
}

type ConfirmStatus int

const (
	Pending                     ConfirmStatus = iota
	ApplyFailed
	UserValidFailed
	SignatureVerificationFailed
	TransactionAbandoned
	AddressValidFail
	SendTxFail
	TxFail
	SendFail
	SendSuccess
	BlockConfirm
	Success
)

func (cs ConfirmStatus) String() string {
	switch cs {
	case Pending:
		return "PENDING"
	case TransactionAbandoned:
		return "TRANSACTION_ABANDONED"
	case BlockConfirm:
		return "BLOCK_CONFIRM"
	case Success:
		return "SUCCESS"
	case ApplyFailed:
		return "APPLY_FAILED"
	case SignatureVerificationFailed:
		return "SIGNATURE_VERIFICATION_FAILED"
	case UserValidFailed:
		return "USER_VALID_FAILED"
	case SendSuccess:
		return "SEND_SUCCESS"
	case SendFail:
		return "SEND_FAIL"
	case SendTxFail:
		return "SEND_TX_FAIL"
	case TxFail:
		return "TX_FAIL"
	case AddressValidFail:
		return "ADDRESS_VALID_FAIL"
	default:
		return "PENDING"
	}
}
