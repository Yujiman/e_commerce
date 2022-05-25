package group

import (
	"net/http"

	helperError "github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/helper/http"
	serviceGroup "github.com/Yujiman/e_commerce/goods/gatway/internal/service/group"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	//Handler of gRPC
	resp, err := serviceGroup.NewClient().GetAllGroups()
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	// Response
	helperHttp.JsonResponse(response, resp)
}
