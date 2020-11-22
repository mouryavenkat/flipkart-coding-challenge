package seller

import (
	"errors"
	"flipkart/internal"
	"fmt"
)

func NewSeller(auctions internal.SellerActions) internal.Seller {
	return &sellerDetails{
		auction: auctions,
		details: make(map[int]sellerInfo),
	}
}

func (s *sellerDetails) RegisterSeller(sellerID int, sellerName string) (err error) {
	if _, ok := s.details[sellerID]; ok {
		return errors.New("seller already exists")
	}
	s.details[sellerID] = sellerInfo{
		name: sellerName,
	}
	return nil
}

func (s *sellerDetails) CreateAuction(auctionID int, lowestBidLimit int, highestBidLimit int, participationCost int, sellerID int) (err error) {
	if _, ok := s.details[sellerID]; !ok {
		return errors.New("seller doesn't exists")
	}
	return s.auction.AddNewAuction(auctionID, sellerID, participationCost, lowestBidLimit, highestBidLimit)
}

func (s *sellerDetails) CloseAuction(auctionID int, sellerID int) (winningBidderName string, err error) {
	if _, ok := s.details[sellerID]; !ok {
		return "", errors.New("seller doesn't exists")
	}
	biddingDetails, err := s.auction.CloseAuction(auctionID, sellerID)
	if err != nil {
		return "", errors.New("Error while closing auction " + err.Error())
	}

	biddingCounter := map[int][]int{}
	for _, biddingInfo := range biddingDetails {
		biddingCounter[biddingInfo[1]] = append(biddingCounter[biddingInfo[1]], biddingInfo[0])
	}
	var bidderID int
	var bidAmount int
	var max = -1
	for _, biddingInfo := range biddingDetails {
		if biddingInfo[1] > max && len(biddingCounter[biddingInfo[1]]) == 1 {
			bidderID = biddingInfo[0]
			bidAmount = biddingInfo[1]
			max = bidAmount
		}
	}
	if max == -1 {
		s.auction.UpdateBidWinner(auctionID, 0, 0, false)
	} else {
		fmt.Println("Winner Details", bidderID, bidAmount)
		s.auction.UpdateBidWinner(auctionID, bidderID, bidAmount, true)
	}
	return "", nil
}

func (s *sellerDetails) GetProfit(sellerID int, auctionID int) (profit float64, err error) {
	if _, ok := s.details[sellerID]; !ok {
		return 0, errors.New("seller doesn't exists")
	}
	auctionDetails, err := s.auction.GetClosedAuctionDetails(auctionID, sellerID)
	if err != nil {
		return 0, err
	}
	fmt.Println(auctionDetails)
	sellerBaseProfit := float64(auctionDetails.ParticipantCount*auctionDetails.ParticipationCost) * .2
	if auctionDetails.WinnerExists {
		return float64(auctionDetails.FinalBidAmount) + sellerBaseProfit -
			float64((auctionDetails.LowestBidLimit+auctionDetails.HighestBidLimit)/2), nil
	}
	return sellerBaseProfit, nil
}
