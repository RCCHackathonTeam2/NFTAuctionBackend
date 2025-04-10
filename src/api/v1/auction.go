package v1

import (
	"encoding/json"

	"github.com/RCCHackathonTeam2/NFTAuctionBase/errcode"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/xhttp"
	"github.com/gin-gonic/gin"

	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/service/svc"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/service/v1"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/types/v1"
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

		res, err := service.GetAuctions(c.Request.Context(), svcCtx, filter.Category, filter.AuctionType, filter.ChainId, filter.MinPrice, filter.MaxPrice, filter.Page, filter.PageSize)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(c, struct {
			Result interface{} `json:"result"`
		}{Result: res})
	}
}
