package dao

import (
	"context"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/types/v1"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (d *Dao) QueryAuctions(ctx context.Context, category string, auctionType []string, chainId []int, minPrice, maxPrice float32, orderBy string, page, pageSize int) (interface{}, int64, error) {
	db := d.DB.WithContext(ctx).Table("auctions").
		Select("auctions.auction_id, auctions.nft_id, auctions.auction_type, auctions.current_price, " +
			"auctions.currency_symbol, auctions.end_time, auctions.status, auctions.created_at," +
			"nfts.name as nft_name, nfts.chain_id, nfts.category, users.username as nft_creator")
	if auctionType != nil && len(auctionType) > 0 {
		db = db.Where("auctions.auction_type in (?)", auctionType)
	}
	if minPrice > 0 {
		db = db.Where("auctions.current_price >= ?", minPrice)
	}
	if maxPrice > 0 {
		db = db.Where("auctions.current_price <= ?", maxPrice)
	}
	if category != "" {
		db.Where("nfts.category = ?", category)
	}
	if chainId != nil && len(chainId) > 0 {
		db = db.Where("nfts.chain_id in (?)", chainId)
	}
	db.Joins("left join nfts on nfts.nft_id = auctions.nft_id and nfts.is_minted = 1")
	db.Joins("left join users on users.user_id = nfts.creator_id")

	// 查询总记录数
	var count int64
	countTx := db.Session(&gorm.Session{})
	if err := countTx.Count(&count).Error; err != nil {
		return nil, 0, errors.Wrap(db.Error, "failed on count auctions")
	}
	// 如果没有记录直接返回
	var auctions []types.Auctions
	if count == 0 {
		return auctions, count, nil
	}

	// 分页查询拍卖列表
	switch orderBy {
	case "updated_at":
		orderBy = "auctions.updated_at desc"
	case "low_to_high_price":
		orderBy = "auctions.current_price asc"
	case "high_to_low_price":
		orderBy = "auctions.current_price desc"
	case "end_time":
		orderBy = "auctions.end_time desc"
	case "bid_count":
		orderBy = "auctions.bid_count desc"
	default:
		orderBy = "auctions.updated_at desc"
	}
	if err := db.Order(orderBy).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Scan(&auctions).Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed on get auctions")
	}

	return auctions, count, nil
}
