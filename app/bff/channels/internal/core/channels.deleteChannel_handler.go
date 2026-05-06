package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChannelsDeleteChannel deletes a channel
func (c *ChannelsCore) ChannelsDeleteChannel(in *mtproto.TLChannelsDeleteChannel) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()

	mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})
	if err != nil {
		return nil, err
	}

	if mutableChat.GetCreatorId() != c.MD.UserId {
		return nil, mtproto.ErrChatAdminRequired
	}

	_, err = c.svcCtx.Dao.ChatClient.Client().ChatDeleteChat(c.ctx, &chat.TLChatDeleteChat{
		ChatId:     channelId,
		OperatorId: c.MD.UserId,
	})
	if err != nil {
		return nil, err
	}

	updatesHelper := mtproto.MakeUpdatesHelper()
	return updatesHelper.ToUpdates(), nil
}

// ChannelsGetParticipants returns channel participants list
func (c *ChannelsCore) ChannelsGetParticipants(in *mtproto.TLChannelsGetParticipants) (*mtproto.Channels_ChannelParticipants, error) {
	channelId := in.GetChannel().GetChannelId()

	participantIdList, err := c.svcCtx.Dao.ChatClient.Client().ChatGetChatParticipantIdList(c.ctx, &chat.TLChatGetChatParticipantIdList{
		ChatId: channelId,
	})
	if err != nil {
		return nil, err
	}

	participants := make([]*mtproto.ChannelParticipant, 0, len(participantIdList.GetDatas()))
	for _, userId := range participantIdList.GetDatas() {
		participant := mtproto.MakeTLChannelParticipant(&mtproto.ChannelParticipant{
			UserId: userId,
			Date:   int32(mtproto.Now().Unix()),
		}).To_ChannelParticipant()
		participants = append(participants, participant)
	}

	return mtproto.MakeTLChannelsChannelParticipants(&mtproto.Channels_ChannelParticipants{
		Count:        int32(len(participants)),
		Participants: participants,
		Chats:        []*mtproto.Chat{},
		Users:        []*mtproto.User{},
	}).To_Channels_ChannelParticipants(), nil
}
