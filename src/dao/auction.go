package dao

import (
	"context"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/types/v1"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (d *Dao) QueryAuctions(ctx context.Context, Category, auctionType string, ChainId int, MinPrice, MaxPrice float32, Page, PageSize int) (interface{}, int64, error) {
	db := d.DB.WithContext(ctx).Table("auctions").
		Select("auctions.auction_id, auctions.nft_id, auctions.auction_type, auctions.current_price, " +
			"auctions.currency_symbol, auctions.end_time, auctions.status, auctions.created_at," +
			"nfts.name as nft_name, nfts.chain_id, nfts.category, users.username as nft_creator")
	if auctionType != "" {
		db = db.Where("auctions.auction_type = ?", auctionType)
	}
	if MinPrice > 0 {
		db = db.Where("auctions.current_price >= ?", MinPrice)
	}
	if MaxPrice > 0 {
		db = db.Where("auctions.current_price <= ?", MaxPrice)
	}
	if Category != "" {
		db.Where("nfts.category = ?", Category)
	}
	if ChainId > 0 {
		db = db.Where("nfts.chain_id = ?", ChainId)
	}
	db.Joins("left join nfts on nfts.nft_id = auctions.nft_id")
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

	// 分页查询拍卖列表,按价格降序排列
	if err := db.Order("auctions.current_price").
		Offset((Page - 1) * PageSize).
		Limit(PageSize).
		Scan(&auctions).Error; err != nil {
		return nil, 0, errors.Wrap(err, "failed on get auctions")
	}

	return auctions, count, nil
}
