package service

import (
	"context"
	"time"

	"github.com/RCCHackathonTeam2/NFTAuctionBase/logger/xzap"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/stores/gdb/orderbookmodel/multi"
	"go.uber.org/zap"

	"NFTAuctionBackend/src/service/svc"
	"NFTAuctionBackend/src/types/v1"
)

func CreateNft(ctx context.Context, svcCtx *svc.ServerCtx, chain int, categorie string, royaltyPercentage string,
	imageUrl string, description string, name string, currentUserAddress string, tokenId string) (*types.CreateNftResp, error) {
	newNFT := multi.Nft{
		TokenId:           tokenId,
		ContractAddress:   svcCtx.C.ContractAddress.NftAuctionAddress,
		ChainId:           int64(chain),
		Category:          categorie,
		Name:              name,
		Description:       description,
		ImagUrl:           imageUrl,
		ThumbnailUrl:      imageUrl,
		MetadataUrl:       imageUrl,
		CreatorId:         currentUserAddress,
		OwnerId:           currentUserAddress,
		RoyaltyPercentage: royaltyPercentage,
		TokenStandard:     "ERC721",
		TotalSupply:       1,
		IsMinted:          0,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Status:            "",
	}

	isSuccess, nftId, err := svcCtx.Dao.CreateNft(ctx, newNFT)
	if err != nil {
		xzap.WithContext(ctx).Error("failed on CreateNft", zap.Error(err))
	}
	return &types.CreateNftResp{Result: isSuccess, NftId: nftId}, nil
}
