package v1

import (
	"encoding/json"

	"github.com/RCCHackathonTeam2/NFTAuctionBase/errcode"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/xhttp"
	"github.com/gin-gonic/gin"

	"NFTAuctionBackend/src/service/svc"
	"NFTAuctionBackend/src/service/v1"
	"NFTAuctionBackend/src/types/v1"
)

func AuctionsHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		filterParam := c.Query("filters")
		if filterParam == "" {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		var filter types.AuctionsParam
		err := json.Unmarshal([]byte(filterParam), &filter)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		res, err := service.GetAuctions(c.Request.Context(), svcCtx, filter.Category, filter.AuctionType, filter.ChainId, filter.MinPrice, filter.MaxPrice, filter.OrderBy, filter.Page, filter.PageSize)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(c, res)
	}
}

func AuctionDetailHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		filterParam := c.Query("filters")
		if filterParam == "" {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		var filter types.AuctionDetailParam
		err := json.Unmarshal([]byte(filterParam), &filter)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}
		if filter.AuctionId <= 0 || filter.ChainId <= 0 || filter.TokenId == "" || filter.ContractAddress == "" {
			xhttp.Error(c, errcode.NewCustomErr("Invalid params."))
			return
		}
		Chain, ok := chainIDToChain[filter.ChainId]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}
		res, err := service.GetAuctionDetail(c.Request.Context(), svcCtx, filter.AuctionId, filter.ChainId, Chain, filter.TokenId, filter.ContractAddress)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(c, res)
	}
}
