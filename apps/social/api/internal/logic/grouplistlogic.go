package logic

import (
	"context"

	"ZeZeIM/apps/social/api/internal/svc"
	"ZeZeIM/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户申群列表
func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList(req *types.GroupListReq) (resp *types.GroupListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
