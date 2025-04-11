package service

import (
	"context"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/service/svc"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/types/v1"
	"github.com/pkg/errors"
)

func GetAuctions(ctx context.Context, svcCtx *svc.ServerCtx, category string, auctionType []string, chainId []int, minPrice, maxPrice float32, orderBy string, page, pageSize int) (*types.AuctionsResp, error) {
	auctions, count, err := svcCtx.Dao.QueryAuctions(ctx, category, auctionType, chainId, minPrice, maxPrice, orderBy, page, pageSize)
	if err != nil {
		return nil, errors.Wrap(err, "failed on get auctions")
	}
	return &types.AuctionsResp{
		Result: auctions,
		Count:  count,
	}, nil
}
