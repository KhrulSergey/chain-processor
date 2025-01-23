package types

type DeliveryGuarantee string

const (
	AtMostOnceDeliveryGuarantee  DeliveryGuarantee = "AT_MOST_ONCE"
	AtLeastOnceDeliveryGuarantee DeliveryGuarantee = "AT_LEAST_ONCE"
	ExactlyOnceDeliveryGuarantee DeliveryGuarantee = "EXACTLY_ONCE"
)
