package internal

type Seller interface {
	RegisterSeller(sellerID int, sellerName string) (err error)
	CreateAuction(auctionID int, lowestBidLimit int, highestBidLimit int, participationCost int, sellerID int) (err error)
	CloseAuction(auctionID int, sellerID int) (winningBidderName string, err error)
	GetProfit(sellerID int, auctionID int) (profit float64, err error)
}
