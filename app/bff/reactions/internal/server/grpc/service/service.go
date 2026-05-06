package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/reactions/internal/core"
	"github.com/teamgram/teamgram-server/app/bff/reactions/internal/svc"
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

var _ mtproto.RPCReactionsServer = (*Service)(nil)

func (s *Service) MessagesSendReaction(ctx context.Context, in *mtproto.TLMessagesSendReaction) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesSendReaction(in)
}

func (s *Service) MessagesGetMessagesReactions(ctx context.Context, in *mtproto.TLMessagesGetMessagesReactions) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetMessagesReactions(in)
}

func (s *Service) MessagesGetMessageReactionsList(ctx context.Context, in *mtproto.TLMessagesGetMessageReactionsList) (*mtproto.Messages_MessageReactionsList, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetMessageReactionsList(in)
}

func (s *Service) MessagesSetChatAvailableReactions(ctx context.Context, in *mtproto.TLMessagesSetChatAvailableReactions) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesSetChatAvailableReactions(in)
}

func (s *Service) MessagesGetAvailableReactions(ctx context.Context, in *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetAvailableReactions(in)
}

func (s *Service) MessagesSetDefaultReaction(ctx context.Context, in *mtproto.TLMessagesSetDefaultReaction) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesSetDefaultReaction(in)
}

func (s *Service) MessagesGetUnreadReactions(ctx context.Context, in *mtproto.TLMessagesGetUnreadReactions) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetUnreadReactions(in)
}

func (s *Service) MessagesReadReactions(ctx context.Context, in *mtproto.TLMessagesReadReactions) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesReadReactions(in)
}

func (s *Service) MessagesReportReaction(ctx context.Context, in *mtproto.TLMessagesReportReaction) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesReportReaction(in)
}

func (s *Service) MessagesGetTopReactions(ctx context.Context, in *mtproto.TLMessagesGetTopReactions) (*mtproto.Messages_Reactions, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetTopReactions(in)
}

func (s *Service) MessagesGetRecentReactions(ctx context.Context, in *mtproto.TLMessagesGetRecentReactions) (*mtproto.Messages_Reactions, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetRecentReactions(in)
}

func (s *Service) MessagesClearRecentReactions(ctx context.Context, in *mtproto.TLMessagesClearRecentReactions) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesClearRecentReactions(in)
}

func (s *Service) MessagesSendPaidReaction(ctx context.Context, in *mtproto.TLMessagesSendPaidReaction) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesSendPaidReaction(in)
}

func (s *Service) MessagesTogglePaidReactionPrivacy(ctx context.Context, in *mtproto.TLMessagesTogglePaidReactionPrivacy) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesTogglePaidReactionPrivacy(in)
}

func (s *Service) MessagesGetPaidReactionPrivacy(ctx context.Context, in *mtproto.TLMessagesGetPaidReactionPrivacy) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetPaidReactionPrivacy(in)
}
