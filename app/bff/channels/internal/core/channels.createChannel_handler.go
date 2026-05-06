package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/crypto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChannelsCreateChannel creates a new channel or megagroup
func (c *ChannelsCore) ChannelsCreateChannel(in *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	channelId, err := c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
	if err != nil {
		return nil, err
	}

	accessHash := crypto.GenerateAccessHash(channelId)
	isBroadcast := mtproto.FromBool(in.GetBroadcast())
	isMegagroup := mtproto.FromBool(in.GetMegagroup())

	channelType := int32(1)
	if isMegagroup {
		channelType = 2
	}

	mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatCreateChat2(c.ctx, &chat.TLChatCreateChat2{
		CreatorId:   c.MD.UserId,
		UserIdList:  []int64{c.MD.UserId},
		Title:       in.GetTitle(),
	})
	if err != nil {
		return nil, err
	}

	chatObj := mutableChat.ToUnsafeChat(c.MD.UserId)
	updatesHelper := mtproto.MakeUpdatesHelper(chatObj)

	return updatesHelper.ToUpdates(), nil
}
