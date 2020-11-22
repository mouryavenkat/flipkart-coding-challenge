package main

import (
	"flipkart/internal/auction"
	"flipkart/internal/buyer"
	"flipkart/internal/seller"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testcases := [][]string{
		{
			"create buyer 1 mourya mobile1",
			"create buyer 2 mourya1 mobile2",
			"create buyer 3 mourya2 mobile3",
			"create buyer 4 mourya mobile1",
			"create buyer 5 mourya mobile1",
			"create buyer 6 mourya mobile1",
			"create seller 9876 mourya",
			"create auction 12345 10 50 1 9876",
			"create auction 12345 10 50 1 9876",
			"create bid 1 12345 17",
			"create bid 2 12345 15",
			"create bid 6 12345 19",
			"create bid 3 12345 19",
			"create bid 4 12345 17",
			"create bid 5 12345 15",
			"close auction 12345 9876",
			"get profit 9876 12345",
		},
		//{
		//	"create seller 9876 mourya",
		//	"create buyer 1 mourya mobile1",
		//	"create buyer 2 mourya1 mobile2",
		//	"create buyer 3 mourya2 mobile3",
		//	"create auction 12345 5 20 2 9876",
		//	"create bid 3 12345 25",
		//	"create bid 2 12345 5",
		//	"withdraw bid 2 12345",
		//	"close auction 12345 9876",
		//	"get profit 9876 12345",
		//},
		//// If no bidders in the auction. (Expected profit: 0)
		//{
		//	"create buyer 1 mourya mobile1",
		//	"create buyer 2 mourya1 mobile2",
		//	"create buyer 3 mourya2 mobile3",
		//	"create seller 9876 mourya",
		//	"create auction 12345 10 50 1 9876",
		//	"close auction 12345 9876",
		//	"get profit 9876 12345",
		//},
		//
		//// When seller tries to fetch profit without closing auction. (Expected profit: 0)
		//{
		//	"create buyer 1 mourya mobile1",
		//	"create buyer 2 mourya1 mobile2",
		//	"create buyer 3 mourya2 mobile3",
		//	"create seller 9876 mourya",
		//	"create auction 12345 10 50 1 9876",
		//	"get profit 9876 12345",
		//},
		//
		//// When one seller tries to fetch profits of another seller
		//{
		//	"create buyer 1 mourya mobile1",
		//	"create buyer 2 mourya1 mobile2",
		//	"create buyer 3 mourya2 mobile3",
		//	"create seller 9876 mourya",
		//	"create seller 98765 mourya",
		//	"create auction 12345 10 50 1 9876",
		//	"close auction 12345 9876",
		//	"get profit 98765 12345",
		//},
		//
		//// When user trying to bid on a closed auction
		//{
		//	"create buyer 1 mourya mobile1",
		//	"create buyer 2 mourya1 mobile2",
		//	"create buyer 3 mourya2 mobile3",
		//	"create seller 9876 mourya",
		//	"create seller 98765 mourya",
		//	"create auction 12345 10 50 1 9876",
		//	"close auction 12345 9876",
		//	"create bid 2 12345 25",
		//},
	}
	for index, testcase := range testcases {
		fmt.Println(fmt.Sprintf("Executing test case %d", index+1))
		auctionClient := auction.NewAction()
		buyerClient := buyer.NewBuyer(auctionClient)
		sellerClient := seller.NewSeller(auctionClient)
		for _, instruction := range testcase {
			splitText := strings.Split(instruction, " ")
			if len(splitText) == 0 {
				fmt.Println("Enter proper command")
				continue
			}
			switch splitText[0] {
			case "create":
				switch splitText[1] {
				case "buyer":
					buyerID, _ := strconv.ParseInt(splitText[2], 10, 64)
					if err := buyerClient.RegisterBuyer(int(buyerID), splitText[3], splitText[4]); err != nil {
						fmt.Println("Error registering user", err.Error())
					}
				case "seller":
					sellerID, _ := strconv.ParseInt(splitText[2], 10, 64)
					sellerClient.RegisterSeller(int(sellerID), splitText[3])
				case "auction":
					auctionID, _ := strconv.ParseInt(splitText[2], 10, 64)
					lowestBid, _ := strconv.ParseInt(splitText[3], 10, 64)
					highestBid, _ := strconv.ParseInt(splitText[4], 10, 64)
					participationCost, _ := strconv.ParseInt(splitText[5], 10, 64)
					sellerID, _ := strconv.ParseInt(splitText[6], 10, 64)
					if err := sellerClient.CreateAuction(int(auctionID), int(lowestBid), int(highestBid), int(participationCost),
						int(sellerID)); err != nil {
						fmt.Println("Error creating auction", err.Error())
					}
				case "bid":
					buyerID, _ := strconv.ParseInt(splitText[2], 10, 64)
					auctionID, _ := strconv.ParseInt(splitText[3], 10, 64)
					bidAmount, _ := strconv.ParseInt(splitText[4], 10, 64)
					if err := buyerClient.CreateBid(int(buyerID), int(auctionID), int(bidAmount)); err != nil {
						fmt.Println("Error creating bid", err.Error())
					}
				}
			case "update":
				switch splitText[1] {
				case "bid":
					buyerID, _ := strconv.ParseInt(splitText[2], 10, 64)
					auctionID, _ := strconv.ParseInt(splitText[3], 10, 64)
					bidAmount, _ := strconv.ParseInt(splitText[4], 10, 64)
					buyerClient.UpdateBid(int(buyerID), int(auctionID), int(bidAmount))
				}
			case "get":
				switch splitText[1] {
				case "profit":
					sellerID, _ := strconv.ParseInt(splitText[2], 10, 64)
					auctionID, _ := strconv.ParseInt(splitText[3], 10, 64)
					fmt.Println(sellerClient.GetProfit(int(sellerID), int(auctionID)))
				}
			case "close":
				auctionID, _ := strconv.ParseInt(splitText[2], 10, 64)
				sellerID, _ := strconv.ParseInt(splitText[3], 10, 64)
				sellerClient.CloseAuction(int(auctionID), int(sellerID))
			case "withdraw":
				buyerID, _ := strconv.ParseInt(splitText[2], 10, 64)
				auctionID, _ := strconv.ParseInt(splitText[3], 10, 64)
				buyerClient.WithdrawBid(int(buyerID), int(auctionID))
			}
		}
		fmt.Println("---------------------")
	}
}
