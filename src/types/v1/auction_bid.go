package types

import "time"

type AuctionBid struct {
	Bidder          string    `json:"bidder"`
	BidAmount       float32   `json:"bid_amount"`
	TransactionHash string    `json:"transaction_hash"`
	BidStatus       string    `json:"bid_status"`
	CreatedAt       time.Time `json:"created_at"`
}
