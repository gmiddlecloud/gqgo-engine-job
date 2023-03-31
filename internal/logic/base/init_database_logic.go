package base

import (
	"context"

	"github.com/gmiddlecloud/gqgo-engine-job/internal/utils/dberrorhandler"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/gmiddlecloud/gqgo-engine-common/enum/errorcode"
	"github.com/gmiddlecloud/gqgo-engine-common/i18n"
	"github.com/gmiddlecloud/gqgo-engine-common/msg/logmsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/gmiddlecloud/gqgo-engine-job/internal/mqs/amq/types/pattern"
	"github.com/gmiddlecloud/gqgo-engine-job/internal/svc"
	"github.com/gmiddlecloud/gqgo-engine-job/types/job"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *job.Empty) (*job.BaseResp, error) {

	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewCodeError(errorcode.Internal, err.Error())
	}

	count, err := l.svcCtx.DB.Task.Query().Count(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, "database error")
	}

	if count != 0 {
		return nil, errorx.NewInvalidArgumentError(i18n.AlreadyInit)
	}

	err = l.insertTaskData()
	if err != nil {
		return nil, err
	}

	return &job.BaseResp{
		Msg: i18n.Success,
	}, nil
}

func (l *InitDatabaseLogic) insertTaskData() error {
	err := l.svcCtx.DB.Task.Create().
		SetName("hello_world").
		SetTaskGroup("base").
		SetCronExpression("@every 5s").
		SetPattern(pattern.RecordHelloWorld).
		SetPayload("{\"name\": \"Mike (DPTask 5s)\"}").
		Exec(l.ctx)

	if err != nil {
		return err
	}

	return nil
}
