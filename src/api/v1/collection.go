package v1

import (
	"encoding/json"
	"fmt"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/service/v1"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/types/v1"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/errcode"

	//"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/errcode"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/api/middleware"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/service/svc"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/logger/xzap"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/xhttp"
)

// CollectionItemsHandler godoc
// @Summary 获取集合中的物品列表
// @Description 根据过滤条件获取指定集合中的NFT物品列表
// @Tags collections
// @Accept json
// @Produce json
// @Param address path string true "集合地址"
// @Param filters query string true "过滤参数，JSON格式"
// @Success 200 {object} interface{} "物品列表"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /collections/{address}/items [get]
func CollectionItemsHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		filterParam := c.Query("filters")
		if filterParam == "" {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		var filter types.CollectionItemFilterParams
		err := json.Unmarshal([]byte(filterParam), &filter)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[filter.ChainID]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}
		res, err := service.GetItems(c.Request.Context(), svcCtx, chain, filter, collectionAddr)
		if err != nil {
			xhttp.Error(c, errcode.ErrUnexpected)
			return
		}
		xhttp.OkJson(c, res)
	}
}

// CollectionBidsHandler godoc
// @Summary 获取集合的出价信息
// @Description 获取指定集合的所有出价信息
// @Tags collections
// @Accept json
// @Produce json
// @Param address path string true "集合地址"
// @Param filters query string true "过滤参数，JSON格式"
// @Success 200 {object} interface{} "出价信息列表"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /collections/{address}/bids [get]
func CollectionBidsHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Query：写到query的参数
		filterParam := c.Query("filters")
		if filterParam == "" {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		var filter types.CollectionBidFilterParams
		err := json.Unmarshal([]byte(filterParam), &filter)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		//c.Params.ByName:写到路径上的参数
		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[int(filter.ChainID)]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		res, err := service.GetBids(c.Request.Context(), svcCtx, chain, collectionAddr, filter.Page, filter.PageSize)
		if err != nil {
			xhttp.Error(c, errcode.ErrUnexpected)
			return
		}
		xhttp.OkJson(c, res)
	}
}

// CollectionItemBidsHandler godoc
// @Summary 获取集合中特定物品的出价信息
// @Description 获取集合中指定物品的所有出价信息
// @Tags collections
// @Accept json
// @Produce json
// @Param address path string true "集合地址"
// @Param token_id path string true "物品Token ID"
// @Param filters query string true "过滤参数，JSON格式"
// @Success 200 {object} interface{} "物品出价信息列表"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /collections/{address}/{token_id}/bids [get]
func CollectionItemBidsHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		filterParam := c.Query("filters")
		if filterParam == "" {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		var filter types.CollectionBidFilterParams
		err := json.Unmarshal([]byte(filterParam), &filter)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		tokenID := c.Params.ByName("token_id")
		if tokenID == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[int(filter.ChainID)]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		res, err := service.GetItemBidsInfo(c.Request.Context(), svcCtx, chain, collectionAddr, tokenID, filter.Page, filter.PageSize)
		if err != nil {
			xhttp.Error(c, errcode.ErrUnexpected)
			return
		}
		xhttp.OkJson(c, res)
	}
}

// ItemDetailHandler godoc
// @Summary 获取物品详情
// @Description 获取指定物品的详细信息
// @Tags collections
// @Accept json
// @Produce json
// @Param address path string true "集合地址"
// @Param token_id path string true "物品Token ID"
// @Param chain_id query int true "链ID"
// @Success 200 {object} interface{} "物品详情"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /collections/{address}/{token_id} [get]
func ItemDetailHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		tokenID := c.Params.ByName("token_id")
		if tokenID == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chainID, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
		if err != nil {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[int(chainID)]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		res, err := service.GetItem(c.Request.Context(), svcCtx, chain, int(chainID), collectionAddr, tokenID)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("get item error"))
			return

		}
		xhttp.OkJson(c, res)
	}
}

// ItemTopTraitPriceHandler godoc
// @Summary 获取物品特性的最高价格信息
// @Description 获取指定集合中物品特性的最高价格信息
// @Tags collections
// @Accept json
// @Produce json
// @Param address path string true "集合地址"
// @Param filters query string true "过滤参数，JSON格式"
// @Success 200 {object} interface{} "特性价格信息"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /collections/{address}/top-trait [get]
func ItemTopTraitPriceHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		filterParam := c.Query("filters")
		if filterParam == "" {
			xhttp.Error(c, errcode.NewCustomErr("Filter param is nil."))
			return
		}

		var filter types.TopTraitFilterParams
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

		res, err := service.GetItemTopTraitPrice(c.Request.Context(), svcCtx, chain, collectionAddr, filter.TokenIds)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("get item error"))
			return
		}
		xhttp.OkJson(c, res)
	}
}

// HistorySalesHandler godoc
// @Summary 获取历史销售价格信息
// @Description 获取指定集合的历史销售价格信息
// @Tags collections
// @Accept json
// @Produce json
// @Param address path string true "集合地址"
// @Param chain_id query int true "链ID"
// @Param duration query string false "时间范围(24h/7d/30d)" default(7d)
// @Success 200 {object} interface{} "历史销售价格信息"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /collections/{address}/history-sales [get]
func HistorySalesHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chainID, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
		if err != nil {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[int(chainID)]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		duration := c.Query("duration")
		if duration != "" {
			validParams := map[string]bool{
				"24h": true,
				"7d":  true,
				"30d": true,
			}
			if ok := validParams[duration]; !ok {
				xzap.WithContext(c).Error("duration parse error: ", zap.String("duration", duration))
				xhttp.Error(c, errcode.ErrInvalidParams)
				return
			}
		} else {
			duration = "7d"
		}

		res, err := service.GetHistorySalesPrice(c.Request.Context(), svcCtx, chain, collectionAddr, duration)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("get history sales price error"))
			return
		}

		xhttp.OkJson(c, struct {
			Result interface{} `json:"result"`
		}{
			Result: res,
		})
	}
}

