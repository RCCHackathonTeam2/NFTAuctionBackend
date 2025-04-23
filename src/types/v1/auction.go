package types

import (
	"github.com/RCCHackathonTeam2/NFTAuctionBase/stores/gdb/orderbookmodel/multi"
	"time"
)

type AuctionsParam struct {
	Category    string   `json:"category"`
	AuctionType []string `json:"auction_type"`
	ChainId     []int    `json:"chain_id"`
	MinPrice    float32  `json:"min_price"`
	MaxPrice    float32  `json:"max_price"`
	OrderBy     string   `json:"order_by"`
	Page        int      `json:"page"`
	PageSize    int      `json:"page_size"`
}

type AuctionDetailParam struct {
	AuctionId       int    `json:"auction_id"`
	ChainId         int    `json:"chain_id"`
	TokenId         string `json:"token_id"`
	ContractAddress string `json:"contract_address"`
}

type Auctions struct {
	AuctionId       int       `json:"auction_id"`
	TokenId         string    `json:"token_id"`
	Category        string    `json:"category"`
	AuctionType     string    `json:"auction_type"`
	ChainId         int       `json:"chain_id"`
	CurrentPrice    float32   `json:"current_price"`
	CurrencySymbol  string    `json:"currency_symbol"`
	EndTime         time.Time `json:"end_time"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	NftName         string    `json:"nft_name"`
	NftCreator      string    `json:"nft_creator"`
	ContractAddress string    `json:"contract_address"`
	ThumbnailUrl    string    `json:"thumbnail_url"`
	AvatarUrl       string    `json:"avatar_url"`
}

type AuctionsResp struct {
	Result interface{} `json:"result"`
	Count  int64       `json:"count"`
}

type AuctionDetailResp struct {
	Result interface{} `json:"result"`
}

type AuctionDetail struct {
	BidCount          int                   `json:"bid_count"`
	Winner            string                `json:"winner"`
	Description       string                `json:"description"`
	ImageUrl          string                `json:"image_url"`
	MetadataUrl       string                `json:"metadata_url"`
	OwnerId           string                `json:"owner_id"`
	RoyaltyPercentage string                `json:"royalty_percentage"`
	TokenStandard     string                `json:"token_standard"`
	MintedAt          time.Time             `json:"minted_at"`
	NftStatus         string                `json:"nft_status"`
	NftAttributes     []multi.NftAttributes `gorm:"-" json:"nft_attributes"`
	AuctionBids       []multi.AuctionBid    `gorm:"-" json:"auction_bids"`
}
