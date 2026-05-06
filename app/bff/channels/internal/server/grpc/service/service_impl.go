package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/channels/internal/core"
)

// ChannelsReadHistory implements RPCChannelsServer
func (s *Service) ChannelsReadHistory(ctx context.Context, in *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsReadHistory(in)
}

// ChannelsDeleteMessages implements RPCChannelsServer
func (s *Service) ChannelsDeleteMessages(ctx context.Context, in *mtproto.TLChannelsDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsDeleteMessages(in)
}

// ChannelsGetMessages implements RPCChannelsServer
func (s *Service) ChannelsGetMessages(ctx context.Context, in *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetMessages(in)
}

// ChannelsGetParticipants implements RPCChannelsServer
func (s *Service) ChannelsGetParticipants(ctx context.Context, in *mtproto.TLChannelsGetParticipants) (*mtproto.Channels_ChannelParticipants, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetParticipants(in)
}

// ChannelsGetParticipant implements RPCChannelsServer
func (s *Service) ChannelsGetParticipant(ctx context.Context, in *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetParticipant(in)
}

// ChannelsGetChannels implements RPCChannelsServer
func (s *Service) ChannelsGetChannels(ctx context.Context, in *mtproto.TLChannelsGetChannels) (*mtproto.Messages_Chats, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetChannels(in)
}

// ChannelsGetFullChannel implements RPCChannelsServer
func (s *Service) ChannelsGetFullChannel(ctx context.Context, in *mtproto.TLChannelsGetFullChannel) (*mtproto.Messages_ChatFull, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetFullChannel(in)
}

// ChannelsCreateChannel implements RPCChannelsServer
func (s *Service) ChannelsCreateChannel(ctx context.Context, in *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsCreateChannel(in)
}

// ChannelsEditAdmin implements RPCChannelsServer
func (s *Service) ChannelsEditAdmin(ctx context.Context, in *mtproto.TLChannelsEditAdmin) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsEditAdmin(in)
}

// ChannelsEditTitle implements RPCChannelsServer
func (s *Service) ChannelsEditTitle(ctx context.Context, in *mtproto.TLChannelsEditTitle) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsEditTitle(in)
}

// ChannelsEditPhoto implements RPCChannelsServer
func (s *Service) ChannelsEditPhoto(ctx context.Context, in *mtproto.TLChannelsEditPhoto) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsEditPhoto(in)
}

// ChannelsJoinChannel implements RPCChannelsServer
func (s *Service) ChannelsJoinChannel(ctx context.Context, in *mtproto.TLChannelsJoinChannel) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsJoinChannel(in)
}

// ChannelsLeaveChannel implements RPCChannelsServer
func (s *Service) ChannelsLeaveChannel(ctx context.Context, in *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsLeaveChannel(in)
}

// ChannelsInviteToChannelC9E33D54 implements RPCChannelsServer
func (s *Service) ChannelsInviteToChannelC9E33D54(ctx context.Context, in *mtproto.TLChannelsInviteToChannelC9E33D54) (*mtproto.Messages_InvitedUsers, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsInviteToChannelC9E33D54(in)
}

// ChannelsDeleteChannel implements RPCChannelsServer
func (s *Service) ChannelsDeleteChannel(ctx context.Context, in *mtproto.TLChannelsDeleteChannel) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsDeleteChannel(in)
}

// ChannelsExportMessageLink implements RPCChannelsServer
func (s *Service) ChannelsExportMessageLink(ctx context.Context, in *mtproto.TLChannelsExportMessageLink) (*mtproto.ExportedMessageLink, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsExportMessageLink(in)
}

// ChannelsToggleSignatures implements RPCChannelsServer
func (s *Service) ChannelsToggleSignatures(ctx context.Context, in *mtproto.TLChannelsToggleSignatures) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsToggleSignatures(in)
}

// ChannelsGetAdminedPublicChannels implements RPCChannelsServer
func (s *Service) ChannelsGetAdminedPublicChannels(ctx context.Context, in *mtproto.TLChannelsGetAdminedPublicChannels) (*mtproto.Messages_Chats, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetAdminedPublicChannels(in)
}

// ChannelsEditBanned implements RPCChannelsServer
func (s *Service) ChannelsEditBanned(ctx context.Context, in *mtproto.TLChannelsEditBanned) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsEditBanned(in)
}

