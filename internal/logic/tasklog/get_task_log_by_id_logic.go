package tasklog

import (
	"context"

	"github.com/gmiddlecloud/gqgo-engine-job/internal/svc"
	"github.com/gmiddlecloud/gqgo-engine-job/internal/utils/dberrorhandler"
	"github.com/gmiddlecloud/gqgo-engine-job/types/job"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogByIdLogic {
	return &GetTaskLogByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskLogByIdLogic) GetTaskLogById(in *job.IDReq) (*job.TaskLogInfo, error) {
	result, err := l.svcCtx.DB.TaskLog.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.TaskLogInfo{
		Id:         result.ID,
		StartedAt:  result.StartedAt.UnixMilli(),
		FinishedAt: result.FinishedAt.UnixMilli(),
		Result:     uint32(result.Result),
	}, nil
}
