package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/bots/internal/core"
	"github.com/teamgram/teamgram-server/app/bff/bots/internal/svc"
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

var _ mtproto.RPCBotsServer = (*Service)(nil)

func (s *Service) BotsSetBotCommands(ctx context.Context, in *mtproto.TLBotsSetBotCommands) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.BotsSetBotCommands(in)
}

func (s *Service) BotsResetBotCommands(ctx context.Context, in *mtproto.TLBotsResetBotCommands) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.BotsResetBotCommands(in)
}

func (s *Service) BotsGetBotCommands(ctx context.Context, in *mtproto.TLBotsGetBotCommands) (*mtproto.Vector_BotCommand, error) {
	c := core.New(ctx, s.svcCtx)
	return c.BotsGetBotCommands(in)
}

func (s *Service) BotsSetBotInfo(ctx context.Context, in *mtproto.TLBotsSetBotInfo) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.BotsSetBotInfo(in)
}

func (s *Service) BotsGetBotInfoDCD914FD(ctx context.Context, in *mtproto.TLBotsGetBotInfoDCD914FD) (*mtproto.Bots_BotInfo, error) {
	c := core.New(ctx, s.svcCtx)
	return c.BotsGetBotInfoDCD914FD(in)
}

func (s *Service) BotsGetAdminedBots(ctx context.Context, in *mtproto.TLBotsGetAdminedBots) (*mtproto.Vector_User, error) {
	c := core.New(ctx, s.svcCtx)
	return c.BotsGetAdminedBots(in)
}

func (s *Service) BotsGetBotInfo75EC12E6(ctx context.Context, in *mtproto.TLBotsGetBotInfo75EC12E6) (*mtproto.Vector_String, error) {
	c := core.New(ctx, s.svcCtx)
	return c.BotsGetBotInfo75EC12E6(in)
}
