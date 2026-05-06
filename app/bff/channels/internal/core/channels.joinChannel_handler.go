package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChannelsJoinChannel joins a user to a channel
func (c *ChannelsCore) ChannelsJoinChannel(in *mtproto.TLChannelsJoinChannel) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()

	mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})
	if err != nil {
		return nil, err
	}

	// Add user as participant
	mutableChat, err = c.svcCtx.Dao.ChatClient.Client().ChatAddChatUser(c.ctx, &chat.TLChatAddChatUser{
		ChatId:   channelId,
		InviterId: c.MD.UserId,
		UserId:   c.MD.UserId,
	})
	if err != nil {
		return nil, err
	}

	updatesHelper := mtproto.MakeUpdatesHelper(mutableChat.ToUnsafeChat(c.MD.UserId))
	return updatesHelper.ToUpdates(), nil
}