// ItemTraitsHandler godoc
// @Summary 获取物品特性信息
// @Description 获取指定物品的特性(Attribute)信息
// @Tags collections
// @Accept json
// @Produce json
// @Param address path string true "集合地址"
// @Param token_id path string true "物品Token ID"
// @Param chain_id query int true "链ID"
// @Success 200 {object} interface{} "物品特性信息"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /collections/{address}/{token_id}/traits [get]
func ItemTraitsHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		tokenID := c.Params.ByName("token_id")
		if tokenID == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chainID, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
		if err != nil {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[int(chainID)]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		itemTraits, err := service.GetItemTraits(c.Request.Context(), svcCtx, chain, collectionAddr, tokenID)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("get item traits error"))
			return
		}

		xhttp.OkJson(c, types.ItemTraitsResp{Result: itemTraits})
	}
}

func ItemOwnerHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		tokenID := c.Params.ByName("token_id")
		if tokenID == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chainID, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
		if err != nil {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[int(chainID)]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		owner, err := service.GetItemOwner(c.Request.Context(), svcCtx, chainID, chain, collectionAddr, tokenID)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("get item owner error"))
			return
		}

		xhttp.OkJson(c, struct {
			Result interface{} `json:"result"`
		}{
			Result: owner,
		})
	}
}

func GetItemImageHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		tokenID := c.Params.ByName("token_id")
		if tokenID == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chainID, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
		if err != nil {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[int(chainID)]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		result, err := service.GetItemImage(c.Request.Context(), svcCtx, chain, collectionAddr, tokenID)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr("failed on get item image"))
			return
		}

		xhttp.OkJson(c, struct {
			Result interface{} `json:"result"`
		}{Result: result})
	}
}

func ItemMetadataRefreshHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		chainId, err := strconv.ParseInt(c.Query("chain_id"), 10, 32)
		if err != nil {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[int(chainId)]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		tokenId := c.Params.ByName("token_id")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		err = service.RefreshItemMetadata(c.Request.Context(), svcCtx, chain, chainId, collectionAddr, tokenId)
		if err != nil {
			xhttp.Error(c, err)
			return
		}

		successStr := "Success to joined the refresh queue and waiting for refresh."
		xhttp.OkJson(c, types.CommonResp{Result: successStr})
	}
}

func CollectionDetailHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		chainID, err := strconv.ParseInt(c.Query("chain_id"), 10, 32)
		if err != nil {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		chain, ok := chainIDToChain[int(chainID)]
		if !ok {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		collectionAddr := c.Params.ByName("address")
		if collectionAddr == "" {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}
		res, err := service.GetCollectionDetail(c.Request.Context(), svcCtx, chain, collectionAddr)
		if err != nil {
			xhttp.Error(c, errcode.ErrUnexpected)
			return
		}

		xhttp.OkJson(c, res)
	}
}

type CreateNftRequest struct {
	name              string `form:"name" binding:"required"`               //
	description       string `form:"description" binding:"required"`        //
	imageUrl          string `form:"image_url" binding:"required"`          //
	royaltyPercentage string `form:"royalty_percentage" binding:"required"` //
	chainId           string `form:"chain_id" binding:"required"`           //
	categorieId       string `form:"categorie_id" binding:"required"`       //
}

// CreateNft godoc
// @Summary 保存 NFT 信息
// @Description 保存 NFT 信息
// @Tags collections
// @Accept json
// @Produce json
// @Param name query string true "nft名称"
// @Param description query string true "NFT描述"
// @Param imageUrl query string true "NFT图片URL"
// @Param royaltyPercentage query string true "版税百分比"
// @Param chainId query int true "所属区块链网络id"
// @Param categorieId query string true "类别id"
// @Success 200 {object} interface{} "出价信息列表"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /collections/createNft [post]
func CreateNft(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {

		currentUserAddress, _ := middleware.GetAuthUserAddress(c, svcCtx.KvStore)
		fmt.Println(currentUserAddress, "我看看看看")
		//req := new(CreateNftRequest)
		name := c.Query("name")
		if name == "" {
			xhttp.Error(c, errcode.NewCustomErr("name param is nil."))
			return
		}
		description := c.Query("description")
		if description == "" {
			xhttp.Error(c, errcode.NewCustomErr("description param is nil."))
			return
		}
		imageUrl := c.Query("imageUrl")
		if imageUrl == "" {
			xhttp.Error(c, errcode.NewCustomErr("imageUrl param is nil."))
			return
		}
		royaltyPercentage := c.Query("royaltyPercentage")
		if royaltyPercentage == "" {
			xhttp.Error(c, errcode.NewCustomErr("royaltyPercentage param is nil."))
			return
		}

		chainId, err := strconv.ParseInt(c.Query("chainId"), 10, 32)
		if err != nil {
			xhttp.Error(c, errcode.ErrInvalidParams)
			return
		}

		categorieId := c.Query("categorieId")
		if categorieId == "" {
			xhttp.Error(c, errcode.NewCustomErr("categorieId param is nil."))
			return
		}

		res, err := service.CreateNft(c.Request.Context(), svcCtx, int(chainId), categorieId, royaltyPercentage, imageUrl, description, name, currentUserAddress)
		if err != nil {
			xhttp.Error(c, errcode.ErrUnexpected)
			return
		}

		xhttp.OkJson(c, res)
	}
}
