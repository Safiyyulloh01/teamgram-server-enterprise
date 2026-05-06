package core

import (
	"github.com/teamgram/proto/mtproto"
)

func (c *BotsCore) BotsSendCustomRequest(in *mtproto.TLBotsSendCustomRequest) (*mtproto.DataJSON, error) {
	return mtproto.MakeTLDataJSON(&mtproto.DataJSON{
		Data: "{}",
	}).To_DataJSON(), nil
}

func (c *BotsCore) BotsAnswerWebhookJSONQuery(in *mtproto.TLBotsAnswerWebhookJSONQuery) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *BotsCore) BotsSetBotCommands(in *mtproto.TLBotsSetBotCommands) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *BotsCore) BotsResetBotCommands(in *mtproto.TLBotsResetBotCommands) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *BotsCore) BotsGetBotCommands(in *mtproto.TLBotsGetBotCommands) (*mtproto.Vector_BotCommand, error) {
	return &mtproto.Vector_BotCommand{
		Datas: []*mtproto.BotCommand{},
	}, nil
}

func (c *BotsCore) BotsSetBotInfo(in *mtproto.TLBotsSetBotInfo) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}

func (c *BotsCore) BotsGetBotInfoDCD914FD(in *mtproto.TLBotsGetBotInfoDCD914FD) (*mtproto.Bots_BotInfo, error) {
	return mtproto.MakeTLBotsBotInfo(&mtproto.Bots_BotInfo{
		Description: "",
		Name:        "",
		About:       "",
	}).To_Bots_BotInfo(), nil
}

func (c *BotsCore) BotsGetAdminedBots(in *mtproto.TLBotsGetAdminedBots) (*mtproto.Vector_User, error) {
	return &mtproto.Vector_User{
		Datas: []*mtproto.User{},
	}, nil
}

func (c *BotsCore) BotsGetBotInfo75EC12E6(in *mtproto.TLBotsGetBotInfo75EC12E6) (*mtproto.Vector_String, error) {
	return &mtproto.Vector_String{
		Datas: []string{},
	}, nil
}
