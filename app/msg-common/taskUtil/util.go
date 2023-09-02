package taskUtil

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/jsonx"
	"msg/app/msg-common/enums/channelType"
	"msg/app/msg-common/enums/messageType"
	"msg/app/msg-common/types"
	"strings"
	"time"
)

const PhoneRegex = "^((13[0-9])|(14[5,7,9])|(15[0-3,5-9])|(166)|(17[0-9])|(18[0-9])|(19[1,8,9]))\\d{8}$"
const PARAM = "?"

func GenerateBusinessId(templateId int64, templateType int) int64 {
	str := fmt.Sprintf("%d%s", int64(templateType*1000000)+templateId, time.Now().Format("20060102"))
	return cast.ToInt64(str)
}
func GenerateUrl(url string, templateId int64, templateType int) string {
	businessId := GenerateBusinessId(templateId, templateType)
	if strings.Contains(url, "?") {
		return fmt.Sprintf("%s?track_code_bid=%d", url, businessId)
	}
	return fmt.Sprintf("%s&track_code_bid=%d", url, businessId)
}

// ReplaceByMap returns a copy of `origin`,
// which is replaced by a map in unordered way, case-sensitively.
func ReplaceByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = strings.Replace(origin, "{$"+k+"}", v, -1)
	}
	return origin
}

func GetAllGroupIds() []string {
	list := make([]string, 0)
	for _, ct := range channelType.TypeCodeEn {
		for _, mt := range messageType.TypeCodeEn {
			list = append(list, ct+"."+mt)
		}
	}
	return list
}
func GetGroupIdByTaskInfo(info types.TaskInfo) string {
	channelCodeEn := channelType.TypeCodeEn[info.SendChannel]
	msgCodeEn := messageType.TypeCodeEn[info.MsgType]
	return channelCodeEn + "." + msgCodeEn
}

func GetContentModel(contentModel interface{}, v interface{}) {
	marshal, _ := jsonx.Marshal(contentModel)
	_ = jsonx.Unmarshal(marshal, &v)
}