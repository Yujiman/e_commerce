package cities

import (
	"log"
	"net/http"

	helperError "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/error"
	helperHttp "github.com/Yujiman/e_commerce/userProfile/gatway/internal/server/http/helper/http"
	serviceCity "github.com/Yujiman/e_commerce/userProfile/gatway/internal/service/city"
)

func Handle(response http.ResponseWriter, request *http.Request) {
	//Handler of gRPC
	data, err := serviceCity.GetAllCity()
	if err != nil {
		helperHttp.ErrorResponse(err, response, helperError.GetStatusCodeErrFromGRPC(err))
		return
	}
	log.Println(data)
	// Response
	helperHttp.JsonResponse(response, data)
}
