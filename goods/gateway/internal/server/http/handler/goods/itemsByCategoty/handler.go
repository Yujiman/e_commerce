package itemsByCategoty

import (
	"net/http"

	helperError "github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/helper/http"
	helperValidator "github.com/Yujiman/e_commerce/goods/gatway/internal/server/http/helper/validator"
	serviceAggregatorItems "github.com/Yujiman/e_commerce/goods/gatway/internal/service/aggregatorItems"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	dto := &DTO{}

	// Filling DTO from request
	if !helperHttp.BindRequest(response, request, dto) {
		return
	}

	// Validate DTO
	if !helperValidator.Validate(dto, response) {
		return
	}

	//Handler of gRPC
	resp, err := serviceAggregatorItems.NewClient().GetItemsByCategory(dto.Category)
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}

	// Response
	helperHttp.JsonResponse(response, resp)
}
