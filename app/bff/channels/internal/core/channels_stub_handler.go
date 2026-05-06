package core

import (
	"github.com/teamgram/proto/mtproto"
	"time"
)

// ChannelsReadHistory marks channel history as read
func (c *ChannelsCore) ChannelsReadHistory(in *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

// ChannelsDeleteMessages deletes messages in a channel
func (c *ChannelsCore) ChannelsDeleteMessages(in *mtproto.TLChannelsDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      1,
		PtsCount: int32(len(in.GetId())),
	}).To_Messages_AffectedMessages(), nil
}

// ChannelsGetMessages gets messages from a channel
func (c *ChannelsCore) ChannelsGetMessages(in *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	return mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages(), nil
}

// ChannelsGetParticipant gets a specific channel participant
func (c *ChannelsCore) ChannelsGetParticipant(in *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	participant := mtproto.MakeTLChannelParticipant(&mtproto.ChannelParticipant{
		UserId: in.GetUserId().GetUserId(),
		Date:   int32(time.Now().Unix()),
	}).To_ChannelParticipant()

	return mtproto.MakeTLChannelsChannelParticipant(&mtproto.Channels_ChannelParticipant{
		Participant: participant,
		Chats:       []*mtproto.Chat{},
		Users:       []*mtproto.User{},
	}).To_Channels_ChannelParticipant(), nil
}

// ChannelsEditPhoto edits channel photo
func (c *ChannelsCore) ChannelsEditPhoto(in *mtproto.TLChannelsEditPhoto) (*mtproto.Updates, error) {
	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}

// ChannelsExportMessageLink exports a message link
func (c *ChannelsCore) ChannelsExportMessageLink(in *mtproto.TLChannelsExportMessageLink) (*mtproto.ExportedMessageLink, error) {
	return mtproto.MakeTLExportedMessageLink(&mtproto.ExportedMessageLink{
		Link: "",
	}).To_ExportedMessageLink(), nil
}

// ChannelsGetAdminedPublicChannels gets admin public channels
func (c *ChannelsCore) ChannelsGetAdminedPublicChannels(in *mtproto.TLChannelsGetAdminedPublicChannels) (*mtproto.Messages_Chats, error) {
	return mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
	}).To_Messages_Chats(), nil
}

// ChannelsGetAdminLog gets admin log
func (c *ChannelsCore) ChannelsGetAdminLog(in *mtproto.TLChannelsGetAdminLog) (*mtproto.Channels_AdminLogResults, error) {
	return mtproto.MakeTLChannelsAdminLogResults(&mtproto.Channels_AdminLogResults{
		Events: []*mtproto.ChannelAdminLogEvent{},
		Chats:  []*mtproto.Chat{},
		Users:  []*mtproto.User{},
	}).To_Channels_AdminLogResults(), nil
}

// ChannelsSetStickers sets channel sticker pack
func (c *ChannelsCore) ChannelsSetStickers(in *mtproto.TLChannelsSetStickers) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

// ChannelsReadMessageContents marks message contents as read
func (c *ChannelsCore) ChannelsReadMessageContents(in *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

// ChannelsDeleteHistory9BAA9647 deletes channel history
func (c *ChannelsCore) ChannelsDeleteHistory9BAA9647(in *mtproto.TLChannelsDeleteHistory9BAA9647) (*mtproto.Updates, error) {
	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}

// ChannelsTogglePreHistoryHidden toggles pre-history hidden
func (c *ChannelsCore) ChannelsTogglePreHistoryHidden(in *mtproto.TLChannelsTogglePreHistoryHidden) (*mtproto.Updates, error) {
	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}

// ChannelsGetGroupsForDiscussion gets groups for discussion
func (c *ChannelsCore) ChannelsGetGroupsForDiscussion(in *mtproto.TLChannelsGetGroupsForDiscussion) (*mtproto.Messages_Chats, error) {
	return mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
	}).To_Messages_Chats(), nil
}

// ChannelsSetDiscussionGroup sets discussion group
func (c *ChannelsCore) ChannelsSetDiscussionGroup(in *mtproto.TLChannelsSetDiscussionGroup) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

// ChannelsEditLocation edits channel location
func (c *ChannelsCore) ChannelsEditLocation(in *mtproto.TLChannelsEditLocation) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

// ChannelsToggleSlowMode toggles slow mode
func (c *ChannelsCore) ChannelsToggleSlowMode(in *mtproto.TLChannelsToggleSlowMode) (*mtproto.Updates, error) {
	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}

// ChannelsGetInactiveChannels gets inactive channels
func (c *ChannelsCore) ChannelsGetInactiveChannels(in *mtproto.TLChannelsGetInactiveChannels) (*mtproto.Messages_InactiveChats, error) {
	return mtproto.MakeTLMessagesInactiveChats(&mtproto.Messages_InactiveChats{
		Dates: []int32{},
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}).To_Messages_InactiveChats(), nil
}

// ChannelsDeleteParticipantHistory deletes participant history
func (c *ChannelsCore) ChannelsDeleteParticipantHistory(in *mtproto.TLChannelsDeleteParticipantHistory) (*mtproto.Messages_AffectedHistory, error) {
	return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      1,
		PtsCount: 0,
		Offset:   0,
	}).To_Messages_AffectedHistory(), nil
}

// ChannelsToggleParticipantsHidden toggles participants hidden
func (c *ChannelsCore) ChannelsToggleParticipantsHidden(in *mtproto.TLChannelsToggleParticipantsHidden) (*mtproto.Updates, error) {
	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}

// ChannelsEditCreator edits channel creator
func (c *ChannelsCore) ChannelsEditCreator(in *mtproto.TLChannelsEditCreator) (*mtproto.Updates, error) {
	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}

// ChannelsGetFutureCreatorAfterLeave gets future creator
func (c *ChannelsCore) ChannelsGetFutureCreatorAfterLeave(in *mtproto.TLChannelsGetFutureCreatorAfterLeave) (*mtproto.User, error) {
	return mtproto.MakeTLUserEmpty(&mtproto.User{}).To_User(), nil
}

// ChannelsDeleteHistoryAF369D42 deletes channel history (alt version)
func (c *ChannelsCore) ChannelsDeleteHistoryAF369D42(in *mtproto.TLChannelsDeleteHistoryAF369D42) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
