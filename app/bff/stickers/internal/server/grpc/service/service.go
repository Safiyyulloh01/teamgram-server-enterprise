package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/stickers/internal/core"
	"github.com/teamgram/teamgram-server/app/bff/stickers/internal/svc"
)

type Service struct {
	svcCtx *svc.ServiceContext
}

func New(ctx *svc.ServiceContext) *Service {
	return &Service{
		svcCtx: ctx,
		// Logger removed,
	}
}

var _ mtproto.RPCStickersServer = (*Service)(nil)

func (s *Service) MessagesGetStickers(ctx context.Context, in *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetStickers(in)
}

func (s *Service) MessagesGetAllStickers(ctx context.Context, in *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetAllStickers(in)
}

func (s *Service) MessagesGetStickerSet(ctx context.Context, in *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetStickerSet(in)
}

func (s *Service) MessagesInstallStickerSet(ctx context.Context, in *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesInstallStickerSet(in)
}

func (s *Service) MessagesUninstallStickerSet(ctx context.Context, in *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesUninstallStickerSet(in)
}

func (s *Service) MessagesReorderStickerSets(ctx context.Context, in *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesReorderStickerSets(in)
}

func (s *Service) MessagesGetFeaturedStickers(ctx context.Context, in *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetFeaturedStickers(in)
}

func (s *Service) MessagesReadFeaturedStickers(ctx context.Context, in *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesReadFeaturedStickers(in)
}

func (s *Service) MessagesGetRecentStickers(ctx context.Context, in *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetRecentStickers(in)
}

func (s *Service) MessagesSaveRecentSticker(ctx context.Context, in *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesSaveRecentSticker(in)
}

func (s *Service) MessagesClearRecentStickers(ctx context.Context, in *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesClearRecentStickers(in)
}

func (s *Service) MessagesGetArchivedStickers(ctx context.Context, in *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetArchivedStickers(in)
}

func (s *Service) MessagesGetMaskStickers(ctx context.Context, in *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetMaskStickers(in)
}

func (s *Service) MessagesGetAttachedStickers(ctx context.Context, in *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector_StickerSetCovered, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetAttachedStickers(in)
}

func (s *Service) MessagesGetFavedStickers(ctx context.Context, in *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetFavedStickers(in)
}

func (s *Service) MessagesFaveSticker(ctx context.Context, in *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesFaveSticker(in)
}

func (s *Service) MessagesSearchStickerSets(ctx context.Context, in *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesSearchStickerSets(in)
}

func (s *Service) MessagesToggleStickerSets(ctx context.Context, in *mtproto.TLMessagesToggleStickerSets) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesToggleStickerSets(in)
}

func (s *Service) MessagesGetOldFeaturedStickers(ctx context.Context, in *mtproto.TLMessagesGetOldFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetOldFeaturedStickers(in)
}

func (s *Service) MessagesSearchEmojiStickerSets(ctx context.Context, in *mtproto.TLMessagesSearchEmojiStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesSearchEmojiStickerSets(in)
}

func (s *Service) MessagesGetMyStickers(ctx context.Context, in *mtproto.TLMessagesGetMyStickers) (*mtproto.Messages_MyStickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesGetMyStickers(in)
}

func (s *Service) MessagesSearchStickers(ctx context.Context, in *mtproto.TLMessagesSearchStickers) (*mtproto.Messages_FoundStickers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.MessagesSearchStickers(in)
}

func (s *Service) StickersCreateStickerSet(ctx context.Context, in *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersCreateStickerSet(in)
}

func (s *Service) StickersRemoveStickerFromSet(ctx context.Context, in *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersRemoveStickerFromSet(in)
}

func (s *Service) StickersChangeStickerPosition(ctx context.Context, in *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersChangeStickerPosition(in)
}

func (s *Service) StickersAddStickerToSet(ctx context.Context, in *mtproto.TLStickersAddStickerToSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersAddStickerToSet(in)
}

func (s *Service) StickersSetStickerSetThumb(ctx context.Context, in *mtproto.TLStickersSetStickerSetThumb) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersSetStickerSetThumb(in)
}

func (s *Service) StickersCheckShortName(ctx context.Context, in *mtproto.TLStickersCheckShortName) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersCheckShortName(in)
}

func (s *Service) StickersSuggestShortName(ctx context.Context, in *mtproto.TLStickersSuggestShortName) (*mtproto.Stickers_SuggestedShortName, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersSuggestShortName(in)
}

func (s *Service) StickersChangeSticker(ctx context.Context, in *mtproto.TLStickersChangeSticker) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersChangeSticker(in)
}

func (s *Service) StickersDeleteStickerSet(ctx context.Context, in *mtproto.TLStickersDeleteStickerSet) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersDeleteStickerSet(in)
}

func (s *Service) StickersRenameStickerSet(ctx context.Context, in *mtproto.TLStickersRenameStickerSet) (*mtproto.Messages_StickerSet, error) {
	c := core.New(ctx, s.svcCtx)
	return c.StickersRenameStickerSet(in)
}
