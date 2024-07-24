package record

import (
	"ZChat/apps/single-chat/api/internal/svc"
	"ZChat/apps/single-chat/api/internal/types"
	"ZChat/pkg/constants"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type RecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordListLogic {
	return &RecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordListLogic) RecordList(req *types.RecordListRequest) (resp *types.RecordListResponse, err error) {
	pattern := fmt.Sprintf("%s:from:%d:to:%d", constants.SingleChatMsg, req.FriendUid, req.Uid) // 生成Redis键
	recodeJsonList, err := l.svcCtx.Redis.LrangeCtx(l.ctx, pattern, 0, -1)
	if err != nil {
		logx.Errorf("l.svcCtx.Redis.KeysCtx error: %v", err)
		return nil, err
	}
	var messages []types.MessageInfo

	for _, recodeJson := range recodeJsonList {
		m := types.MessageInfo{}
		fmt.Println(recodeJson)
		err := json.Unmarshal([]byte(recodeJson), &m)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		messages = append(messages, m)
	}

	return &types.RecordListResponse{
		RecordList: messages,
	}, err
}
