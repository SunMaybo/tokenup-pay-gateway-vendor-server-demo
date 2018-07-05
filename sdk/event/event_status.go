package event

type ConfirmEvent int

const (
	BindWithdrawAddressEvent  ConfirmEvent = iota
	UserWithdrawEvent
	ProxyTransactionEvent
	ApplyRechargeAddressEvent
	RechargeEvent
)

func (cs ConfirmEvent) String() string {
	switch cs {
	case BindWithdrawAddressEvent:
		return "USER_WITHDRAW_ADDRESS_BIND_RECEIVED"
	case UserWithdrawEvent:
		return "USER_WITHDRAW_DATA_RECEIVED"
	case ProxyTransactionEvent:
		return "PROXY_TRANSACTION_ORDER_DONE"
	case ApplyRechargeAddressEvent:
		return "USER_RECHARGE_ADDRESS_APPLY_DATA_RECEIVED"
	case RechargeEvent:
		return "USER_RECHARGE_DATA_RECEIVED"
	default:
		return ""
	}
}
