package service

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/channels/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type Service struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func New(ctx *svc.ServiceContext) *Service {
	return &Service{
		svcCtx: ctx,
		Logger: logx.WithContext(ctx.Ctx()),
	}
}

var _ mtproto.RPCChannelsServer = (*Service)(nil)
