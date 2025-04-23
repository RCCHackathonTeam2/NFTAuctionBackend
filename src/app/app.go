package app

import (
	"context"

	"github.com/RCCHackathonTeam2/NFTAuctionBase/logger/xzap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"NFTAuctionBackend/src/config"
	"NFTAuctionBackend/src/service/svc"
)

type Platform struct {
	config    *config.Config
	router    *gin.Engine
	serverCtx *svc.ServerCtx
}

func NewPlatform(config *config.Config, router *gin.Engine, serverCtx *svc.ServerCtx) (*Platform, error) {
	return &Platform{
		config:    config,
		router:    router,
		serverCtx: serverCtx,
	}, nil
}

func (p *Platform) Start() {
	xzap.WithContext(context.Background()).Info("EasySwap-End run", zap.String("port", p.config.Api.Port))
	if err := p.router.Run(p.config.Api.Port); err != nil {
		panic(err)
	}
}
