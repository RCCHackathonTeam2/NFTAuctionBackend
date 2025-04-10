package types

import (
	"time"
)

type AuctionsParam struct {
	Category    string  `json:"category"`
	AuctionType string  `json:"auction_type"`
	ChainId     int     `json:"chain_id"`
	MinPrice    float32 `json:"min_price"`
	MaxPrice    float32 `json:"max_price"`
	Page        int     `json:"page"`
	PageSize    int     `json:"page_size"`
}

type Auctions struct {
	AuctionId      int       `json:"auction_id"`
	NftId          int       `json:"nft_id"`
	Category       string    `json:"category"`
	AuctionType    string    `json:"auction_type"`
	ChainId        int       `json:"chain_id"`
	CurrentPrice   float32   `json:"current_price"`
	CurrencySymbol string    `json:"currency_symbol"`
	EndTime        time.Time `json:"end_time"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	NftName        string    `json:"nft_name"`
	NftCreator     string    `json:"nft_creator"`
}