// ChannelsGetAdminLog implements RPCChannelsServer
func (s *Service) ChannelsGetAdminLog(ctx context.Context, in *mtproto.TLChannelsGetAdminLog) (*mtproto.Channels_AdminLogResults, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetAdminLog(in)
}

// ChannelsSetStickers implements RPCChannelsServer
func (s *Service) ChannelsSetStickers(ctx context.Context, in *mtproto.TLChannelsSetStickers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsSetStickers(in)
}

// ChannelsReadMessageContents implements RPCChannelsServer
func (s *Service) ChannelsReadMessageContents(ctx context.Context, in *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsReadMessageContents(in)
}

// ChannelsDeleteHistory9BAA9647 implements RPCChannelsServer
func (s *Service) ChannelsDeleteHistory9BAA9647(ctx context.Context, in *mtproto.TLChannelsDeleteHistory9BAA9647) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsDeleteHistory9BAA9647(in)
}

// ChannelsTogglePreHistoryHidden implements RPCChannelsServer
func (s *Service) ChannelsTogglePreHistoryHidden(ctx context.Context, in *mtproto.TLChannelsTogglePreHistoryHidden) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsTogglePreHistoryHidden(in)
}

// ChannelsGetGroupsForDiscussion implements RPCChannelsServer
func (s *Service) ChannelsGetGroupsForDiscussion(ctx context.Context, in *mtproto.TLChannelsGetGroupsForDiscussion) (*mtproto.Messages_Chats, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetGroupsForDiscussion(in)
}

// ChannelsSetDiscussionGroup implements RPCChannelsServer
func (s *Service) ChannelsSetDiscussionGroup(ctx context.Context, in *mtproto.TLChannelsSetDiscussionGroup) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsSetDiscussionGroup(in)
}

// ChannelsEditLocation implements RPCChannelsServer
func (s *Service) ChannelsEditLocation(ctx context.Context, in *mtproto.TLChannelsEditLocation) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsEditLocation(in)
}

// ChannelsToggleSlowMode implements RPCChannelsServer
func (s *Service) ChannelsToggleSlowMode(ctx context.Context, in *mtproto.TLChannelsToggleSlowMode) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsToggleSlowMode(in)
}

// ChannelsGetInactiveChannels implements RPCChannelsServer
func (s *Service) ChannelsGetInactiveChannels(ctx context.Context, in *mtproto.TLChannelsGetInactiveChannels) (*mtproto.Messages_InactiveChats, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetInactiveChannels(in)
}

// ChannelsDeleteParticipantHistory implements RPCChannelsServer
func (s *Service) ChannelsDeleteParticipantHistory(ctx context.Context, in *mtproto.TLChannelsDeleteParticipantHistory) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsDeleteParticipantHistory(in)
}

// ChannelsToggleParticipantsHidden implements RPCChannelsServer
func (s *Service) ChannelsToggleParticipantsHidden(ctx context.Context, in *mtproto.TLChannelsToggleParticipantsHidden) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsToggleParticipantsHidden(in)
}

// ChannelsEditCreator implements RPCChannelsServer
func (s *Service) ChannelsEditCreator(ctx context.Context, in *mtproto.TLChannelsEditCreator) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsEditCreator(in)
}

// ChannelsGetFutureCreatorAfterLeave implements RPCChannelsServer
func (s *Service) ChannelsGetFutureCreatorAfterLeave(ctx context.Context, in *mtproto.TLChannelsGetFutureCreatorAfterLeave) (*mtproto.User, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsGetFutureCreatorAfterLeave(in)
}

// ChannelsInviteToChannel199F3A6C implements RPCChannelsServer
func (s *Service) ChannelsInviteToChannel199F3A6C(ctx context.Context, in *mtproto.TLChannelsInviteToChannel199F3A6C) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsInviteToChannel199F3A6C(in)
}

// ChannelsDeleteHistoryAF369D42 implements RPCChannelsServer
func (s *Service) ChannelsDeleteHistoryAF369D42(ctx context.Context, in *mtproto.TLChannelsDeleteHistoryAF369D42) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	return c.ChannelsDeleteHistoryAF369D42(in)
}
