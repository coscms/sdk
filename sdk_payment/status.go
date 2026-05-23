package sdk_payment

// TradeStatus represents a payment notification status.
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
	// TradeStatusDeleted 已删除
	TradeStatusDeleted TradeStatus = "deleted"
)

// IsSuccess returns true if the status is success.
func (s TradeStatus) IsSuccess() bool {
	return s == TradeStatusSuccess
}

// IsFailure returns true if the status is failure.
func (s TradeStatus) IsFailure() bool {
	return s == TradeStatusFailure
}

// IsCancelled returns true if the status is cancelled.
func (s TradeStatus) IsCancelled() bool {
	return s == TradeStatusCancelled
}

// -----

// IsExpired returns true if the status is expired.
func (s TradeStatus) IsExpired() bool {
	return s == TradeStatusExpired
}

// IsDeleted returns true if the status is deleted.
func (s TradeStatus) IsDeleted() bool {
	return s == TradeStatusDeleted
}
