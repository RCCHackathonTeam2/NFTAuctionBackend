package main

import (
	"flag"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/api/router"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/app"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/config"
	"github.com/RCCHackathonTeam2/NFTAuctionBackend/src/service/svc"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	// port       = ":9000"
	repoRoot          = ""
	defaultConfigPath = "./config/config.toml"
)

// @title NFT Auction API
// @version 1.0
// @description NFT Auction Backend API Documentation
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.nftauction.com/support
// @contact.email support@nftauction.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
func main() {
	conf := flag.String("conf", defaultConfigPath, "conf file path")
	flag.Parse()
	c, err := config.UnmarshalConfig(*conf)
	if err != nil {
		panic(err)
	}

	for _, chain := range c.ChainSupported {
		if chain.ChainID == 0 || chain.Name == "" {
			panic("invalid chain_suffix config")
		}
	}

	serverCtx, err := svc.NewServiceContext(c)
	if err != nil {
		panic(err)
	}
	// Initialize router
	r := router.NewRouter(serverCtx)
	// 添加swagger路由
	url := ginSwagger.URL("/src/docs/swagger.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	app, err := app.NewPlatform(c, r, serverCtx)
	if err != nil {
		panic(err)
	}
	app.Start()
}
