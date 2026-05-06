package reactions_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/reactions/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/reactions/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/reactions/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
