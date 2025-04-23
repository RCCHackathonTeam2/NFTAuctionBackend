package dao

import (
	"context"
	"fmt"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/logger/xzap"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/stores/gdb/orderbookmodel/multi"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (d *Dao) CreateNft(ctx context.Context, newNFT multi.Nft) (bool, int64, error) {
	// 开启 Debug 模式打印 SQL
	db := d.DB.WithContext(ctx).Debug().Table(multi.NftTableName(""))

	// 执行插入（移除了冲突策略以验证问题）
	result := db.Create(&newNFT)
	if result.Error != nil {
		xzap.WithContext(ctx).Error("插入失败", zap.Error(result.Error))
		return false, 0, result.Error
	}

	// 检查是否实际插入
	if result.RowsAffected == 0 {
		return false, 0, fmt.Errorf("未插入数据（可能冲突或条件不满足）")
	}

	xzap.WithContext(ctx).Info("插入成功", zap.Int64("nft_id", newNFT.NftId))
	return true, newNFT.NftId, nil
}

func (d *Dao) QueryNftAttributes(ctx context.Context, TokenId string) ([]multi.NftAttributes, error) {
	var nftAttributes []multi.NftAttributes
	db := d.DB.WithContext(ctx).Table("nft_attributes").
		Select("attribute_id, token_id, trait_type, trait_value, display_type, rarity_percentage, created_at").
		Where("token_id = ?", TokenId)
	if err := db.Scan(&nftAttributes).Error; err != nil {
		return nil, errors.Wrap(err, "failed on get NftAttributes")
	}
	return nftAttributes, nil
}
