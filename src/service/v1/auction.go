package service

import (
	"context"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/service/svc"
)

func GetAuctions(ctx context.Context, svcCtx *svc.ServerCtx, Category, auctionType string, ChainId int, MinPrice, MaxPrice float32, Page, PageSize int) (interface{}, error) {
	// todo
	auctions, _, err := svcCtx.Dao.QueryAuctions(ctx, Category, auctionType, ChainId, MinPrice, MaxPrice, Page, PageSize)
	return auctions, err
}
