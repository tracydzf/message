package interfaces

import (
	"context"
	"msg/app/msg-common/types"
)

type Process interface {
	Process(ctx context.Context, sendTaskModel *types.SendTaskModel) error
}
