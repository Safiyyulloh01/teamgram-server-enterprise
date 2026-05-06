package core

import (
	"context"

	"github.com/teamgram/marmota/pkg/metadata"
	"github.com/teamgram/teamgram-server/app/bff/reactions/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReactionsCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *ReactionsCore {
	return &ReactionsCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
