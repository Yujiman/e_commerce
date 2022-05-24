package deliveryPoint

import (
	"net/http"

	helperError "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/http"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/middleware"
	serviceDeliveryPoint "github.com/Yujiman/e_commerce/userProfile/gatway/internal/service/deliveryPoint"
	"github.com/Yujiman/e_commerce/userProfile/gatway/internal/service/deliveryPointUser"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	//Handler of gRPC
	dto := &DTO{UserId: request.Header.Get(middleware.OauthUserId)}

	deliveryId, err := deliveryPointUser.GetDeliveryPointById(dto.UserId)
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}
	data, err := serviceDeliveryPoint.GetDeliveryPoint(*deliveryId)
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	// Response
	helperHttp.JsonResponse(response, data)
}
