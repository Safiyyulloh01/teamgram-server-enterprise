package core

import (
	"github.com/teamgram/proto/mtproto"
)

func (c *ChannelsCore) ChannelsCreateChannel(in *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsJoinChannel(in *mtproto.TLChannelsJoinChannel) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsLeaveChannel(in *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsGetChannels(in *mtproto.TLChannelsGetChannels) (*mtproto.Messages_Chats, error) {
	return mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
	}).To_Messages_Chats(), nil
}

func (c *ChannelsCore) ChannelsGetFullChannel(in *mtproto.TLChannelsGetFullChannel) (*mtproto.Messages_ChatFull, error) {
	return mtproto.MakeTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: &mtproto.ChatFull{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_ChatFull(), nil
}

func (c *ChannelsCore) ChannelsEditTitle(in *mtproto.TLChannelsEditTitle) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsDeleteChannel(in *mtproto.TLChannelsDeleteChannel) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsGetParticipants(in *mtproto.TLChannelsGetParticipants) (*mtproto.Channels_ChannelParticipants, error) {
	return mtproto.MakeTLChannelsChannelParticipants(&mtproto.Channels_ChannelParticipants{
		Count:        0,
		Participants: []*mtproto.ChannelParticipant{},
		Chats:        []*mtproto.Chat{},
		Users:        []*mtproto.User{},
	}).To_Channels_ChannelParticipants(), nil
}

func (c *ChannelsCore) ChannelsEditAdmin(in *mtproto.TLChannelsEditAdmin) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsEditBanned(in *mtproto.TLChannelsEditBanned) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsToggleSignatures(in *mtproto.TLChannelsToggleSignatures) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsInviteToChannel199F3A6C(in *mtproto.TLChannelsInviteToChannel199F3A6C) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsInviteToChannelC9E33D54(in *mtproto.TLChannelsInviteToChannelC9E33D54) (*mtproto.Messages_InvitedUsers, error) {
	return mtproto.MakeTLMessagesInvitedUsers(&mtproto.Messages_InvitedUsers{
		Updates:         &mtproto.Updates{},
		MissingInvitees: []*mtproto.MissingInvitee{},
	}).To_Messages_InvitedUsers(), nil
}

func (c *ChannelsCore) ChannelsGetParticipant(in *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	return mtproto.MakeTLChannelsChannelParticipant(&mtproto.Channels_ChannelParticipant{
		Participant: &mtproto.ChannelParticipant{},
		Chats:       []*mtproto.Chat{},
		Users:       []*mtproto.User{},
	}).To_Channels_ChannelParticipant(), nil
}

func (c *ChannelsCore) ChannelsReadHistory(in *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsDeleteMessages(in *mtproto.TLChannelsDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      1,
		PtsCount: int32(len(in.GetId())),
	}).To_Messages_AffectedMessages(), nil
}

func (c *ChannelsCore) ChannelsGetMessages(in *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	return mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages(), nil
}

func (c *ChannelsCore) ChannelsEditPhoto(in *mtproto.TLChannelsEditPhoto) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsExportMessageLink(in *mtproto.TLChannelsExportMessageLink) (*mtproto.ExportedMessageLink, error) {
	return mtproto.MakeTLExportedMessageLink(&mtproto.ExportedMessageLink{
		Link: "",
	}).To_ExportedMessageLink(), nil
}

func (c *ChannelsCore) ChannelsGetAdminedPublicChannels(in *mtproto.TLChannelsGetAdminedPublicChannels) (*mtproto.Messages_Chats, error) {
	return mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
	}).To_Messages_Chats(), nil
}

func (c *ChannelsCore) ChannelsGetAdminLog(in *mtproto.TLChannelsGetAdminLog) (*mtproto.Channels_AdminLogResults, error) {
	return mtproto.MakeTLChannelsAdminLogResults(&mtproto.Channels_AdminLogResults{
		Events: []*mtproto.ChannelAdminLogEvent{},
		Chats:  []*mtproto.Chat{},
		Users:  []*mtproto.User{},
	}).To_Channels_AdminLogResults(), nil
}

func (c *ChannelsCore) ChannelsSetStickers(in *mtproto.TLChannelsSetStickers) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsReadMessageContents(in *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsDeleteHistory9BAA9647(in *mtproto.TLChannelsDeleteHistory9BAA9647) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsTogglePreHistoryHidden(in *mtproto.TLChannelsTogglePreHistoryHidden) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsGetGroupsForDiscussion(in *mtproto.TLChannelsGetGroupsForDiscussion) (*mtproto.Messages_Chats, error) {
	return mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
	}).To_Messages_Chats(), nil
}

func (c *ChannelsCore) ChannelsSetDiscussionGroup(in *mtproto.TLChannelsSetDiscussionGroup) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsEditLocation(in *mtproto.TLChannelsEditLocation) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsToggleSlowMode(in *mtproto.TLChannelsToggleSlowMode) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsGetInactiveChannels(in *mtproto.TLChannelsGetInactiveChannels) (*mtproto.Messages_InactiveChats, error) {
	return mtproto.MakeTLMessagesInactiveChats(&mtproto.Messages_InactiveChats{
		Dates: []int32{},
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}).To_Messages_InactiveChats(), nil
}

func (c *ChannelsCore) ChannelsDeleteParticipantHistory(in *mtproto.TLChannelsDeleteParticipantHistory) (*mtproto.Messages_AffectedHistory, error) {
	return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      1,
		PtsCount: 0,
		Offset:   0,
	}).To_Messages_AffectedHistory(), nil
}

func (c *ChannelsCore) ChannelsToggleParticipantsHidden(in *mtproto.TLChannelsToggleParticipantsHidden) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsEditCreator(in *mtproto.TLChannelsEditCreator) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    0,
		Seq:     0,
	}).To_Updates(), nil
}

func (c *ChannelsCore) ChannelsGetFutureCreatorAfterLeave(in *mtproto.TLChannelsGetFutureCreatorAfterLeave) (*mtproto.User, error) {
	return mtproto.MakeTLUserEmpty(&mtproto.User{}).To_User(), nil
}

func (c *ChannelsCore) ChannelsDeleteHistoryAF369D42(in *mtproto.TLChannelsDeleteHistoryAF369D42) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
