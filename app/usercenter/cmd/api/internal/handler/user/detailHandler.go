package user

import (
	"go-zero-for-k8s/common/result"
	"net/http"

	"go-zero-for-k8s/app/usercenter/cmd/api/internal/logic/user"
	"go-zero-for-k8s/app/usercenter/cmd/api/internal/svc"
	"go-zero-for-k8s/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewDetailLogic(r.Context(), ctx)
		l.Logger.Infof("api-handler-req: %+v", req)
		resp, err := l.Detail(req)
		result.HttpResult(r, w, resp, err)
	}
}
