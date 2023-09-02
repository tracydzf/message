package handler

import (
	"msg/app/msg-web/msg_api/internal/logic"
	"msg/app/msg-web/msg_api/internal/svc"
	"msg/app/msg-web/msg_api/internal/types"
	"msg/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendLogic(r.Context(), svcCtx)
		resp, err := l.Send(req)
		result.HttpResult(r, w, resp, err)

	}
}
