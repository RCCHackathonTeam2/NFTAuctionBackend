package v1

import (
	"github.com/RCCHackathonTeam2/NFTAuctionBase/errcode"
	//"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/errcode"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/kit/validator"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/xhttp"
	"github.com/gin-gonic/gin"

	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/service/svc"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/service/v1"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/types/v1"
)

// UserLoginHandler godoc
// @Summary 用户登录qaq
// @Description 用户通过签名信息登录系统
// @Tags user
// @Accept json
// @Produce json
// @Param request body types.LoginReq true "登录请求参数"
// @Success 200 {object} types.UserLoginResp "登录成功响应"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /user/login [post]
func UserLoginHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := types.LoginReq{}
		if err := c.BindJSON(&req); err != nil {
			xhttp.Error(c, err)
			return
		}

		if err := validator.Verify(&req); err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}

		res, err := service.UserLogin(c.Request.Context(), svcCtx, req)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}

		xhttp.OkJson(c, types.UserLoginResp{
			Result: res,
		})
	}
}

// GetLoginMessageHandler godoc
// @Summary 获取登录签名信息
// @Description 获取用户登录所需的签名信息
// @Tags user
// @Accept json
// @Produce json
// @Param address path string true "用户地址"
// @Success 200 {object} interface{} "登录签名信息"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /user/{address}/login-message [get]
func GetLoginMessageHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Params.ByName("address")
		if address == "" {
			xhttp.Error(c, errcode.NewCustomErr("user addr is null"))
			return
		}

		res, err := service.GetUserLoginMsg(c.Request.Context(), svcCtx, address)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}

		xhttp.OkJson(c, res)
	}
}

// GetSigStatusHandler godoc
// @Summary 获取用户签名状态
// @Description 获取指定用户的签名状态信息
// @Tags user
// @Accept json
// @Produce json
// @Param address path string true "用户地址"
// @Success 200 {object} interface{} "签名状态信息"
// @Failure 400 {object} errcode.Error "参数错误"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /user/{address}/sig-status [get]
func GetSigStatusHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAddr := c.Params.ByName("address")
		if userAddr == "" {
			xhttp.Error(c, errcode.NewCustomErr("user addr is null"))
			return
		}

		res, err := service.GetSigStatusMsg(c.Request.Context(), svcCtx, userAddr)
		if err != nil {
			xhttp.Error(c, errcode.NewCustomErr(err.Error()))
			return
		}

		xhttp.OkJson(c, res)
	}
}
