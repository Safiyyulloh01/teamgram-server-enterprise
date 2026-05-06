package core

import (
	"github.com/teamgram/proto/mtproto"
)

func (c *StickersCore) MessagesGetAllStickers(in *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	return mtproto.MakeTLMessagesAllStickers(&mtproto.Messages_AllStickers{
		Hash: 0,
		Sets: []*mtproto.StickerSet{},
	}).To_Messages_AllStickers(), nil
}

func (c *StickersCore) MessagesGetStickers(in *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error) {
	return mtproto.MakeTLMessagesStickers(&mtproto.Messages_Stickers{
		Hash:     0,
		Stickers: []*mtproto.Document{},
	}).To_Messages_Stickers(), nil
}

func (c *StickersCore) MessagesGetStickerSet(in *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	return mtproto.MakeTLMessagesStickerSetNotModified(&mtproto.Messages_StickerSet{}).To_Messages_StickerSet(), nil
}

func (c *StickersCore) MessagesInstallStickerSet(in *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error) {
	return mtproto.MakeTLMessagesStickerSetInstallResultArchive(&mtproto.Messages_StickerSetInstallResult{
		Sets: []*mtproto.StickerSetCovered{},
	}).To_Messages_StickerSetInstallResult(), nil
}

func (c *StickersCore) MessagesUninstallStickerSet(in *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *StickersCore) MessagesReorderStickerSets(in *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *StickersCore) MessagesGetFeaturedStickers(in *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	return mtproto.MakeTLMessagesFeaturedStickers(&mtproto.Messages_FeaturedStickers{
		Hash:   0,
		Sets:   []*mtproto.StickerSetCovered{},
		Unread: []int64{},
	}).To_Messages_FeaturedStickers(), nil
}

func (c *StickersCore) MessagesReadFeaturedStickers(in *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *StickersCore) MessagesGetRecentStickers(in *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	return mtproto.MakeTLMessagesRecentStickers(&mtproto.Messages_RecentStickers{
		Hash:     0,
		Packs:    []*mtproto.StickerPack{},
		Stickers: []*mtproto.Document{},
		Dates:    []int32{},
	}).To_Messages_RecentStickers(), nil
}

func (c *StickersCore) MessagesSaveRecentSticker(in *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *StickersCore) MessagesClearRecentStickers(in *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *StickersCore) MessagesGetArchivedStickers(in *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	return mtproto.MakeTLMessagesArchivedStickers(&mtproto.Messages_ArchivedStickers{
		Count: 0,
		Sets:  []*mtproto.StickerSetCovered{},
	}).To_Messages_ArchivedStickers(), nil
}

func (c *StickersCore) MessagesGetMaskStickers(in *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	return mtproto.MakeTLMessagesAllStickers(&mtproto.Messages_AllStickers{
		Hash: 0,
		Sets: []*mtproto.StickerSet{},
	}).To_Messages_AllStickers(), nil
}

func (c *StickersCore) MessagesGetAttachedStickers(in *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector_StickerSetCovered, error) {
	return &mtproto.Vector_StickerSetCovered{
		Datas: []*mtproto.StickerSetCovered{},
	}, nil
}

func (c *StickersCore) MessagesGetFavedStickers(in *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	return mtproto.MakeTLMessagesFavedStickers(&mtproto.Messages_FavedStickers{
		Hash:     0,
		Packs:    []*mtproto.StickerPack{},
		Stickers: []*mtproto.Document{},
	}).To_Messages_FavedStickers(), nil
}

func (c *StickersCore) MessagesFaveSticker(in *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *StickersCore) MessagesSearchStickerSets(in *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	return mtproto.MakeTLMessagesFoundStickerSets(&mtproto.Messages_FoundStickerSets{
		Hash: 0,
		Sets: []*mtproto.StickerSetCovered{},
	}).To_Messages_FoundStickerSets(), nil
}

func (c *StickersCore) MessagesToggleStickerSets(in *mtproto.TLMessagesToggleStickerSets) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *StickersCore) MessagesGetOldFeaturedStickers(in *mtproto.TLMessagesGetOldFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	return mtproto.MakeTLMessagesFeaturedStickers(&mtproto.Messages_FeaturedStickers{
		Hash:   0,
		Sets:   []*mtproto.StickerSetCovered{},
		Unread: []int64{},
	}).To_Messages_FeaturedStickers(), nil
}

func (c *StickersCore) MessagesSearchEmojiStickerSets(in *mtproto.TLMessagesSearchEmojiStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	return mtproto.MakeTLMessagesFoundStickerSets(&mtproto.Messages_FoundStickerSets{
		Hash: 0,
		Sets: []*mtproto.StickerSetCovered{},
	}).To_Messages_FoundStickerSets(), nil
}

func (c *StickersCore) MessagesGetMyStickers(in *mtproto.TLMessagesGetMyStickers) (*mtproto.Messages_MyStickers, error) {
	return mtproto.MakeTLMessagesMyStickers(&mtproto.Messages_MyStickers{
		Count: 0,
		Sets:  []*mtproto.StickerSetCovered{},
	}).To_Messages_MyStickers(), nil
}

func (c *StickersCore) MessagesSearchStickers(in *mtproto.TLMessagesSearchStickers) (*mtproto.Messages_FoundStickers, error) {
	return mtproto.MakeTLMessagesFoundStickers(&mtproto.Messages_FoundStickers{
		Hash:     0,
		Stickers: []*mtproto.Document{},
	}).To_Messages_FoundStickers(), nil
}

func (c *StickersCore) StickersCreateStickerSet(in *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error) {
	return mtproto.MakeTLMessagesStickerSetNotModified(&mtproto.Messages_StickerSet{}).To_Messages_StickerSet(), nil
}

func (c *StickersCore) StickersRemoveStickerFromSet(in *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
	return mtproto.MakeTLMessagesStickerSetNotModified(&mtproto.Messages_StickerSet{}).To_Messages_StickerSet(), nil
}

func (c *StickersCore) StickersChangeStickerPosition(in *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error) {
	return mtproto.MakeTLMessagesStickerSetNotModified(&mtproto.Messages_StickerSet{}).To_Messages_StickerSet(), nil
}

func (c *StickersCore) StickersAddStickerToSet(in *mtproto.TLStickersAddStickerToSet) (*mtproto.Messages_StickerSet, error) {
	return mtproto.MakeTLMessagesStickerSetNotModified(&mtproto.Messages_StickerSet{}).To_Messages_StickerSet(), nil
}

func (c *StickersCore) StickersSetStickerSetThumb(in *mtproto.TLStickersSetStickerSetThumb) (*mtproto.Messages_StickerSet, error) {
	return mtproto.MakeTLMessagesStickerSetNotModified(&mtproto.Messages_StickerSet{}).To_Messages_StickerSet(), nil
}

func (c *StickersCore) StickersCheckShortName(in *mtproto.TLStickersCheckShortName) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *StickersCore) StickersSuggestShortName(in *mtproto.TLStickersSuggestShortName) (*mtproto.Stickers_SuggestedShortName, error) {
	return mtproto.MakeTLStickersSuggestedShortName(&mtproto.Stickers_SuggestedShortName{
		ShortName: "",
	}).To_Stickers_SuggestedShortName(), nil
}

func (c *StickersCore) StickersChangeSticker(in *mtproto.TLStickersChangeSticker) (*mtproto.Messages_StickerSet, error) {
	return mtproto.MakeTLMessagesStickerSetNotModified(&mtproto.Messages_StickerSet{}).To_Messages_StickerSet(), nil
}
