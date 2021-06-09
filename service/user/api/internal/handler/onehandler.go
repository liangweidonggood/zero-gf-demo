package handler

import (
	"net/http"

	"zero-gf-demo/service/user/api/internal/logic"
	"zero-gf-demo/service/user/api/internal/svc"
	"zero-gf-demo/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func oneHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewOneLogic(r.Context(), ctx)
		user, err := l.One(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, user)
		}
	}
}
