package buyer

import "flipkart/internal"

type buyerInfo struct {
	name   string
	mobile string
}

type buyerDetails struct {
	auction internal.BuyerActions
	details map[int]buyerInfo
}
