package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChannelsToggleSignatures toggles signatures on/off for channel posts
func (c *ChannelsCore) ChannelsToggleSignatures(in *mtproto.TLChannelsToggleSignatures) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()

	mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})
	if err != nil {
		return nil, err
	}

	updatesHelper := mtproto.MakeUpdatesHelper(mutableChat.ToUnsafeChat(c.MD.UserId))
	return updatesHelper.ToUpdates(), nil
}

// ChannelsInviteToChannel invites users to a channel
func (c *ChannelsCore) ChannelsInviteToChannel199F3A6C(in *mtproto.TLChannelsInviteToChannel199F3A6C) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()
	users := in.GetUsers()

	for _, userId := range users {
		_, err := c.svcCtx.Dao.ChatClient.Client().ChatAddChatUser(c.ctx, &chat.TLChatAddChatUser{
			ChatId:   channelId,
			InviterId: c.MD.UserId,
			UserId:   userId.GetUserId(),
		})
		if err != nil {
			return nil, err
		}
	}

	mutableChat, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})

	updatesHelper := mtproto.MakeUpdatesHelper(mutableChat.ToUnsafeChat(c.MD.UserId))
	return updatesHelper.ToUpdates(), nil
}

// ChannelsInviteToChannel (newer version)
func (c *ChannelsCore) ChannelsInviteToChannelC9E33D54(in *mtproto.TLChannelsInviteToChannelC9E33D54) (*mtproto.Messages_InvitedUsers, error) {
	channelId := in.GetChannel().GetChannelId()
	users := in.GetUsers()

	for _, userId := range users {
		_, err := c.svcCtx.Dao.ChatClient.Client().ChatAddChatUser(c.ctx, &chat.TLChatAddChatUser{
			ChatId:   channelId,
			InviterId: c.MD.UserId,
			UserId:   userId.GetUserId(),
		})
		if err != nil {
			return nil, err
		}
	}

	return mtproto.MakeTLMessagesInvitedUsers(&mtproto.Messages_InvitedUsers{
		Updates: &mtproto.Updates{},
		MissingInvitees: []*mtproto.MissingInvitee{},
	}).To_Messages_InvitedUsers(), nil
}
