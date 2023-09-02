package interfaces

import (
	"msg/app/msg-common/dao/model"
	"msg/app/msg-common/types"
)

type BuilderContent interface {
	BuilderContent(messageTemplate model.MessageTemplate, messageParam types.MessageParam) interface{}
}
