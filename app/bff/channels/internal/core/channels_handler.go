package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

func (c *ChannelsCore) ChannelsCreateChannel(in *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	channelId := c.svcCtx.Dao.IDGenClient2.NextId(c.ctx)
	_ = channelId

	// Use chat service to create
	mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})
	if err != nil {
		mutableChat = nil
	}

	chatObj := &mtproto.Chat{}
	if mutableChat != nil {
		chatObj = mutableChat.ToUnsafeChat(c.MD.UserId)
	}

	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{chatObj},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()

	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsJoinChannel(in *mtproto.TLChannelsJoinChannel) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()
	mutableChat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})
	if err != nil {
		return nil, err
	}

	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{mutableChat.ToUnsafeChat(c.MD.UserId)},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsLeaveChannel(in *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsGetChannels(in *mtproto.TLChannelsGetChannels) (*mtproto.Messages_Chats, error) {
	chats := make([]*mtproto.Chat, 0)
	for _, id := range in.GetId() {
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
	userData, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &user.TLUserGetMutableUsers{
		Id: []int64{creatorId},
	})

	return mtproto.MakeTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: fullChat,
		Chats:    []*mtproto.Chat{mutableChat.ToUnsafeChat(c.MD.UserId)},
		Users:    userData.GetUserListByIdList(c.MD.UserId, creatorId),
	}).To_Messages_ChatFull(), nil
}

func (c *ChannelsCore) ChannelsEditTitle(in *mtproto.TLChannelsEditTitle) (*mtproto.Updates, error) {
	channelId := in.GetChannel().GetChannelId()
	mutableChat, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chat.TLChatGetMutableChat{
		ChatId: channelId,
	})

	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{mutableChat.ToUnsafeChat(c.MD.UserId)},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsDeleteChannel(in *mtproto.TLChannelsDeleteChannel) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsGetParticipants(in *mtproto.TLChannelsGetParticipants) (*mtproto.Channels_ChannelParticipants, error) {
	participantIdList, err := c.svcCtx.Dao.ChatClient.Client().ChatGetChatParticipantIdList(c.ctx, &chat.TLChatGetChatParticipantIdList{
		ChatId: in.GetChannel().GetChannelId(),
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

func (c *ChannelsCore) ChannelsEditAdmin(in *mtproto.TLChannelsEditAdmin) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsEditBanned(in *mtproto.TLChannelsEditBanned) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsToggleSignatures(in *mtproto.TLChannelsToggleSignatures) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsInviteToChannel199F3A6C(in *mtproto.TLChannelsInviteToChannel199F3A6C) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsInviteToChannelC9E33D54(in *mtproto.TLChannelsInviteToChannelC9E33D54) (*mtproto.Messages_InvitedUsers, error) {
	return mtproto.MakeTLMessagesInvitedUsers(&mtproto.Messages_InvitedUsers{
		Updates: &mtproto.Updates{},
		MissingInvitees: []*mtproto.MissingInvitee{},
	}).To_Messages_InvitedUsers(), nil
}

func (c *ChannelsCore) ChannelsGetParticipant(in *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	participant := mtproto.MakeTLChannelParticipant(&mtproto.ChannelParticipant{
		UserId: in.GetUserId().GetUserId(),
		Date:   int32(mtproto.Now().Unix()),
	}).To_ChannelParticipant()

	return mtproto.MakeTLChannelsChannelParticipant(&mtproto.Channels_ChannelParticipant{
		Participant: participant,
		Chats:       []*mtproto.Chat{},
		Users:       []*mtproto.User{},
	}).To_Channels_ChannelParticipant(), nil
}

func (c *ChannelsCore) ChannelsReadHistory(in *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsDeleteMessages(in *mtproto.TLChannelsDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      1,
		PtsCount: int32(len(in.GetId())),
	}).To_Messages_AffectedMessages(), nil
}

func (c *ChannelsCore) ChannelsGetMessages(in *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	return mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages(), nil
}

func (c *ChannelsCore) ChannelsEditPhoto(in *mtproto.TLChannelsEditPhoto) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsExportMessageLink(in *mtproto.TLChannelsExportMessageLink) (*mtproto.ExportedMessageLink, error) {
	return mtproto.MakeTLExportedMessageLink(&mtproto.ExportedMessageLink{
		Link: "",
	}).To_ExportedMessageLink(), nil
}

func (c *ChannelsCore) ChannelsGetAdminedPublicChannels(in *mtproto.TLChannelsGetAdminedPublicChannels) (*mtproto.Messages_Chats, error) {
	return mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
	}).To_Messages_Chats(), nil
}

func (c *ChannelsCore) ChannelsGetAdminLog(in *mtproto.TLChannelsGetAdminLog) (*mtproto.Channels_AdminLogResults, error) {
	return mtproto.MakeTLChannelsAdminLogResults(&mtproto.Channels_AdminLogResults{
		Events: []*mtproto.ChannelAdminLogEvent{},
		Chats:  []*mtproto.Chat{},
		Users:  []*mtproto.User{},
	}).To_Channels_AdminLogResults(), nil
}

func (c *ChannelsCore) ChannelsSetStickers(in *mtproto.TLChannelsSetStickers) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsReadMessageContents(in *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsDeleteHistory9BAA9647(in *mtproto.TLChannelsDeleteHistory9BAA9647) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsTogglePreHistoryHidden(in *mtproto.TLChannelsTogglePreHistoryHidden) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsGetGroupsForDiscussion(in *mtproto.TLChannelsGetGroupsForDiscussion) (*mtproto.Messages_Chats, error) {
	return mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
	}).To_Messages_Chats(), nil
}

func (c *ChannelsCore) ChannelsSetDiscussionGroup(in *mtproto.TLChannelsSetDiscussionGroup) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsEditLocation(in *mtproto.TLChannelsEditLocation) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *ChannelsCore) ChannelsToggleSlowMode(in *mtproto.TLChannelsToggleSlowMode) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsGetInactiveChannels(in *mtproto.TLChannelsGetInactiveChannels) (*mtproto.Messages_InactiveChats, error) {
	return mtproto.MakeTLMessagesInactiveChats(&mtproto.Messages_InactiveChats{
		Dates: []int32{},
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}).To_Messages_InactiveChats(), nil
}

func (c *ChannelsCore) ChannelsDeleteParticipantHistory(in *mtproto.TLChannelsDeleteParticipantHistory) (*mtproto.Messages_AffectedHistory, error) {
	return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      1,
		PtsCount: 0,
		Offset:   0,
	}).To_Messages_AffectedHistory(), nil
}

func (c *ChannelsCore) ChannelsToggleParticipantsHidden(in *mtproto.TLChannelsToggleParticipantsHidden) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsEditCreator(in *mtproto.TLChannelsEditCreator) (*mtproto.Updates, error) {
	rUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(mtproto.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return rUpdates, nil
}

func (c *ChannelsCore) ChannelsGetFutureCreatorAfterLeave(in *mtproto.TLChannelsGetFutureCreatorAfterLeave) (*mtproto.User, error) {
	return mtproto.MakeTLUserEmpty(&mtproto.User{}).To_User(), nil
}

func (c *ChannelsCore) ChannelsDeleteHistoryAF369D42(in *mtproto.TLChannelsDeleteHistoryAF369D42) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
