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

// OrderInfosHandler godoc
// @Summary 批量查询出价信息
// @Description 批量查询出价信息
// @Tags user
// @Accept json
// @Produce json
// @Param filters query string true "过滤参数，JSON格式"
// @Success 200 {object} interface{} "出价信息列表"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /portfolio/bid-orders [post]
func OrderInfosHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		filterParam := c.Query("filters")
		if filterParam == "" {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		var filter types.OrderInfosParam
		err := json.Unmarshal([]byte(filterParam), &filter)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		chain, ok := chainIDToChain[filter.ChainID]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		res, err := service.GetOrderInfos(c.Request.Context(), svcCtx, filter.ChainID, chain, filter.UserAddress, filter.CollectionAddress, filter.TokenIds)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}
		xhttp.OkJson(c, struct {
			Result interface{} `json:"result"`
		}{Result: res})
	}
}
