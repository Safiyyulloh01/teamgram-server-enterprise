package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
	"time"
)

// MessagesGetAvailableReactions returns available message reactions
func (c *ReactionsCore) MessagesGetAvailableReactions(in *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error) {
	reactions := []*mtproto.AvailableReaction{
		makeReaction("👍", "\U0001F44D", 1),
		makeReaction("👎", "\U0001F44E", 2),
		makeReaction("❤", "\u2764\uFE0F", 3),
		makeReaction("🔥", "\U0001F525", 4),
		makeReaction("\U0001F60A", "\U0001F60A", 5),
		makeReaction("😁", "\U0001F601", 6),
		makeReaction("😮", "\U0001F62E", 7),
		makeReaction("🎉", "\U0001F389", 8),
		makeReaction("💩", "\U0001F4A9", 9),
		makeReaction("\U0001F622", "\U0001F622", 10),
		makeReaction("🤩", "\U0001F929", 11),
		makeReaction("🤔", "\U0001F914", 12),
		makeReaction("👏", "\U0001F44F", 13),
		makeReaction("\U0001F62D", "\U0001F62D", 14),
		makeReaction("😈", "\U0001F608", 15),
		makeReaction("🎃", "\U0001F383", 16),
		makeReaction("💯", "\U0001F4AF", 17),
		makeReaction("🐳", "\U0001F433", 18),
		makeReaction("\U0001F937", "\U0001F937", 19),
		makeReaction("\U0001F496", "\U0001F496", 20),
		makeReaction("💘", "\U0001F498", 21),
	}

	return mtproto.MakeTLMessagesAvailableReactions(&mtproto.Messages_AvailableReactions{
		Hash:      int32(crypto.GenerateHash(reactions)),
		Reactions: reactions,
	}).To_Messages_AvailableReactions(), nil
}

func makeReaction(emoticon, staticIcon string, idx int) *mtproto.AvailableReaction {
	return mtproto.MakeTLAvailableReaction(&mtproto.AvailableReaction{
		Reaction:         emoticon,
		Title:            emoticon,
		StaticIcon:       mtproto.MakeTLDocumentEmpty(&mtproto.Document{}).To_Document(),
		AppearAnimation:  mtproto.MakeTLDocumentEmpty(&mtproto.Document{}).To_Document(),
		SelectAnimation:  mtproto.MakeTLDocumentEmpty(&mtproto.Document{}).To_Document(),
		ActivateAnimation: mtproto.MakeTLDocumentEmpty(&mtproto.Document{}).To_Document(),
		EffectAnimation:  mtproto.MakeTLDocumentEmpty(&mtproto.Document{}).To_Document(),
		AroundAnimation:  mtproto.MakeTLDocumentEmpty(&mtproto.Document{}).To_Document(),
		CenterIcon:       mtproto.MakeTLDocumentEmpty(&mtproto.Document{}).To_Document(),
	}).To_AvailableReaction()
}

// MessagesSendReaction sends a reaction to a message
func (c *ReactionsCore) MessagesSendReaction(in *mtproto.TLMessagesSendReaction) (*mtproto.Updates, error) {
	peer := in.GetPeer()
	msgId := in.GetMsgId()
	reaction := in.GetReaction()
	big := mtproto.FromBool(in.GetBig())

	peerType, peerId := mtproto.ResolvePeer(peer)

	updatesHelper := mtproto.MakeUpdatesHelper()
	updateReaction := mtproto.MakeTLUpdateMessageReaction(&mtproto.Update{
		UserId:    c.MD.UserId,
		Peer:      peer,
		MsgId:     msgId,
		Reactions: reaction,
	})
	updatesHelper.AddUpdate(updateReaction)

	return updatesHelper.ToUpdates(), nil
}

// MessagesGetMessagesReactions gets reactions for specific messages
func (c *ReactionsCore) MessagesGetMessagesReactions(in *mtproto.TLMessagesGetMessagesReactions) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdatesEmpty(&mtproto.Updates{}).To_Updates(), nil
}

// MessagesGetMessageReactionsList gets reactions list for a message
func (c *ReactionsCore) MessagesGetMessageReactionsList(in *mtproto.TLMessagesGetMessageReactionsList) (*mtproto.Messages_MessageReactionsList, error) {
	return mtproto.MakeTLMessagesMessageReactionsList(&mtproto.Messages_MessageReactionsList{
		Count: 0,
		Reactions: []*mtproto.MessageUserReaction{},
		Users: []*mtproto.User{},
		Chats: []*mtproto.Chat{},
	}).To_Messages_MessageReactionsList(), nil
}

// MessagesSetChatAvailableReactions sets available reactions for a chat
func (c *ReactionsCore) MessagesSetChatAvailableReactions(in *mtproto.TLMessagesSetChatAvailableReactions) (*mtproto.Updates, error) {
	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}

// MessagesSetDefaultReaction sets a user's default reaction
func (c *ReactionsCore) MessagesSetDefaultReaction(in *mtproto.TLMessagesSetDefaultReaction) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

// MessagesGetUnreadReactions gets unread reactions
func (c *ReactionsCore) MessagesGetUnreadReactions(in *mtproto.TLMessagesGetUnreadReactions) (*mtproto.Messages_Messages, error) {
	return mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages(), nil
}

// MessagesReadReactions marks reactions as read
func (c *ReactionsCore) MessagesReadReactions(in *mtproto.TLMessagesReadReactions) (*mtproto.Messages_AffectedHistory, error) {
	return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      1,
		PtsCount: 0,
		Offset:   0,
	}).To_Messages_AffectedHistory(), nil
}

// MessagesReportReaction reports a reaction
func (c *ReactionsCore) MessagesReportReaction(in *mtproto.TLMessagesReportReaction) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

// MessagesGetTopReactions returns top reactions
func (c *ReactionsCore) MessagesGetTopReactions(in *mtproto.TLMessagesGetTopReactions) (*mtproto.Messages_Reactions, error) {
	return mtproto.MakeTLMessagesReactions(&mtproto.Messages_Reactions{
		Reactions: []*mtproto.Reaction{},
		Period:    0,
		Hash:      0,
	}).To_Messages_Reactions(), nil
}

// MessagesGetRecentReactions returns recent reactions
func (c *ReactionsCore) MessagesGetRecentReactions(in *mtproto.TLMessagesGetRecentReactions) (*mtproto.Messages_Reactions, error) {
	return mtproto.MakeTLMessagesReactions(&mtproto.Messages_Reactions{
		Reactions: []*mtproto.Reaction{},
		Period:    0,
		Hash:      0,
	}).To_Messages_Reactions(), nil
}

// MessagesClearRecentReactions clears recent reactions
func (c *ReactionsCore) MessagesClearRecentReactions(in *mtproto.TLMessagesClearRecentReactions) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

// MessagesSendPaidReaction sends a paid reaction
func (c *ReactionsCore) MessagesSendPaidReaction(in *mtproto.TLMessagesSendPaidReaction) (*mtproto.Updates, error) {
	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}

// MessagesTogglePaidReactionPrivacy toggles paid reaction privacy
func (c *ReactionsCore) MessagesTogglePaidReactionPrivacy(in *mtproto.TLMessagesTogglePaidReactionPrivacy) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

// MessagesGetPaidReactionPrivacy gets paid reaction privacy
func (c *ReactionsCore) MessagesGetPaidReactionPrivacy(in *mtproto.TLMessagesGetPaidReactionPrivacy) (*mtproto.Updates, error) {
	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}
