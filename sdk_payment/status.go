package sdk_payment

// TradeStatus 通知交易状态
type TradeStatus string

const (
	// TradeStatusSuccess 成功
	TradeStatusSuccess TradeStatus = "success"
	// TradeStatusFailure 失败
	TradeStatusFailure TradeStatus = "failure"
	// TradeStatusCancelled 被取消
	TradeStatusCancelled TradeStatus = "cancelled"

	// -----

	// TradeStatusExpired 已过期
	TradeStatusExpired TradeStatus = "expired"
	// TradeStatusExpired 已删除
	TradeStatusDeleted TradeStatus = "deleted"
)

func (s TradeStatus) IsSuccess() bool {
	return s == TradeStatusSuccess
}

func (s TradeStatus) IsFailure() bool {
	return s == TradeStatusFailure
}

func (s TradeStatus) IsCancelled() bool {
	return s == TradeStatusCancelled
}

// -----

func (s TradeStatus) IsExpired() bool {
	return s == TradeStatusExpired
}

func (s TradeStatus) IsDeleted() bool {
	return s == TradeStatusDeleted
}
