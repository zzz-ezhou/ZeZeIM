package friend

import (
	"ZeZeIM/apps/social/rpc/socialclient"
	"ZeZeIM/util/ctxdata"
	"context"

	"ZeZeIM/apps/social/api/internal/svc"
	"ZeZeIM/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友申请
func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInLogic) FriendPutIn(req *types.FriendPutInReq) (resp *types.FriendPutInResp, err error) {
	// todo: add your logic here and delete this line

	uid := ctxdata.GetUId(l.ctx)

	_, err = l.svcCtx.Social.FriendPutIn(l.ctx, &socialclient.FriendPutInReq{
		UserId:  uid,
		ReqUid:  req.UserId,
		ReqMsg:  req.ReqMsg,
		ReqTime: req.ReqTime,
	})

	return
}
