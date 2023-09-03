package srv

import (
	"context"
	"msg/app/msg-common/types"
	"msg/app/msg_job/internal/services/deduplication/structs"
	"msg/common/utils/arrayUtils"
)

type deduplicationService struct {
}

func (c deduplicationService) Deduplication(ctx context.Context,
	limit structs.LimitService,
	service structs.DeduplicationService,
	taskInfo *types.TaskInfo,
	param structs.DeduplicationConfigItem) error {

	var newRows []string
	filter, err := limit.LimitFilter(ctx, service, taskInfo, param)
	if err != nil {
		return err
	}
	for _, s := range taskInfo.Receiver {
		if !arrayUtils.ArrayStringIn(filter, s) {
			newRows = append(newRows, s)
		}
	}
	taskInfo.Receiver = newRows
	return nil
}
