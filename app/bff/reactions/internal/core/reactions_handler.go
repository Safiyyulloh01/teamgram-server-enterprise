package core

import (
	"github.com/teamgram/proto/mtproto"
	"time"
)

func (c *ReactionsCore) MessagesGetAvailableReactions(in *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error) {
	return mtproto.MakeTLMessagesAvailableReactions(&mtproto.Messages_AvailableReactions{
		Hash:      0,
		Reactions: []*mtproto.AvailableReaction{},
	}).To_Messages_AvailableReactions(), nil
}

func (c *ReactionsCore) MessagesSendReaction(in *mtproto.TLMessagesSendReaction) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ReactionsCore) MessagesGetMessagesReactions(in *mtproto.TLMessagesGetMessagesReactions) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ReactionsCore) MessagesGetMessageReactionsList(in *mtproto.TLMessagesGetMessageReactionsList) (*mtproto.Messages_MessageReactionsList, error) {
	return mtproto.MakeTLMessagesMessageReactionsList(&mtproto.Messages_MessageReactionsList{
		Count:     0,
		Reactions: []*mtproto.MessagePeerReaction{},
		Users:     []*mtproto.User{},
		Chats:     []*mtproto.Chat{},
	}).To_Messages_MessageReactionsList(), nil
}

func (c *ReactionsCore) MessagesSetChatAvailableReactions(in *mtproto.TLMessagesSetChatAvailableReactions) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ReactionsCore) MessagesSetDefaultReaction(in *mtproto.TLMessagesSetDefaultReaction) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ReactionsCore) MessagesGetUnreadReactions(in *mtproto.TLMessagesGetUnreadReactions) (*mtproto.Messages_Messages, error) {
	return mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages(), nil
}

func (c *ReactionsCore) MessagesReadReactions(in *mtproto.TLMessagesReadReactions) (*mtproto.Messages_AffectedHistory, error) {
	return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      1,
		PtsCount: 0,
		Offset:   0,
	}).To_Messages_AffectedHistory(), nil
}

func (c *ReactionsCore) MessagesReportReaction(in *mtproto.TLMessagesReportReaction) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ReactionsCore) MessagesGetTopReactions(in *mtproto.TLMessagesGetTopReactions) (*mtproto.Messages_Reactions, error) {
	return mtproto.MakeTLMessagesReactions(&mtproto.Messages_Reactions{
		Reactions: []*mtproto.Reaction{},
		Hash:      0,
	}).To_Messages_Reactions(), nil
}

func (c *ReactionsCore) MessagesGetRecentReactions(in *mtproto.TLMessagesGetRecentReactions) (*mtproto.Messages_Reactions, error) {
	return mtproto.MakeTLMessagesReactions(&mtproto.Messages_Reactions{
		Reactions: []*mtproto.Reaction{},
		Hash:      0,
	}).To_Messages_Reactions(), nil
}

func (c *ReactionsCore) MessagesClearRecentReactions(in *mtproto.TLMessagesClearRecentReactions) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ReactionsCore) MessagesSendPaidReaction(in *mtproto.TLMessagesSendPaidReaction) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ReactionsCore) MessagesTogglePaidReactionPrivacy(in *mtproto.TLMessagesTogglePaidReactionPrivacy) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ReactionsCore) MessagesGetPaidReactionPrivacy(in *mtproto.TLMessagesGetPaidReactionPrivacy) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}
