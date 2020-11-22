package internal

type SellerActions interface {
	AddNewAuction(auctionID int, sellerID int, participationCost int, lowestBidLimit int, highestBidLimit int) (err error)
	CloseAuction(auctionID int, sellerID int) (biddingDetails [][2]int, err error)
	UpdateBidWinner(auctionID int, participantID int, finalBidAmount int, winnerExists bool) (err error)
	GetClosedAuctionDetails(auctionID int, sellerID int) (AuctionCalculations, error)
}

type BuyerActions interface {
	AddBidderToAuction(auctionID int, participantID int, bidAmount int) (err error)
	UpdateBidderAmountInAuction(auctionID int, participantID int, bidAmount int) (err error)
	WithdrawBidderFromAuction(auctionID int, participantID int) (err error)
}

type Auction interface {
	BuyerActions
	SellerActions
}

type AuctionCalculations struct {
	FinalBidAmount    int
	ParticipantCount  int
	ParticipationCost int
	LowestBidLimit    int
	HighestBidLimit   int
	WinnerExists      bool
}
