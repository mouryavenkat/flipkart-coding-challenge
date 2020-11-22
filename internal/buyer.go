package internal

type Buyer interface {
	RegisterBuyer(buyerID int, buyerName string, mobile string) (err error)
	CreateBid(buyerID int, auctionID int, bidAmount int) (err error)
	WithdrawBid(buyerID int, auctionID int) (err error)
	UpdateBid(buyerID int, auctionID int, bidAmount int) (err error)
}
