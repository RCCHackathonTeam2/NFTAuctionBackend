package dao

import (
	"context"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/logger/xzap"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/stores/gdb/orderbookmodel/multi"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

func (d *Dao) CreateNft(ctx context.Context, newNFT multi.Nft) (bool, error) {
	if err := d.DB.WithContext(ctx).Table(multi.NftTableName("")).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&newNFT).Error; err != nil { // 将nft信息存入数据库
		xzap.WithContext(ctx).Error("failed on create nft", zap.Error(err))
		return false, nil
	}
	return true, nil
}
