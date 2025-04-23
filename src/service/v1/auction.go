package service

import (
	"NFTAuctionBackend/src/service/svc"
	"NFTAuctionBackend/src/types/v1"
	"context"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/stores/gdb/orderbookmodel/multi"
	"github.com/pkg/errors"
	"sync"
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

func GetAuctionDetail(ctx context.Context, svcCtx *svc.ServerCtx, AuctionId, ChainId int, Chain, TokenId, ContractAddress string) (*types.AuctionDetailResp, error) {
	var queryErr error
	var wg sync.WaitGroup

	// 并发查询以下信息:
	// 1.查询AuctionAndNft信息
	var auctionDetail *types.AuctionDetail
	wg.Add(1)
	go func() {
		defer wg.Done()
		auctionDetail, queryErr = svcCtx.Dao.QueryAuctionAndNft(ctx, AuctionId, ChainId, Chain, TokenId, ContractAddress)
		if queryErr != nil {
			queryErr = errors.Wrap(queryErr, "failed on get AuctionAndNft Info")
			return
		}
	}()

	// 2.查询NFT属性信息
	var nftAttributes []multi.NftAttributes
	wg.Add(1)
	go func() {
		defer wg.Done()
		nftAttributes, queryErr = svcCtx.Dao.QueryNftAttributes(ctx, TokenId)
		if queryErr != nil {
			queryErr = errors.Wrap(queryErr, "failed on get NftAttributes")
			return
		}
	}()

	// 3.查询出价信息
	var auctionBids []multi.AuctionBid
	wg.Add(1)
	go func() {
		defer wg.Done()
		auctionBids, queryErr = svcCtx.Dao.QueryAuctionBids(ctx, AuctionId)
		if queryErr != nil {
			queryErr = errors.Wrap(queryErr, "failed on get auctionBids")
			return
		}
	}()

	// 4. 等待所有查询完成
	wg.Wait()
	if queryErr != nil {
		return nil, errors.Wrap(queryErr, "failed on get auctionDetail")
	}

	// 5. 整合所有信息
	if auctionDetail != nil {
		if nftAttributes != nil {
			auctionDetail.NftAttributes = nftAttributes
		}
		if auctionBids != nil {
			auctionDetail.AuctionBids = auctionBids
		}
	}

	return &types.AuctionDetailResp{
		Result: auctionDetail,
	}, nil
}
