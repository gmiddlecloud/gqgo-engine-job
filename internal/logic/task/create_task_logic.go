package task

import (
	"context"

	"github.com/gmiddlecloud/gqgo-engine-job/internal/svc"
	"github.com/gmiddlecloud/gqgo-engine-job/internal/utils/dberrorhandler"
	"github.com/gmiddlecloud/gqgo-engine-job/types/job"

	"github.com/gmiddlecloud/gqgo-engine-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTaskLogic) CreateTask(in *job.TaskInfo) (*job.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Task.Create().
		SetStatus(uint8(in.Status)).
		SetName(in.Name).
		SetTaskGroup(in.TaskGroup).
		SetCronExpression(in.CronExpression).
		SetPattern(in.Pattern).
		SetPayload(in.Payload).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
