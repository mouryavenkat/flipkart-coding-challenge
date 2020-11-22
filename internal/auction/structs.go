package auction

type auctionInfo struct {
	participationCost int
	lowestBidLimit    int
	highestBidLimit   int
	// key is participantID, value is bidAmount
	participants map[int]int
	sellerID     int
	// open or close
	auctionStatus    string
	participantCount int
	winner           int
	winnerExists     bool
	finalBidAmount   int
}

type auctionImp struct {
	auctionDetails map[int]*auctionInfo
}
