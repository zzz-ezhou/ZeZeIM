package logic

import (
	"context"

	"ZeZeIM/apps/social/api/internal/svc"
	"ZeZeIM/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请进群
func NewGroupPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInLogic {
	return &GroupPutInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPutInLogic) GroupPutIn(req *types.GroupPutInReq) (resp *types.GroupPutInResp, err error) {
	// todo: add your logic here and delete this line

	return
}
