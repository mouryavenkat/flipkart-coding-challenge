package seller

import "flipkart/internal"

type sellerInfo struct {
	name string
}

type sellerDetails struct {
	auction internal.SellerActions
	details map[int]sellerInfo
}
