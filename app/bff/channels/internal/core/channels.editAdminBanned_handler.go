package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChannelsEditAdmin edits admin permissions in a channel
func (c *ChannelsCore) ChannelsEditAdmin(in *mtproto.TLChannelsEditAdmin) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()

	_, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatAdmin(c.ctx, &chat.TLChatEditChatAdmin{
		ChatId:           channelId,
		OperatorId:       c.MD.UserId,
		EditChatAdminId:  in.GetUserId().GetUserId(),
		IsAdmin:          mtproto.MakeFlagsBool(in.GetIsAdmin()),
	})
	if err != nil {
		return nil, err
	}

	mutableChat, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})

	updatesHelper := mtproto.MakeUpdatesHelper(mutableChat.ToUnsafeChat(c.MD.UserId))
	return updatesHelper.ToUpdates(), nil
}

// ChannelsEditBanned edits banned rights in a channel
func (c *ChannelsCore) ChannelsEditBanned(in *mtproto.TLChannelsEditBanned) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()

	_, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatDefaultBannedRights(c.ctx, &chat.TLChatEditChatDefaultBannedRights{
		ChatId:      channelId,
		OperatorId:  c.MD.UserId,
		BannedRights: in.GetBannedRights(),
	})
	if err != nil {
		return nil, err
	}

	mutableChat, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})

	updatesHelper := mtproto.MakeUpdatesHelper(mutableChat.ToUnsafeChat(c.MD.UserId))
	return updatesHelper.ToUpdates(), nil
}
