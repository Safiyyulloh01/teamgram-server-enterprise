package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChannelsLeaveChannel removes a user from a channel
func (c *ChannelsCore) ChannelsLeaveChannel(in *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()

	_, err := c.svcCtx.Dao.ChatClient.Client().ChatDeleteChatUser(c.ctx, &chat.TLChatDeleteChatUser{
		ChatId:       channelId,
		OperatorId:   c.MD.UserId,
		DeleteUserId: c.MD.UserId,
	})
	if err != nil {
		return nil, err
	}

	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}
