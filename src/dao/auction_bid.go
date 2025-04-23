package dao

import (
	"context"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/stores/gdb/orderbookmodel/multi"
	"github.com/pkg/errors"
)

func (d *Dao) QueryAuctionBids(ctx context.Context, AuctionId int) ([]multi.AuctionBid, error) {
	var auctionBis []multi.AuctionBid
	db := d.DB.WithContext(ctx).Table("bids").
		Select("bid_id, auction_id, bidder, bid_amount, transaction_hash, status, created_at, updated_at").
		Where("auction_id = ?", AuctionId)
	if err := db.Scan(&auctionBis).Error; err != nil {
		return nil, errors.Wrap(err, "failed on get AuctionBids")
	}
	return auctionBis, nil
}
