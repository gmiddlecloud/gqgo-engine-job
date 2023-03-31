package base

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gmiddlecloud/gqgo-engine-job/ent/task"
	"github.com/gmiddlecloud/gqgo-engine-job/internal/enum/taskresult"
	"github.com/gmiddlecloud/gqgo-engine-job/internal/mqs/amq/types/pattern"
	"github.com/gmiddlecloud/gqgo-engine-job/internal/utils/dberrorhandler"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"github.com/gmiddlecloud/gqgo-engine-job/internal/mqs/amq/types/payload"
	"github.com/gmiddlecloud/gqgo-engine-job/internal/svc"
)

type HelloWorldHandler struct {
	svcCtx *svc.ServiceContext
	taskId uint64
}

func NewHelloWorldHandler(svcCtx *svc.ServiceContext) *HelloWorldHandler {
	taskId, _ := svcCtx.DB.Task.Query().Where(task.PatternEQ(pattern.RecordHelloWorld)).First(context.Background())
	return &HelloWorldHandler{
		svcCtx: svcCtx,
		taskId: taskId.ID,
	}
}

// ProcessTask if return err != nil , asynq will retry | 如果返回错误不为空则会重试
func (l *HelloWorldHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p payload.HelloWorldPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(err, "failed to umarshal the payload :%s", string(t.Payload()))
	}

	startTime := time.Now()
	fmt.Printf("Hi! %s\n", p.Name)
	finishTime := time.Now()

	err := l.svcCtx.DB.TaskLog.Create().
		SetStartedAt(startTime).
		SetFinishedAt(finishTime).
		SetResult(taskresult.Success).
		SetTasksID(l.taskId).
		Exec(context.Background())

	if err != nil {
		return dberrorhandler.DefaultEntError(logx.WithContext(context.Background()), err,
			"failed to save task log to database")
	}

	return nil
}
