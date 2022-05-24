package deliveryPoints

import (
	"net/http"

	helperError "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/http"
	serviceDeliveryPoint "github.com/Yujiman/e_commerce/userProfile/gatway/internal/service/deliveryPoint"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	//Handler of gRPC

	data, err := serviceDeliveryPoint.GetAllDeliveryPoint()
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	// Response
	helperHttp.JsonResponse(response, data)
}
