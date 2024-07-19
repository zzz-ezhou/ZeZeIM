package logic

import (
	"ZeZeIM/apps/user/models"
	"context"
	"github.com/jinzhu/copier"

	"ZeZeIM/apps/user/rpc/internal/svc"
	"ZeZeIM/apps/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllUserLogic {
	return &AllUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllUserLogic) AllUser(in *user.AllUserReq) (*user.AllUserResp, error) {
	var (
		userEntitys []*models.Users
		err         error
	)

	userEntitys, err = l.svcCtx.UsersModel.AllUser(l.ctx)

	if err != nil {
		return nil, err
	}
	var resp []*user.UserEntity
	err = copier.Copy(&resp, &userEntitys)
	if err != nil {
		return nil, err
	}

	return &user.AllUserResp{
		User: resp,
	}, nil
}
