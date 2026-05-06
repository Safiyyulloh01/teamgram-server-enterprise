package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChannelsEditTitle edits the channel title
func (c *ChannelsCore) ChannelsEditTitle(in *mtproto.TLChannelsEditTitle) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()

	mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatTitle(c.ctx, &chat.TLChatEditChatTitle{
		ChatId:     channelId,
		EditUserId: c.MD.UserId,
		Title:      in.GetTitle(),
	})
	if err != nil {
		return nil, err
	}

	updatesHelper := mtproto.MakeUpdatesHelper(mutableChat.ToUnsafeChat(c.MD.UserId))
	return updatesHelper.ToUpdates(), nil
}

// ChannelsEditAbout edits the channel description
func (c *ChannelsCore) ChannelsEditAbout(in *mtproto.TLChannelsEditAbout) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()

	_, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatAbout(c.ctx, &chat.TLChatEditChatAbout{
		ChatId:     channelId,
		EditUserId: c.MD.UserId,
		About:      in.GetAbout(),
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
