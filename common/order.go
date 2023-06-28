package common

const (
	// -1 取消
	OrderTypeCancelled = iota - 1
	_
	// 1 未支付
	OrderTypePending
	// 2 已支付
	OrderTypePaid
	// 3 已完成
	OrderTypeDone
)

func OrderTypeToString(OrderType int) string {
	switch OrderType {
	case OrderTypeCancelled:
		return "CANCELLED"
	case OrderTypePending:
		return "PENDING"
	case OrderTypePaid:
		return "PAID"
	case OrderTypeDone:
		return "DONE"
	default:
		return "UNDEFINED"
	}
}
