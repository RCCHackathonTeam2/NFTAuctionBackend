package svc

import (
	"github.com/RCCHackathonTeam2/NFTAuctionBase/evm/erc"
	//"github.com/RCCHackathonTeam2/NFTAuctionBase/image"
	"github.com/RCCHackathonTeam2/NFTAuctionBase/stores/xkv"
	"gorm.io/gorm"

	"NFTAuctionBackend/src/dao"
)

type CtxConfig struct {
	db *gorm.DB
	//imageMgr image.ImageManager
	dao     *dao.Dao
	KvStore *xkv.Store
	Evm     erc.Erc
}

type CtxOption func(conf *CtxConfig)

func NewServerCtx(options ...CtxOption) *ServerCtx {
	c := &CtxConfig{}
	for _, opt := range options {
		opt(c)
	}
	return &ServerCtx{
		DB: c.db,
		//ImageMgr: c.imageMgr,
		KvStore: c.KvStore,
		Dao:     c.dao,
	}
}

func WithKv(kv *xkv.Store) CtxOption {
	return func(conf *CtxConfig) {
		conf.KvStore = kv
	}
}

func WithDB(db *gorm.DB) CtxOption {
	return func(conf *CtxConfig) {
		conf.db = db
	}
}

func WithDao(dao *dao.Dao) CtxOption {
	return func(conf *CtxConfig) {
		conf.dao = dao
	}
}
