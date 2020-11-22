package buyer

import (
	"errors"
	"flipkart/internal"
)

func NewBuyer(auctionGetter internal.BuyerActions) internal.Buyer {
	return &buyerDetails{
		auction: auctionGetter,
		details: make(map[int]buyerInfo),
	}
}

func (b *buyerDetails) RegisterBuyer(buyerID int, buyerName string, mobile string) (err error) {
	if _, ok := b.details[buyerID]; ok {
		return errors.New("buyer already exists")
	}
	b.details[buyerID] = buyerInfo{
		name:   buyerName,
		mobile: mobile,
	}

	return nil
}

func (b *buyerDetails) CreateBid(buyerID int, auctionID int, bidAmount int) (err error) {
	if _, ok := b.details[buyerID]; !ok {
		return errors.New("buyer doesn't exists")
	}
	return b.auction.AddBidderToAuction(auctionID, buyerID, bidAmount)
}

func (b *buyerDetails) WithdrawBid(buyerID int, auctionID int) (err error) {
	if _, ok := b.details[buyerID]; !ok {
		return errors.New("buyer doesn't exists")
	}
	return b.auction.WithdrawBidderFromAuction(auctionID, buyerID)
}

func (b *buyerDetails) UpdateBid(buyerID int, auctionID int, bidAmount int) (err error) {
	if _, ok := b.details[buyerID]; !ok {
		return errors.New("buyer doesn't exists")
	}
	return b.auction.UpdateBidderAmountInAuction(auctionID, buyerID, bidAmount)
}
