package limit

import (
	"msg/app/msg-common/types"
	"msg/app/msg_job/internal/services/deduplication/structs"
)

func deduplicationAllKey(service structs.DeduplicationService, taskInfo *types.TaskInfo) []string {
	var newRows []string
	for _, receiver := range taskInfo.Receiver {
		newRows = append(newRows, deduplicationSingleKey(service, taskInfo, receiver))
	}
	return newRows
}
func deduplicationSingleKey(service structs.DeduplicationService, taskInfo *types.TaskInfo, receiver string) string {
	return service.DeduplicationSingleKey(taskInfo, receiver)
}
