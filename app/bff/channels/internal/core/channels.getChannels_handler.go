package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChannelsGetChannels returns info about multiple channels
func (c *ChannelsCore) ChannelsGetChannels(in *mtproto.TLChannelsGetChannels) (*mtproto.Messages_Chats, error) {
	ids := in.GetId()
	chats := make([]*mtproto.Chat, 0, len(ids))

	for _, id := range ids {
		mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
			ChatId: id.GetChannelId(),
		})
		if err == nil && mutableChat != nil {
			chats = append(chats, mutableChat.ToUnsafeChat(c.MD.UserId))
		}
	}

	return mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: chats,
	}).To_Messages_Chats(), nil
}
