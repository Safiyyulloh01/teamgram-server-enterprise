package dao

import (
	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/marmota/pkg/net/rpcx"
	"github.com/teamgram/teamgram-server/app/bff/channels/internal/config"
	msg_client "github.com/teamgram/teamgram-server/app/messenger/msg/msg/client"
	sync_client "github.com/teamgram/teamgram-server/app/messenger/sync/client"
	chat_client "github.com/teamgram/teamgram-server/app/service/biz/chat/client"
	dialog_client "github.com/teamgram/teamgram-server/app/service/biz/dialog/client"
	message_client "github.com/teamgram/teamgram-server/app/service/biz/message/client"
	user_client "github.com/teamgram/teamgram-server/app/service/biz/user/client"
	idgen_client "github.com/teamgram/teamgram-server/app/service/idgen/client"
	media_client "github.com/teamgram/teamgram-server/app/service/media/client"
)

type Dao struct {
	user_client.UserClient
	ChatClient      *chat_client.ChatClientHelper
	msg_client.MsgClient
	media_client.MediaClient
	dialog_client.DialogClient
	idgen_client.IDGenClient2
	message_client.MessageClient
	sync_client.SyncClient
}

func New(c config.Config) *Dao {
	return &Dao{
		UserClient:      user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:      chat_client.NewChatClientHelper(rpcx.GetCachedRpcClient(c.ChatClient)),
		MsgClient:       msg_client.NewMsgClient(rpcx.GetCachedRpcClient(c.MsgClient)),
		MediaClient:     media_client.NewMediaClient(rpcx.GetCachedRpcClient(c.MediaClient)),
		DialogClient:    dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
		IDGenClient2:    idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
		MessageClient:   message_client.NewMessageClient(rpcx.GetCachedRpcClient(c.MessageClient)),
		SyncClient:      sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
	}
}
