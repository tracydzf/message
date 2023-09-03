package limit

import (
	"context"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"msg/app/msg-common/types"
	"msg/app/msg_job/internal/services/deduplication/structs"
	"msg/app/msg_job/internal/svc"
	"msg/common/utils/timex"
)

const slideWindowLimitServiceTag = "SW_"
const slideWindowLimitServicePrefix = "slideWindowLimitServicePrefix_"

// 滑动窗口去重服务
type slideWindowLimitService struct {
	svcCtx *svc.ServiceContext
}

func (s slideWindowLimitService) LimitFilter(ctx context.Context, duplication structs.DeduplicationService, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) (filterReceiver []string, err error) {
	filterReceiver = make([]string, 0)
	period := timex.GetDisTodayEnd()
	periodLimiter := limit.NewPeriodLimit(int(period), param.Num, s.svcCtx.Redis, slideWindowLimitServicePrefix)

	for _, receiver := range taskInfo.Receiver {
		key := slideWindowLimitServiceTag + deduplicationSingleKey(duplication, taskInfo, receiver)
		code, err := periodLimiter.Take(key)
		if err != nil {
			logx.Errorw("slideWindowLimitService Take ", logx.Field("err", err))
			continue
		}
		//表示到了上限 直接过滤掉
		if code == limit.OverQuota {
			filterReceiver = append(filterReceiver, receiver)
		}

	}
	return filterReceiver, nil
}

func NewSlideWindowLimitService(svcCtx *svc.ServiceContext) structs.LimitService {
	return &slideWindowLimitService{svcCtx: svcCtx}
}
