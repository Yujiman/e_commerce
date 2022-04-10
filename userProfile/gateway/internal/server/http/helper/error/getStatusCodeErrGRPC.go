package error

import (
	"net/http"

	"google.golang.org/grpc/status"
)

func GetStatusCodeErrFromGRPC(err error) int {
	st, ok := status.FromError(err)

	if !ok || st.Code().String() == "Unavailable" {
		return http.StatusInternalServerError
	}
	return int(st.Code())
}
