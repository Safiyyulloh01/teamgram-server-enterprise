package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// ChannelsGetFullChannel returns full channel info including participants
func (c *ChannelsCore) ChannelsGetFullChannel(in *mtproto.TLChannelsGetFullChannel) (*mtproto.Messages_ChatFull, error) {
	channelId := in.GetChannel().GetChannelId()

	mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})
	if err != nil {
		return nil, err
	}

	fullChat := mtproto.MakeTLChannelFull(&mtproto.ChatFull{
		Id:                 channelId,
		Participants:       mtproto.MakeTLChannelParticipants(&mtproto.ChannelParticipants{}).To_ChannelParticipants(),
		ChatPhoto:          mtproto.MakeTLPhotoEmpty(&mtproto.Photo{}).To_Photo(),
		NotifySettings:     mtproto.MakeTLPeerNotifySettings(&mtproto.PeerNotifySettings{}).To_PeerNotifySettings(),
		ExportedInvite:     mtproto.MakeTLChatInviteExported(&mtproto.ExportedChatInvite{}).To_ExportedChatInvite(),
		BotInfo:            []*mtproto.BotInfo{},
		PinnedMsgId:        0,
		FolderId:           0,
		CanSetUsername:     true,
		CanSetStickers:     true,
		CanViewParticipants: true,
		CanViewStatistics:  false,
		Stickerset:         mtproto.MakeTLStickerSetNotSet(&mtproto.StickerSet{}).To_StickerSet(),
	}).To_ChatFull()

	creatorId := mutableChat.GetCreatorId()
	user, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &user.TLUserGetMutableUsers{
		Id: []int64{creatorId},
	})

	return mtproto.MakeTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: fullChat,
		Chats:    []*mtproto.Chat{mutableChat.ToUnsafeChat(c.MD.UserId)},
		Users:    user.GetUserListByIdList(c.MD.UserId, creatorId),
	}).To_Messages_ChatFull(), nil
}
