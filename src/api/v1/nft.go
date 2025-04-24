package v1

import (
	"NFTAuctionBackend/src/service/v1"
	"fmt"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/errcode"

	//"NFTAuctionBackend/src/errcode"
	"strconv"

	"NFTAuctionBackend/src/api/middleware"
	"NFTAuctionBackend/src/service/svc"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/xhttp"
	"github.com/gin-gonic/gin"
)

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
// @Param tokenId query string true "tokenId"
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

		tokenId := c.Query("tokenId")
		if tokenId == "" {
			xhttp.Error(c, errcode.NewCustomErr("tokenId param is nil."))
			return
		}

		res, err := service.CreateNft(c.Request.Context(), svcCtx, int(chainId), categorieId, royaltyPercentage, imageUrl, description, name, currentUserAddress, tokenId)
		if err != nil {
			xhttp.Error(c, errcode.ErrUnexpected)
			return
		}

		xhttp.OkJson(c, res)
	}
}
