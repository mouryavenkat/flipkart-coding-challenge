package auction

import (
	"errors"
	"flipkart/internal"
)

func NewAction() internal.Auction {
	return &auctionImp{
		auctionDetails: make(map[int]*auctionInfo),
	}
}

func (a *auctionImp) AddNewAuction(auctionID int, sellerID int, participationCost int, lowestBidLimit int, highestBidLimit int) (err error) {
	if _, ok := a.auctionDetails[auctionID]; ok {
		return errors.New("auctionID exists already")
	}
	a.auctionDetails[auctionID] = &auctionInfo{
		participationCost: participationCost,
		lowestBidLimit:    lowestBidLimit,
		highestBidLimit:   highestBidLimit,
		participants:      make(map[int]int),
		sellerID:          sellerID,
		auctionStatus:     "open",
		participantCount:  0,
	}
	return nil
}

func (a *auctionImp) CloseAuction(auctionID int, sellerID int) (biddingDetails [][2]int, err error) {
	if _, ok := a.auctionDetails[auctionID]; !ok {
		return nil, errors.New("invalid Auction ID")
	}
	if a.auctionDetails[auctionID].sellerID != sellerID {
		return nil, errors.New("only a valid seller can close the auction")
	}
	a.auctionDetails[auctionID].auctionStatus = "closed"
	bidDetails := make([][2]int, len(a.auctionDetails[auctionID].participants))
	var index int
	for bidderID, bidAmount := range a.auctionDetails[auctionID].participants {
		bidDetails[index] = [2]int{bidderID, bidAmount}
		index++
	}
	return bidDetails, nil
}

func (a *auctionImp) AddBidderToAuction(auctionID int, participantID int, bidAmount int) (err error) {
	if _, ok := a.auctionDetails[auctionID]; !ok {
		return errors.New("invalid Auction ID")
	}
	if _, ok := a.auctionDetails[auctionID].participants[participantID]; ok {
		return errors.New("participant already part of auction. Can't create another")
	}

	if bidAmount < a.auctionDetails[auctionID].lowestBidLimit || bidAmount > a.auctionDetails[auctionID].highestBidLimit {
		return errors.New("the bid doesnt fall in bid range area")
	}
	if a.auctionDetails[auctionID].auctionStatus == "closed" {
		return errors.New("you can't bid on a closed auction")
	}
	// TODO: lock this part as a txn
	a.auctionDetails[auctionID].participants[participantID] = bidAmount
	a.auctionDetails[auctionID].participantCount++

	return nil
}

func (a *auctionImp) UpdateBidWinner(auctionID int, participantID int, finalBidAmount int, winnerExists bool) (err error) {
	if _, ok := a.auctionDetails[auctionID]; !ok {
		return errors.New("invalid Auction ID")
	}
	if _, ok := a.auctionDetails[auctionID].participants[participantID]; !ok {
		return errors.New("bidder not present, please create one")
	}
	if a.auctionDetails[auctionID].auctionStatus != "closed" {
		return errors.New("auction not yet closed. Can't declare winner")
	}
	a.auctionDetails[auctionID].winner = participantID
	a.auctionDetails[auctionID].winnerExists = winnerExists
	a.auctionDetails[auctionID].finalBidAmount = finalBidAmount
	return nil
}

func (a *auctionImp) UpdateBidderAmountInAuction(auctionID int, participantID int, bidAmount int) (err error) {
	if _, ok := a.auctionDetails[auctionID]; !ok {
		return errors.New("invalid Auction ID")
	}
	if _, ok := a.auctionDetails[auctionID].participants[participantID]; !ok {
		return errors.New("bidder not present, please create one")
	}
	if a.auctionDetails[auctionID].auctionStatus == "closed" {
		return errors.New("you can't bid on a closed auction")
	}
	a.auctionDetails[auctionID].participants[participantID] = bidAmount
	return nil
}

func (a *auctionImp) WithdrawBidderFromAuction(auctionID int, participantID int) (err error) {
	if _, ok := a.auctionDetails[auctionID]; !ok {
		return errors.New("invalid Auction ID")
	}
	if _, ok := a.auctionDetails[auctionID].participants[participantID]; !ok {
		return errors.New("participant not available")
	}
	// TODO: lock this part as a txn
	delete(a.auctionDetails[auctionID].participants, participantID)
	auctionDetails := a.auctionDetails[auctionID]
	// TODO: Check this (If bidder withdraws from auction, should that be included or not as participation count)
	// auctionDetails.participantCount = auctionDetails.participantCount - 1
	a.auctionDetails[auctionID] = auctionDetails
	return nil
}

func (a *auctionImp) GetClosedAuctionDetails(auctionID int, sellerID int) (internal.AuctionCalculations, error) {
	auctionDetails, ok := a.auctionDetails[auctionID]
	if auctionDetails.sellerID != sellerID {
		return internal.AuctionCalculations{}, errors.New("you can't see profits of other sellers")
	}
	if !ok || auctionDetails.auctionStatus != "closed" {
		return internal.AuctionCalculations{}, errors.New("can't get details at this moment")
	}
	return internal.AuctionCalculations{
		FinalBidAmount:    auctionDetails.finalBidAmount,
		ParticipantCount:  auctionDetails.participantCount,
		ParticipationCost: auctionDetails.participationCost,
		LowestBidLimit:    auctionDetails.lowestBidLimit,
		HighestBidLimit:   auctionDetails.highestBidLimit,
		WinnerExists:      auctionDetails.winnerExists,
	}, nil
}
